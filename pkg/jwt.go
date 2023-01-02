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

// func ValidateToken(tokenString string) (*jwt.Token, error) {

// 	splitToken := strings.Split(tokenString, "Bearer ")

// 	fmt.Println(splitToken[1])

// 	token, _, err := new(jwt.Parser).ParseUnverified(splitToken[1], jwt.MapClaims{})
// 	if err != nil {
// 		fmt.Printf("Error %s", err)
// 	}
// 	claims := token.Claims.(jwt.MapClaims)
// 	sub := fmt.Sprint(claims["id"])
// 	fmt.Println(sub)

// 	// token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
// 	// 	_, ok := token.Method.(*jwt.SigningMethodHMAC)
// 	// 	//cek tokennya
// 	// 	if !ok {
// 	// 		return nil, errors.New("invalid token")
// 	// 	}

// 	// 	return []byte("screet"), nil
// 	// })

// 	// if err != nil {
// 	// 	return token, err
// 	// }
// 	return token, err

// }
