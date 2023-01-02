package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amarmaulana95/go-echo/entity"
	"github.com/amarmaulana95/go-echo/middleware"
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

func (h *handlerProduct) HandlerResultAll(c echo.Context) error {

	err := c.Request().ParseForm()
	if err != nil {
		fmt.Println("error parsing form", err)
	}
	responData := schemas.ResponStatusDataViewUser{}
	search := c.FormValue("q")
	page_data := "1"
	if len(c.FormValue("_page")) > 0 {
		page_data = c.FormValue("_page")
	}
	page, err := strconv.ParseUint(page_data, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	limit_data := "100"
	if len(c.FormValue("_limit")) > 0 {
		limit_data = c.FormValue("_limit")
	}
	limit, err := strconv.ParseUint(limit_data, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	offset := ((page - 1) * limit)

	res, err := h.products.EntityResultAll(search, limit, offset)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	theTpage := h.products.EntityResultAllTotal(search)

	selisih := theTpage % limit

	total_pages := 1

	if selisih == 0 {
		total_pages = (int(theTpage) / int(limit))
	} else {
		total_pages = (int(theTpage) / int(limit)) + 1
	}

	if err != nil {
		result := schemas.APIResponse("Bad Request", http.StatusUnprocessableEntity, "Error", "Bad Request")
		return c.JSON(http.StatusOK, result)
	}

	tokenString := c.Request().Header.Get("Authorization")
	varID := middleware.ExtractTokens(tokenString)

	fmt.Printf("hasilnya -> %s\r\n", varID)

	responData = schemas.ResponStatusDataViewUser{uint32(page), uint32(limit), uint32(total_pages), uint32(theTpage), res}
	result := schemas.APIResponse("Data user", http.StatusOK, "Success", responData)
	return c.JSON(http.StatusOK, result)

}
