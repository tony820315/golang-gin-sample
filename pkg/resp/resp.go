package resp

import (
	"net/http"
)

type ResponseBody struct {
	Code    int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"data"`
}

func NewResponseBody(extendError ExtendError) ResponseBody {
	return ResponseBody{Code: extendError.GetCode(), Message: extendError.GetMessage()}
}

func (r *ResponseBody) SetExtendError(extendError ExtendError) {
	r.Code = extendError.GetCode()
	r.Message = extendError.GetMessage()
}

func (r *ResponseBody) StatusCode() int {
	statusCode := 500
	if http.StatusContinue <= r.Code && r.Code <= http.StatusNetworkAuthenticationRequired {
		statusCode = r.Code
	}

	return statusCode
}
