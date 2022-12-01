package service

import "database/sql"

type GachaService struct {
	db *sql.DB
}

func NewGachaService(db *sql.DB) *GachaService {
	return &GachaService{
		db: db,
	}
}
