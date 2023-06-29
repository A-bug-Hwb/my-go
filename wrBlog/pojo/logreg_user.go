package pojo

type LoginUserVo struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginUser struct {
	UserId        int
	Token         string
	LoginTime     string
	ExpireTime    string
	Ipaddr        string
	LoginLocation string
	Browser       string
	Os            string
	Permissions   []string
	UserBo
}

type RegisterUser struct {
	UserName        string `json:"userName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
