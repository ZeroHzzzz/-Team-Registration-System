package utils

// import "errors"

// var (
// 	ErrParameterWrong  = errors.New("参数错误")
// 	ErrInternalServer  = errors.New("网络错误")
// 	ErrOperationFailed = errors.New("操作未成功完成")
// 	ErrFormatWrong     = errors.New("格式错误")
// )

const (
	SuccessCode             = 200
	ParameterErrorCode      = 400
	OperationFailedCode     = 401
	NotFoundCode            = 402
	FormatWrongCode         = 405
	HaveExistCode           = 406
	InternalServerErrorCode = 500
)

const (
	SuccessMsg             = "Accepted"
	ParameterErrorMsg      = "参数错误"
	InternalServerErrorMsg = "Internal server error"
	OperationFailedMsg     = "操作未成功完成"
	UserNotFoundMsg        = "账户不存在或密码错误"
	TeamNotFoundFailMsg    = "团队不存在或密码错误"
	FormatWrongMsg         = "格式错误"
	UserHaveExistMsg       = "邮箱已被注册"
	NotFoundMsg            = "找不到对象"
)
