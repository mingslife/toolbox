package common

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
)

type errorWrapper struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

var options = []httptransport.ServerOption{
	httptransport.ServerErrorEncoder(func(_ context.Context, err error, w http.ResponseWriter) {
		var (
			statusCode int
			errorCode  uint
			errorText  string
		)
		switch e := err.(type) {
		case BaseError:
			statusCode = e.StatusCode()
			errorCode = e.ErrorCode()
			errorText = e.ErrorText()
		case validator.ValidationErrors:
			statusCode = http.StatusBadRequest
			errorCode = errCodeValidation
			errorText = err.Error()
		default:
			statusCode = http.StatusInternalServerError
			errorCode = errCodeUndefined
			errorText = err.Error()
		}

		log.Printf("ERROR statusCode=%d errorCode=%d errorText=%s\n", statusCode, errorCode, errorText)

		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(&errorWrapper{
			Code:    errorCode,
			Message: errorText,
		})
	}),
}

func NewServer(e endpoint.Endpoint, dec httptransport.DecodeRequestFunc) *httptransport.Server {
	return httptransport.NewServer(e, dec, httptransport.EncodeJSONResponse, options...)
}
