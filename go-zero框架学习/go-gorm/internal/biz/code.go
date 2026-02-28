package biz

const OK = 200

var (
	DBError      = NewError(10000, "数据库不存在")
	ParamError   = NewError(10001, "参数不存在")
	DataNotError = NewError(10002, "数据不存在")
	ServerError  = NewError(10003, "服务器不存在")
)
