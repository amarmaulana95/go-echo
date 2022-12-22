package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amarmaulana95/go-echo/entity"
	"github.com/amarmaulana95/go-echo/schemas"
	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	user entity.EntityUsers
}

func NewHandlerUser(user entity.EntityUsers) *handlerUser {
	return &handlerUser{user: user}
}

func (h *handlerUser) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello World")
}

func (h *handlerUser) HandlerResults(c echo.Context) error {
	res, err := h.user.EntityResults()

	if err != nil {
		err := schemas.APIResponse("Bad Request", http.StatusBadGateway, "Error", "Bad Request")
		return c.JSON(http.StatusBadGateway, err)

	}

	result := schemas.APIResponse("Data user", http.StatusOK, "success", res)
	return c.JSON(http.StatusOK, result)
}

func (h *handlerUser) HandlerResultService(c echo.Context) error {
	var body schemas.SchemaUser
	id := c.Param("id")

	body.ID = id

	res, err := h.user.EntityResultID(&body)

	if err != nil {
		result := schemas.APIResponse("Bad Request", http.StatusUnprocessableEntity, "Error", "Bad Request")
		return c.JSON(http.StatusOK, result)
	}

	result := schemas.APIResponse("Ok", http.StatusUnprocessableEntity, "Success", res)
	return c.JSON(http.StatusOK, result)
}

func (h *handlerUser) HandlerResultSearch(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		fmt.Println("error parsing form", err)
	}
	search := c.FormValue("q")
	res, err := h.user.EntityResultIDSearch(search)
	result := schemas.APIResponse("Data user", http.StatusOK, "Success", res)
	return c.JSON(http.StatusOK, result)

}

func (h *handlerUser) HandlerResultAll(c echo.Context) error {
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

	res, err := h.user.EntityResultAll(search, limit, offset)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	theTpage := h.user.EntityResultAllTotal(search)

	fmt.Println(theTpage)

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

	responData = schemas.ResponStatusDataViewUser{uint32(page), uint32(limit), uint32(total_pages), uint32(theTpage), res}
	return c.JSON(http.StatusOK, responData)

}
