package httputil

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
