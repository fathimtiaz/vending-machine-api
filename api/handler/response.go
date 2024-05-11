package handler

import (
	"encoding/json"
	"net/http"
	"vending-machine-api/helper"
)

// swagger:response Response
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseWriter struct {
	logger     *helper.Logger
	httpStatus int
	w          http.ResponseWriter
}

func (rw *ResponseWriter) WriteResponse(resp *Response) {
	rw.w.WriteHeader(rw.httpStatus)

	if err := json.NewEncoder(rw.w).Encode(resp); err != nil {
		rw.logger.Error.Printf("error encoding response: %+v", err)
	}
}
