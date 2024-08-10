package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var secretKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

}

type Claims struct {
	ID       int
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(id int, username string) (string, error) {
	exp := time.Now().Add(time.Hour)
	claims := Claims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyToken(token string) (*Claims, error) {
	claims := &Claims{}
	t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, err
	}

	return claims, nil
}
