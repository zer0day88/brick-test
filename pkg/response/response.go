package response

import (
	"github.com/labstack/echo/v4"
	pkgerr "github.com/zer0day88/brick-test/pkg/error"
)

type ResponseHeader struct {
	HttpCode int
}

type ResponseBody struct {
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	Data            interface{} `json:"data,omitempty"`
}

type Response struct {
	Header ResponseHeader
	Body   ResponseBody
}

func (r *Response) JSON(c echo.Context, data interface{}, err pkgerr.CustomErr) error {
	r.Header.HttpCode = err.HttpCode
	r.Body = ResponseBody{
		ResponseCode:    err.Code,
		ResponseMessage: err.Msg,
	}

	if data != nil {
		r.Body.Data = data
	}

	return c.JSON(r.Header.HttpCode, r.Body)
}
