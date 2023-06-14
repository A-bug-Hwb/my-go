package pojo

import (
	"github.com/gogf/gf/os/gtime"
)

type BaseEntityBo struct {
	CreateBy   string      `json:"createBy"`
	CreateTime *gtime.Time `json:"createTime"`
	UpdateBy   string      `json:"updateBy"`
	UpdateTime *gtime.Time `json:"updateTime"`
	Remark     string      `json:"remark"`
}

type BaseEntityVo struct {
	CreateBy   string      `json:"createBy"`
	CreateTime *gtime.Time `json:"createTime"`
	UpdateBy   string      `json:"updateBy"`
	UpdateTime *gtime.Time `json:"updateTime"`
	Remark     string      `json:"remark"`
}
