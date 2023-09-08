package response

type Login struct {
	Avatar   string   `json:"avatar"`
	Username string   `json:"username"`
	Nickname string   `json:"nickname"`
	RealName string   `json:"realName"`
	Gender   string   `json:"gender"`
	Mobile   string   `json:"mobile"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
	Token    string   `json:"token"`
}
