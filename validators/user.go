package validators

type UserValidator struct{}

type Login struct {
  Email string `json:"email" binding:"required"`
  Password string `json:"password" binding:"required"`
}

type Register struct {
  Username    string `form:"username" json:"username" binding:"required,min=4,max=30"`
	Email       string `form:"email" json:"email" binding:"required,email"`
	Password    string `form:"password" json:"password" binding:"required,min=6"`
	ProfilePath *string `form:"profile_path" json:"profile_path"`
}
