package controllers

type ErrorResponse struct {
	Msg string
	Err error
}

type SuccessResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}
