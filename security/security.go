package security

import (
	client_secret "final-project/client/secret"
	server_secret "final-project/server/secret"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(fields map[string]string) (string, error) {
	var claims jwt.MapClaims
	for k, v := range fields {
		claims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(server_secret.SECRET_TOKEN)
	return tokenStr, err
}

func ParseToken(tokenStr string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(client_secret.SECRET_TOKEN), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
