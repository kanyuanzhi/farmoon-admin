package response

type Login struct {
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	RealName string `json:"realName"`
	Token    string `json:"token"`
}
