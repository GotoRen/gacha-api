package model

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserCreateResponse struct {
	Token string
}

type UserLoginRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserGetResponse struct {
	Name string
}

type UserUpdateRequest struct {
	Name string
}
