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

type UserPossessionCharacter struct {
	CharacterID     int    `json:"id"`
	CharacterReality int   `json:"reality"`
	CharacterName   string `json:"name"`
	CharacterSum	int	   `json:"sum"`
}
