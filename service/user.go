package service

import (
	"context"
	"database/sql"
	"math/rand"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func CreateToken(userID string) (string, error) {
	// tokenの作成
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// claimsの設定
	token.Claims = jwt.MapClaims{
		"user": userID,
	}

	// 署名
	var secretKey = "secret" // 秘密鍵
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
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
	token, err = CreateToken(id)

	return token, err
}
