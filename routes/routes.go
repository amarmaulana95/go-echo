package routes

import (
	"net/http"

	"github.com/amarmaulana95/go-echo/handler"
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

	route := router.Group("/api/hello")
	routerAuth := router.Group("/api/auth")
	routerUser := router.Group("/api/users")

	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routerAuth.POST("/login", handlerAuth.HandlerLogin)
	routerAuth.POST("/register", handlerAuth.HandlerRegister)

	routerUser.GET("/hello", handlerUser.HandlerHello)
	// routerUser.POST("/create", handlerUser.HandlerCreate)
	// routerUser.PUT("/update/:id", handlerUser.HandlerUpdate)
	// routerUser.DELETE("/delete/:id", handlerUser.HandlerDelete)
	routerUser.GET("/", handlerUser.HandlerResults)
	// routerUser.GET("/:id", handlerUser.HandlerResult)

}
