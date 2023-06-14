package pojo

type WrMenuBo struct {
	Id       int    `json:"id"`
	MenuName string `json:"menuName"`
	OrderNum string `json:"orderNum"`
	Path     string `json:"path"`
	Visible  string `json:"visible"`
	Status   string `json:"status"`
	BaseEntityBo
}

func (WrMenuBo) TableName() string {
	return "wr_menu"
}

type WrMenuVo struct {
	Id       int    `json:"id"`
	MenuName string `json:"menuName"`
	OrderNum string `json:"orderNum"`
	Path     string `json:"path"`
	Visible  string `json:"visible"`
	Status   string `json:"status"`
}
