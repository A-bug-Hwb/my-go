package pojo

type RoleVo struct {
	roleId   int
	roleName string
	roleKey  string
	sort     int
	status   int
}

type RoleBo struct {
	roleId   int
	roleName string
	roleKey  string
	sort     int
	status   int
}

type RolePo struct {
	roleId   int
	roleName string
	roleKey  string
	sort     int
	status   int
}

func (RolePo) TableName() string {
	return "role"
}
