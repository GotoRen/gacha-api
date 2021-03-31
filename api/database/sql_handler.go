package database

import (
	"context"
	"log"

	"github.com/fumist23/game-api/model"
)

// nameをtokenを受け取って保存する
func CreateUser(ctx context.Context, name string, token string) error {
	_, err := DB.QueryContext(ctx, "INSERT INTO users(name, token) VALUES(?, ?)", name, token)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return err
	}

	return nil
}

// tokenが存在するかチェックする
func VerifyToken(ctx context.Context, token string) bool {
	row := DB.QueryRowContext(ctx, "SELECT * FROM users WHERE token=?", token)
	if err := row.Err(); err != nil {
		log.Printf("this token is invalid")
		return false
	}

	return true
}

// tokenを受け取って該当するuserのnameを取り出す

func GetUser(ctx context.Context, token string) (model.User, error) {
	row := DB.QueryRowContext(ctx, "SELECT id, name FROM users WHERE token=?", token)
	var user model.User
	if err := row.Scan(&user.Id, &user.Name); err != nil {
		log.Print("failed to get user name from database")
		return user, err
	}

	return user, nil
}

// tokenとnameを受け取ってtokenに該当するuserのnameを更新する
func UpdateUser(ctx context.Context, token string, name string) error {
	if _, err := DB.QueryContext(ctx, "UPDATE users SET name = ? WHERE token = ? ", name, token); err != nil {
		log.Printf("failed to update user: %v", err)
		return err
	}
	return nil
}

// GetGachaConfigs ガチャの設定情報を取得する
func GetGachaConfigs(ctx context.Context) ([]model.GachaConfig, error) {
	rows, err := DB.QueryContext(ctx, "SELECT reality, probability FROM gacha_configs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gachaConfigs := make([]model.GachaConfig, 0)

	for rows.Next() {
		var gachaConfig model.GachaConfig
		if err := rows.Scan(&gachaConfig.Reality, &gachaConfig.Probability); err != nil {
			return gachaConfigs, err
		}
		gachaConfigs = append(gachaConfigs, gachaConfig)
	}

	if err := rows.Err(); err != nil {
		return gachaConfigs, err
	}

	return gachaConfigs, nil
}

// GetCharacters キャラクターの情報を取得する
func GetCharacters(ctx context.Context) ([]model.Character, error) {
	rows, err := DB.QueryContext(ctx, "SELECT id, name, reality FROM characters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	characters := make([]model.Character, 0)

	for rows.Next() {
		var character model.Character
		if err := rows.Scan(&character.ID, &character.Name, &character.Reality); err != nil {
			return characters, err
		}
		characters = append(characters, character)
	}

	if err := rows.Err(); err != nil {
		return characters, err
	}

	return characters, nil
}

func GetCharacter(ctx context.Context, characterId int) (model.Character, error) {
	row := DB.QueryRowContext(ctx, "SELECT id, name, reality FROM characters WHERE id = ?", characterId)

	var character model.Character
	if err := row.Scan(&character.ID, &character.Name, &character.Reality); err != nil {
		return character, err
	}
	return character, nil
}

// ユーザーが引いたキャラクターをDBに保存する
func PostUserCharacters(ctx context.Context, selectedCharacters []model.Character, userId int) error {
	for _, selectedCharacter := range selectedCharacters {
		if _, err := DB.QueryContext(ctx, "INSERT INTO userCharacters(userId, characterId) VALUES(?, ?)", userId, selectedCharacter.ID); err != nil {
			return err
		}
	}
	return nil
}

func GetUserCharactersByID(ctx context.Context, userId int) ([]model.UserCharacter, error) {
	rows, err := DB.QueryContext(ctx, "SELECT user_characters.id, user_characters.characterId, characters.name FROM user_characters WHERE userId = ? INNER JOIN characters ON user_characters.characterId = characters.id", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userCharacters := make([]model.UserCharacter, 0)

	for rows.Next() {
		var userCharacter model.UserCharacter
		err := rows.Scan(&userCharacter.UserCharacterID, &userCharacter.CharacterID, &userCharacter.CharacterName)
		if err != nil {
			return nil, err
		}
		userCharacters = append(userCharacters, userCharacter)
		log.Printf("userCharacters🔥🔥🔥: %v", userCharacters)
	}

	if err := rows.Err(); err != nil {
		return userCharacters, err
	}

	return userCharacters, nil
}
