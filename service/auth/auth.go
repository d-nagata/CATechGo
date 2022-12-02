package auth

import (
	"github.com/dgrijalva/jwt-go"
)

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

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// jwtの検証
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil // CreateTokenにて指定した文字列を使用
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
