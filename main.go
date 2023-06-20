package main

import (
	"log"
	"net/http"
	"rest-crowdfounding/auth"
	"rest-crowdfounding/campaign"
	"rest-crowdfounding/handler"
	"rest-crowdfounding/helper"
	"rest-crowdfounding/user"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfounding-golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()

	input := campaign.CreateCampaignInput{}
	input.Name = "Penggalangan dana untuk disfable"
	input.ShortDescription = "shorttt"
	input.Description = "longggg"
	input.GoalAmount = 200000
	input.Perks = "hadiah satuku, duaku, tigaku"
	inputUser, _ := userService.GetUserByID(1)
	input.User = inputUser

	_, err = campaignService.CreateCampaign(input)

	if err != nil {
		log.Fatal(err.Error())
	}

	userHandler := handler.NewUserHandler(userService, authService)

	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	api.POST("/sessions", userHandler.Login)

	api.POST("/email_checkers", userHandler.CheckEmailAvailable)

	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploudAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)

	api.GET("/campaigns/:id", campaignHandler.GetCampaign)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""

		arrayOfToken := strings.Split(authHeader, " ")

		if len(tokenString) == 2 {
			tokenString = arrayOfToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}

}
