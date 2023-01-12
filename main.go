package main

import (
	"crowdfounding/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "root:@tcp(127.0.0.1:3306)/crowdfounding-golang?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("connection database success")

	// var users []user.User

	// length := len(users)

	// fmt.Println(length)

	// db.Find(&users)
	// length = len(users)
	// fmt.Println(length)

	// for _, user := range users {
	// 	fmt.Println(user.Email)
	// 	fmt.Println(user.Name)

	router := gin.Default()
	router.GET("/user", getUser)
	router.Run()

}

func getUser(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/crowdfounding-golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
