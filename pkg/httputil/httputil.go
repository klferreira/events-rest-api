package httputil

import (
	"encoding/json"
	"net/http"
)

type JSONResponse map[string]interface{}

func GetJSONResponse(keyname string, data interface{}, err error) *JSONResponse {
	response := JSONResponse{
		"data": JSONResponse{keyname: data},
	}

	if err != nil {
		response["error"] = err.Error()
	}

	return &response
}

func (r *JSONResponse) Write(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(r)
}
