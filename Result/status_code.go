package Result

type ResultData struct {
	Code int
	Msg  string
}

// 成功返回
var SUCCESS = ResultData{
	0,
	"ok",
}

// 失败返回
var ERROR = ResultData{
	500,
	"err",
}

// ip不在白名单返回
var IP_NOT_AUTHORIZED = ResultData{
	403,
	"Ip不在白名单",
}

// token认证失败
var TOKEN_ERR = ResultData{
	401,
	"Token无效",
}
