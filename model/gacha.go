package model

type (
	GachaResult struct {
		CharacterID string `json:"characterID"`
		Name        string `json:"name"`
	}
	GachaDrawRequest struct {
		Times int `json:"times"`
	}
	GachaDrawResponse struct {
		Results GachaResult `json:"results"`
	}
)
