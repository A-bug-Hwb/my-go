package pojo

type UserRolePo struct {
	RoleId int
	Userid int
}

func (UserRolePo) TableName() string {
	return "user_role"

}
