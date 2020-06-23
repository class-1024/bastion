package errno

type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

/*
错误码设计
第一位表示错误级别, 1 为系统错误, 2 为普通错误
第二三位表示服务模块代码
第四五位表示具体错误代码
*/
var (
	OK = &Errno{Code: 0, Message: "Success"}

	// 系统错误, 前缀为 100
	InternalServerError = &Errno{Code: 10001, Message: "服务器开小差"}
	InvalidParams       = &Errno{Code: 10002, Message: "请求参数错误"}

	ErrorSession    = &Errno{Code: 10003, Message: "session 发送错误"}
	ErrorIpNotAllow = &Errno{Code: 10004, Message: "ip 不在白名单"}
	ErrorNotFound   = &Errno{Code: 10005, Message: "未找到"}
	ErrorCreateData = &Errno{Code: 10006, Message: "创建错误"}
	ErrorQueryData  = &Errno{Code: 10007, Message: "查询错误"}
	ErrorUpdateData = &Errno{Code: 10007, Message: "更新错误"}

	// 用户错误, 前缀为 203
	ErrorUserNotLogin    = &Errno{Code: 20301, Message: "用户未登录"}
	ErrorUsePassword     = &Errno{Code: 20302, Message: "密码错误"}
	ErrorUserNotFound    = &Errno{Code: 20302, Message: "密码错误"}
	ErrorDecryptUserData = &Errno{Code: 20303, Message: "微信用户数据解密失败"}
)
