package routes

import (
	"net/http"

	"github.com/amarmaulana95/go-echo/handler"
	"github.com/amarmaulana95/go-echo/middleware"
	"github.com/amarmaulana95/go-echo/repository"
	"github.com/amarmaulana95/go-echo/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(db *gorm.DB, router *echo.Echo) {

	repositoryAuth := repository.NewRepositoryAuth(db)
	serviceAuth := services.NewServiceAuth(repositoryAuth)
	handlerAuth := handler.NewHandlerAuth(serviceAuth)

	repositoryUser := repository.NewRepositoryUser(db)
	serviceUser := services.NewServiceUser(repositoryUser)
	handlerUser := handler.NewHandlerUser(serviceUser)

	repositoryProduct := repository.NewRepositoryProduct(db)
	serviceProduct := services.NewserviceProduct(repositoryProduct)
	handlerProduct := handler.NewHandlerProduct(serviceProduct)

	route := router.Group("/api/hello")
	routerAuth := router.Group("/api/auth")
	routerUser := router.Group("/api/users")
	routerProducts := router.Group("/api/products")

	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routerAuth.POST("/login", handlerAuth.HandlerLogin)
	routerAuth.POST("/register", handlerAuth.HandlerRegister)

	middleware.WebSecurityConfig(router)

	routerUser.GET("/hello", handlerUser.HandlerHello)
	routerUser.GET("/users", handlerUser.HandlerResults)
	routerUser.GET("/userall", handlerUser.HandlerResultAll)
	routerUser.GET("/users/:id", handlerUser.HandlerResultService)
	routerUser.GET("/cari", handlerUser.HandlerResultSearch)

	routerProducts.GET("/product", handlerProduct.HandlerResults)
	routerProducts.GET("/products", handlerProduct.HandlerResultAll)
	// routerProducts.GET("/productsID", handlerProduct.HandlerProductSearch)

}

// func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
// 	var (
// 		jwtKey = os.Getenv("JWT_KEY")
// 	)

// 	return func(c echo.Context) error {
// 		authToken := c.Request().Header.Get("Authorization")
// 		if authToken == "" {
// 			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
// 		}

// 		splitToken := strings.Split(authToken, "Bearer ")
// 		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
// 			}

// 			return []byte(jwtKey), nil
// 		})

// 		if !token.Valid || err != nil {
// 			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
// 		}

// 		var id int
// 		destructID := token.Claims.(jwt.MapClaims)["id"]
// 		if destructID != nil {
// 			id = int(destructID.(float64))
// 		} else {
// 			id = 0
// 		}
// 		return id
// 	}
// }

// token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
// 	if err != nil {
// 		fmt.Printf("Error %s", err)
// 	}
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 		// obtains claims
// 		sub := fmt.Sprint(claims["sub"])
// 		name := fmt.Sprint(claims["name"])
// 		// and so on and on
// 		// ...
// 		fmt.Printf("sub = %s\r\n", sub)
// 		fmt.Printf("name = %s", name)

// 	}
