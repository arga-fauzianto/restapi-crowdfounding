package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Ocuppation string `json:"ocuppation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}
