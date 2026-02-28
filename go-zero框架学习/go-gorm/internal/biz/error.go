package biz

type ErrorUtil struct {
	Code int
	Msg  string
}

func NewError(code int, msg string) *ErrorUtil {
	return &ErrorUtil{Code: code, Msg: msg}
}
