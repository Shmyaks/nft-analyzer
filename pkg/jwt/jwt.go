package jwt

import (
	db "golang-template-module/pkg/database"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtCustomClaims struct {
	Id    uint64 `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func Create(m *db.User) (string, error) {
	claims := &jwtCustomClaims{
		m.Id,
		m.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return t, err
}
