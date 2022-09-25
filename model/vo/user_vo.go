package vo

type UserPostVO struct {
	Account  string `json:"account" example:"帳號" default:"alan" binding:"required"`
	Password string `json:"password" example:"密碼" default:"home_work_z" binding:"required"`
}
