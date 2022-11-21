package handler

import (
	"net/http"

	"github.com/amarmaulana95/go-echo/entity"
	"github.com/amarmaulana95/go-echo/pkg"
	"github.com/amarmaulana95/go-echo/schemas"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	auth entity.EntityAuth
}

func NewHandlerAuth(auth entity.EntityAuth) *handlerAuth {
	return &handlerAuth{auth: auth}
}

func (h *handlerAuth) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello Auth")
}

func (h *handlerAuth) HandlerRegister(c echo.Context) error {
	var body schemas.SchemaUser

	if err := c.Bind(&body); err != nil {
		// return c.JSON(400, schemas.SchemaResponses{
		// 	StatusCode: 400,
		// 	Message:    "Bad Request",
		// 	Data:       nil,
		// })
		result := schemas.APIResponse("Bad Request", http.StatusUnprocessableEntity, "Error", "Bad Request")
		return c.JSON(http.StatusOK, result)
	}

	_, error := h.auth.EntityRegister(&body)

	if error != nil {
		result := schemas.APIResponse("Bad Request", http.StatusUnprocessableEntity, "Error", "Bad Request")
		return c.JSON(http.StatusOK, result)
	}

	// return c.JSON(200, schemas.SchemaResponses{
	// 	StatusCode: 200,
	// 	Message:    "Success",
	// 	Data:       body,
	// })
	result := schemas.APIResponse("Ok", http.StatusOK, "success", body)
	return c.JSON(http.StatusOK, result)
}

func (h *handlerAuth) HandlerLogin(c echo.Context) error {
	var body schemas.SchemaUser

	if err := c.Bind(&body); err != nil {
		result := schemas.APIResponse("Error Authtentication", http.StatusUnprocessableEntity, "Error", "email atau password salah")
		return c.JSON(http.StatusOK, result)
	}

	res, err := h.auth.EntityLogin(&body)

	if err != nil {
		result := schemas.APIResponse("Error Authtentication", http.StatusUnprocessableEntity, "Error", "email atau password salah")
		return c.JSON(http.StatusOK, result)
	}

	jwt, err := pkg.Sign(&schemas.SchemaJWT{
		ID:    res.ID,
		Email: res.Email,
	})

	if err != nil {
		result := schemas.APIResponse("Error Authtentication", http.StatusUnprocessableEntity, "Error", "email atau password salah")
		return c.JSON(http.StatusOK, result)
	}

	result := schemas.APIResponse("Success login", http.StatusOK, "success", jwt)
	return c.JSON(http.StatusOK, result)
}
