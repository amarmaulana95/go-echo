package handler

import (
	"net/http"

	"github.com/amarmaulana95/go-echo/entity"
	"github.com/amarmaulana95/go-echo/schemas"
	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	user entity.EntityUsers
}

func NewHandlerUser(user entity.EntityUsers) *handlerUser {
	return &handlerUser{
		user: user,
	}
}

func (h *handlerUser) HandlerHello(c echo.Context) error {
	return c.String(200, "Hello World")
}

func (h *handlerUser) HandlerResults(c echo.Context) error {
	_, err := h.user.EntityResults()

	if err != nil {
		result := schemas.APIResponse("Bad Request", http.StatusUnprocessableEntity, "Error", "Bad Request")
		return c.JSON(http.StatusOK, result)
		// c.JSON(400, schemas.SchemaResponses{
		// 	StatusCode: 400,
		// 	Message:    "Bad Request",
		// 	Data:       nil,
		// })

		// return err
	}

	// return c.JSON(200, schemas.SchemaResponses{
	// 	StatusCode: 200,
	// 	Message:    "Success",
	// 	Data:       res,
	// })

	result := schemas.APIResponse("Success", http.StatusOK, "Error", "Success")
	return c.JSON(http.StatusOK, result)
}

// func (h *handlerUser) HandlerCreate(c echo.Context) error {
// 	var body schemas.SchemaUser

// 	if err := c.Bind(&body); err != nil {
// 		c.JSON(400, schemas.SchemaResponses{
// 			StatusCode: 400,
// 			Message:    "Bad Request",
// 			Data:       nil,
// 		})
// 	}

// 	_, error := h.user.EntityCreate(&body)

// 	if error != nil {
// 		c.JSON(400, schemas.SchemaResponses{
// 			StatusCode: 400,
// 			Message:    "Error",
// 			Data:       nil,
// 		})

// 		return error
// 	}

// 	return c.JSON(200, schemas.SchemaResponses{
// 		StatusCode: 200,
// 		Message:    "Success",
// 		Data:       body,
// 	})
// }

// func (h *handlerUser) HandlerResult(c echo.Context) error {
// 	var body schemas.SchemaUser
// 	id := c.Param("id")

// 	body.ID = id

// 	res, err := h.user.EntityResult(&body)

// 	if err != nil {
// 		c.JSON(400, schemas.SchemaResponses{
// 			StatusCode: 400,
// 			Message:    "Bad Request",
// 			Data:       nil,
// 		})

// 		return err
// 	}

// 	return c.JSON(200, schemas.SchemaResponses{
// 		StatusCode: 200,
// 		Message:    "Success",
// 		Data:       res,
// 	})
// }

// func (h *handlerUser) HandlerUpdate(c echo.Context) error {
// 	var body schemas.SchemaUser

// 	if err := c.Bind(&body); err != nil {
// 		c.JSON(400, schemas.SchemaResponses{
// 			StatusCode: 400,
// 			Message:    "Bad Request",
// 			Data:       nil,
// 		})
// 	}

// 	_, error := h.user.EntityUpdate(&body)

// 	if error != nil {
// 		c.JSON(400, schemas.SchemaResponses{
// 			StatusCode: 400,
// 			Message:    "Error",
// 			Data:       nil,
// 		})

// 		return error
// 	}

// 	return c.JSON(200, schemas.SchemaResponses{
// 		StatusCode: 200,
// 		Message:    "Success",
// 		Data:       body,
// 	})
// }

// func (h *handlerUser) HandlerDelete(c echo.Context) error {
// 	var body schemas.SchemaUser
// 	id := c.Param("id")

// 	body.ID = id

// 	_, error := h.user.EntityDelete(&body)

// 	if error != nil {
// 		c.JSON(400, schemas.SchemaResponses{
// 			StatusCode: 400,
// 			Message:    "Error",
// 			Data:       nil,
// 		})

// 		return error
// 	}

// 	return c.JSON(200, schemas.SchemaResponses{
// 		StatusCode: 200,
// 		Message:    "Success",
// 		Data:       body,
// 	})
// }
