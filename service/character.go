package service

import "database/sql"

type CharacterService struct {
	db *sql.DB
}

func NewCharacterService(db *sql.DB) *CharacterService {
	return &CharacterService{
		db: db,
	}
}
