package Result

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Handler interface {
	parse(detail *Result)
}

type HandleFunc func(*Result)

func (f HandleFunc) parse(detail *Result) {
	f(detail)
}

// 闭包函数，当需要修改默认参数时调用
func SetCode(code int) HandleFunc {
	return func(this *Result) {
		this.Code = code
	}
}

func SetMsg(msg string) HandleFunc {
	return func(this *Result) {
		this.Msg = msg
	}
}

// 返回统一数据格式
func AjaxResult(code int, msg string, data interface{}) Result {
	return Result{Data: data, Code: code, Msg: msg}
}

// 返回类
func ResultInfo(resultData ResultData) Result {
	return Result{Data: nil, Code: resultData.Code, Msg: resultData.Msg}
}

func ResultInfoData(resultData ResultData, data interface{}) Result {
	return Result{Data: data, Code: resultData.Code, Msg: resultData.Msg}
}

func ResultSuc(data interface{}, arg ...Handler) Result {
	// 默认返回数据格式的结构体
	defaultData := &Result{
		Msg:  SUCCESS.Msg,
		Code: SUCCESS.Code,
	}
	for _, item := range arg {
		// 接口函数调用闭包函数
		item.parse(defaultData)
	}
	return Result{Data: data, Code: defaultData.Code, Msg: defaultData.Msg}
}

func ResultErr(code int, msg string, arg ...Handler) Result {
	// 默认返回数据格式的结构体
	defaultData := &Result{
		Msg:  ERROR.Msg,
		Code: ERROR.Code,
		Data: nil,
	}
	for _, item := range arg {
		// 接口函数调用闭包函数
		item.parse(defaultData)
	}
	return Result{Data: defaultData.Data, Code: code, Msg: msg}
}
