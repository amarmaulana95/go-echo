package pkg

import (
	"time"

	"github.com/amarmaulana95/go-echo/schemas"
	"github.com/golang-jwt/jwt"
)

func Sign(schemajwt *schemas.SchemaJWT) (string, error) {
	claims := &schemas.JWtMetaRequest{
		ID:    schemajwt.ID,
		Email: schemajwt.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := token.SignedString([]byte("secret"))

	return jwt, err
}
