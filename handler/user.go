package handler

import (
	"crowdfounding/helper"
	"crowdfounding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.Apiresponse("Account failed register", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")

	response := helper.Apiresponse("Account has been success register", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
