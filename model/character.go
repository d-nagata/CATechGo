package model

type (
	UserCharacter struct {
		UserCharacterID string `json:"userCharacterID"`
		CharacterID     string `json:"characterID"`
		Name            string `json:"name"`
	}

	CharacterListResponse struct {
		Characters []UserCharacter `json:"characters"`
	}
)
