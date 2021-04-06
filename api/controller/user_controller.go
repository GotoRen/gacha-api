package controller

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"

	"github.com/fumist23/game-api/database"
	"github.com/fumist23/game-api/model"
)

// /user/createに対するハンドラ
// requestからnameを取り出してtoken生成してDBに保存して返す
func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body := r.Body
	defer body.Close()

	if r.Method != http.MethodPost {
		log.Printf("only POST methods are permitted")
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	var userCreateRequest model.UserCreateRequest
	if err := json.NewDecoder(body).Decode(&userCreateRequest); err != nil {
		log.Printf("failed to decode userCreateRequest: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	name := userCreateRequest.Name

	tokenStr, err := GenerateTokenWithName(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occurred in GenerateTokenWithName %v", err)
	}

	err = database.CreateUser(ctx, name, tokenStr)
	if err != nil {
		log.Printf("error occurred in database.requestUser: %v", err)
	}

	userCreateResponse := model.UserCreateResponse{
		Token: tokenStr,
	}

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(userCreateResponse)
	if err != nil {
		log.Printf("failed to encode userCreateResponse: %v", err)
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.Printf("failed to w.Write: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

// /user/loginに対するハンドラ
// bodyのidとnameから該当するuserのtokenを取得して返す
func LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body := r.Body
	defer body.Close()

	if r.Method != http.MethodPost {
		log.Printf("only POST methods are permitted")
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	var userLoginRequest model.UserLoginRequest
	if err := json.NewDecoder(body).Decode(&userLoginRequest); err != nil {
		log.Printf("failed to decode userLoginRequest: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	id := userLoginRequest.Id
	name := userLoginRequest.Name

	user, err := database.GetToken(ctx, id, name)

	if err != nil {
		log.Printf("error occurred in database.GetToken: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	token := user.Token

	if err := json.NewEncoder(w).Encode(token); err != nil {
		log.Printf("failed to encode token to json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

// /user/getに対するハンドラ
// headerのx-tokenからtokenを取り出してDBからfetchして該当するuserのnameを取得して返す
func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodGet {
		log.Printf("only GET methods are permitted")
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	token := r.Header.Get("x-token")
	if token == "" {
		log.Println("token must be included in header")
		w.WriteHeader(http.StatusBadRequest)
	}

	user, err := database.GetUser(ctx, token)

	if err != nil {
		log.Printf("error occurred in database.GetUser: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	name := user.Name

	if err := json.NewEncoder(w).Encode(name); err != nil {
		log.Printf("failed to encode name to json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

// /user/updateに対するハンドラ
// x-tokenからtokenを取り出して該当するuserを検証し、受け取ったnameを更新してDB更新して返す
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPut {
		log.Printf("only PUT methods are permitted")
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	token := r.Header.Get("x-token")
	isValidToken := database.VerifyToken(ctx, token)
	if !isValidToken {
		w.WriteHeader(http.StatusUnauthorized)
	}
	body := r.Body
	defer body.Close()

	var userUpdateRequest model.UserUpdateRequest

	if err := json.NewDecoder(body).Decode(&userUpdateRequest); err != nil {
		log.Printf("failed to decode request.Body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	name := userUpdateRequest.Name

	err := database.UpdateUser(ctx, token, name)
	if err != nil {
		log.Printf("failed to UpdateUser: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
