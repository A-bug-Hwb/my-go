package result

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ResultDef 自定义返回
func ResultDef(code int, msg string, data interface{}) Result {
	return Result{code, msg, data}
}

// ResultSuc 成功返回
func ResultSuc(msg string, data interface{}) Result {
	// 默认返回数据格式的结构体
	defaultData := &Result{
		Code: SUCCESS,
		Msg:  "ok",
		Data: nil,
	}
	if msg != "" {
		defaultData.Msg = msg
	}
	if data != nil {
		defaultData.Data = data
	}
	return Result{Code: defaultData.Code, Msg: defaultData.Msg, Data: defaultData.Data}
}

// ResultErr 失败返回
func ResultErr(msg string) Result {
	// 默认返回数据格式的结构体
	defaultData := &Result{
		Code: ERROR,
		Msg:  "error",
		Data: nil,
	}
	if msg != "" {
		defaultData.Msg = msg
	}
	return Result{Code: defaultData.Code, Msg: defaultData.Msg, Data: defaultData.Data}
}
