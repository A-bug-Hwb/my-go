package pojo

type UserVo struct {
	UserId    int
	Uid       string
	NickName  string
	UserName  string
	Mobile    string
	Mailbox   string
	Password  string
	Signature string
	Introduce string
	Status    string
	Deleted   string
}

type UserBo struct {
	UserId    int
	Uid       string
	NickName  string
	UserName  string
	Mobile    string
	Mailbox   string
	Password  string
	Signature string
	Introduce string
	Status    string
	Deleted   string
	RoleBos   []RoleBo
}

type UserPo struct {
	UserId    int
	Uid       string
	NickName  string
	UserName  string
	Mobile    string
	Mailbox   string
	Password  string
	Signature string
	Introduce string
	Status    string
	Deleted   string
	BaseEntityPo
}

func (UserPo) TableName() string {
	return "user"
}
