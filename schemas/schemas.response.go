package schemas

import "github.com/go-playground/validator/v10"

// type SchemaResponses struct {
// 	StatusCode int         `json:"code"`
// 	Message    string      `json:"message"`
// 	Data       interface{} `json:"data"`
// }

type SchemaResponses struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) SchemaResponses {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := SchemaResponses{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
