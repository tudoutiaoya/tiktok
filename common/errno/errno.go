package errno

type ErrResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// NewErrResponse 新建错误响应
func NewErrResponse(statusCode int32, statusMsg string) *ErrResponse {
	return &ErrResponse{StatusCode: statusCode, StatusMsg: statusMsg}
}

// HandleServiceErrRes 处理service层返回的错误
func HandleServiceErrRes(err error) *ErrResponse {
	return &ErrResponse{StatusCode: 50000, StatusMsg: err.Error()}
}

// 鉴权失败，需要登录
var NoLoginErr = NewErrResponse(40100, "请登录")

// Controller层一些参数错误
var (
	ParamIllegal      = NewErrResponse(50001, "参数不合法")
	LoginParamIllegal = NewErrResponse(50003, "用户名或密码格式不正确")
	TokenIllegal      = NewErrResponse(50004, "token非法或已过期")
)

// Service
var ()
