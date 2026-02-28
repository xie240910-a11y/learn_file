package biz

import "go-gorm/internal/types"

// 成功的函数
func Success(data interface{}, message ...string) *types.CommonResponse {
	if len(message) > 0 {
		return &types.CommonResponse{
			Code:    OK,
			Message: message[0],
			Success: true,
			Data:    data,
		}
	}
	return &types.CommonResponse{
		Code:    OK,
		Success: true,
		Data:    data,
	}
}

// 失败的函数
func Error(err *ErrorUtil) *types.CommonResponse {
	return &types.CommonResponse{
		Code:    err.Code,
		Message: err.Msg,
		Success: false,
	}
}
