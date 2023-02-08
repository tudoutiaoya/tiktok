package errno

//func NewHttpErr(statusCode int32, statusMsg string) *controller.Response {
//	return &controller.Response{StatusCode: statusCode, StatusMsg: statusMsg}
//}
//
//func ServiceErr(err error) *controller.Response {
//	return &controller.Response{StatusCode: 50000, StatusMsg: err.Error()}
//}

type ErrResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func NewErrResponse(statusCode int32, statusMsg string) *ErrResponse {
	return &ErrResponse{StatusCode: statusCode, StatusMsg: statusMsg}
}

func HandleServiceErrRes(err error) *ErrResponse {
	return &ErrResponse{StatusCode: 50000, StatusMsg: err.Error()}
}

var (
	ParamIllegalErr = NewErrResponse(50001, "参数不合法")
)
