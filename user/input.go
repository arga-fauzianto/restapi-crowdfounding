package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required, email"`
	Ocuppation string `json:"ocuppation" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
