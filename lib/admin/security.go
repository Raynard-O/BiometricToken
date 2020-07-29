package admin

import (
	"BiometricToken/configuration"
	"BiometricToken/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)


var signingKey = configuration.GetHmacSigningKey()

func GenerateToken(admin *models.Admin) (string, error) {
	claims := jwt.MapClaims{
		"email" : admin.Email,
		"fullname"	:	admin.FullName,
		"admin_id"	:	admin.ID,
		"issued_at"	:	time.Now(),
		"expire_at" :	time.Now().Add(time.Minute * 24),
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err  := rawToken.SignedString(signingKey)
	if err != nil {
		//panic("error generating token")
		return "", err
	}
	return token, nil
}
