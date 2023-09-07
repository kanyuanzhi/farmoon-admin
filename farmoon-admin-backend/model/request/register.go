package request

type Register struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RepeatPassword string `json:"repeatPassword" binding:"required"`
}
