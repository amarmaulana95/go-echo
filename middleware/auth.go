package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/amarmaulana95/go-echo/schemas"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var whiteListPaths = []string{
	"/api/auth/login",
	"/api/auth/register",
}

func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "JWT Token is missing"
}

func WebSecurityConfig(e *echo.Echo) {
	config := middleware.JWTConfig{
		Claims:     &schemas.JWtMetaRequest{},
		SigningKey: []byte("secret"),
		Skipper:    skipAuth,
	}
	e.Use(middleware.JWTWithConfig(config))
}

func skipAuth(e echo.Context) bool {
	path := e.Path()
	for _, p := range whiteListPaths {
		if path == p {
			return true
		}
	}
	return false
}

func RequrieAuth(e echo.Context) {
	// Get the cookie off req
	tokenString := e.Request().Header.Get("Authorization")

	// Decode/validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("screet")), nil
	})

	if token == nil {
		fmt.Println(token.Valid)
		// c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		// 	fmt.Println(sub)

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println(token.Valid)
		}
		// Find the user with token sub

		// Attach to req
		e.Set("user", fmt.Sprint(claims["id"]))

		// Continue

	} else {
		fmt.Println("gagal")
	}
}

func ExtractTokens(s string) string {
	ParsingID := ""
	splitToken := strings.Split(s, "Bearer ")
	token, _, err := new(jwt.Parser).ParseUnverified(splitToken[1], jwt.MapClaims{})
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// obtains claims
		// id  fmt.Sprint(claims["id"])
		ParsingID = fmt.Sprint(claims["id"])
		// fmt.Printf("sub = %s\r\n", id)
		// fmt.Printf("name = %s", name)
	}

	return ParsingID
}

// func ExtractTokens(e echo.Context) (tokenisasi string) {
// 	splitToken := strings.Split(tokenisasi, "Bearer ")
// 	token, _, err := new(jwt.Parser).ParseUnverified(splitToken[1], jwt.MapClaims{})
// 	if err != nil {
// 		fmt.Printf("Error %s", err)
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok {

// 		sub := fmt.Sprint(claims["id"])
// 		name := fmt.Sprint(claims["email"])
// 		fmt.Printf("sub = %s\r\n", sub)
// 		fmt.Printf("email = %s", name)

// 	}

// }
