package model

type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Reality int    `json:"reality"`
}

type UserCharacter struct {
	UserCharacterID int    `json:"id"`
	CharacterID     int    `json:"characterId"`
	CharacterName   string `json:"characterName"`
}
