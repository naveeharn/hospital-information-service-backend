package helper

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Errors  any    `json:"errors"`
	Data    any    `json:"data"`
	Result  Result `json:"result"`
}

type Result struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type EmptyObj struct{}

func CreateResponse(data any, result Result) Response {
	return Response{
		Errors:  nil,
		Data:    data,
		Result:  result,
	}
}

func CreateErrorResponse(err string, data any, result Result) Response {
	return Response{
		Status:  false,
		Errors:  err,
		Data:    data,
		Result:  result,
	}
}
