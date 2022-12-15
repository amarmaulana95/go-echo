package handler

import (
	"net/http"

	"github.com/amarmaulana95/go-echo/entity"
	"github.com/amarmaulana95/go-echo/schemas"
	"github.com/labstack/echo/v4"
)

type handlerProduct struct {
	products entity.EntityProducts
}

func NewHandlerProduct(products entity.EntityProducts) *handlerProduct {
	return &handlerProduct{products: products}
}

func (h *handlerProduct) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello World")
}

func (h *handlerProduct) HandlerResults(c echo.Context) error {

	res, err := h.products.EntityResults()
	if err != nil {
		err := schemas.APIResponse("Bad Request", http.StatusBadGateway, "Error", "Bad Request")
		return c.JSON(http.StatusBadGateway, err)
	}
	result := schemas.APIResponse("Data Products", http.StatusOK, "success", res)
	return c.JSON(http.StatusOK, result)
}
