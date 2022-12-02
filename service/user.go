package service

import (
	"context"
	"database/sql"
	"math/rand"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ikalemmon/CATechGo/service/auth"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name string) (string, error) {
	const (
		insert = `INSERT INTO users(name, id) VALUES (?,?)`
	)
	var token = ""

	//現在時刻からid作成
	rand.Seed(time.Now().UnixNano())
	var id = strconv.Itoa(rand.Intn(10000000000))

	_, err := s.db.ExecContext(ctx, insert, name, id)
	if err != nil {
		return "", err
	}

	//トークン作成
	token, err = auth.CreateToken(id)

	return token, err
}

func (s *UserService) GetUser(ctx context.Context, tokenString string) (string, error) {
	const (
		read = `SELECT name FROM users where id = ?`
	)
	var name = ""

	//トークン読み取り
	token, err := auth.VerifyToken(tokenString)
	if err != nil {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"]

	stmt, err := s.db.PrepareContext(ctx, read)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	row, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return "", err
	}
	defer row.Close()

	if err := row.Scan(name); err != nil {
		return "", err
	}

	return name, err
}
