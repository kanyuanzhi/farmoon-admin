package request

type CountUsers struct {
	Filter             string `json:"filter" form:"filter"`
	EnableGenderFilter bool   `json:"enableGenderFilter" form:"enableGenderFilter"`
	GenderFilter       string `json:"genderFilter" form:"genderFilter"`
	EnableRolesFilter  bool   `json:"enableRolesFilter" form:"enableRolesFilter"`
	RolesFilter        string `json:"rolesFilter" form:"rolesFilter"`
}

type ListUsers struct {
	PageIndex          int    `json:"pageIndex" form:"pageIndex"`
	PageSize           int    `json:"pageSize" form:"pageSize"`
	Filter             string `json:"filter" form:"filter"`
	EnableGenderFilter bool   `json:"enableGenderFilter" form:"enableGenderFilter"`
	GenderFilter       string `json:"genderFilter" form:"genderFilter"`
	EnableRolesFilter  bool   `json:"enableRolesFilter" form:"enableRolesFilter"`
	RolesFilter        string `json:"rolesFilter" form:"rolesFilter"`
}

type UpdateUser struct {
	Id       uint     `json:"id"`
	Nickname string   `json:"nickname"`
	RealName string   `json:"realName"`
	Gender   string   `json:"gender"`
	Mobile   string   `json:"mobile"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
}

type UpdatePassword struct {
	Id       uint   `json:"id"`
	Password string `json:"password"`
}

type DeleteUsers struct {
	Ids []uint `json:"ids"`
}

type AddUser struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Nickname string   `json:"nickname"`
	RealName string   `json:"realName"`
	Gender   string   `json:"gender"`
	Mobile   string   `json:"mobile"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
}
