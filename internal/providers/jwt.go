package providers

import (
	"fmt"
	"time"

	"github.com/donkeysharp/time-to-get-a-job-backend/internal/domain/models"
	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {
	Secret string
}

type ApplicationClaims struct {
	jwt.RegisteredClaims
}

func (me *JWTProvider) CreateJWT(account *models.Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%v", account.Id),
	})

	t, err := token.SignedString([]byte(me.Secret))
	if err != nil {
		return "", nil
	}
	return t, nil
}

func (me *JWTProvider) ValidateJWT(token string) error {
	return nil
}
