package Result

type Pagination struct {
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
	Total    int         `json:"total"`
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Data     interface{} `json:"data"`
}

func PageResultSuc(data interface{}, arg ...Handler) Pagination {
	// 默认返回数据格式的结构体
	defaultData := &Result{
		Msg:  "ok",
		Code: 0,
	}
	for _, item := range arg {
		// 接口函数调用闭包函数
		item.parse(defaultData)
	}
	return Pagination{Data: data, Code: defaultData.Code, Msg: defaultData.Msg}
}
