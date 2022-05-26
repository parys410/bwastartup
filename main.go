package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=1q2w!Q@W dbname=bwastartup port=5432 sslmode=disable TimeZone=Asia/Makassar"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"status": "Error",
		// 	"message": "Unable to Connect to the Database",
		// })
	}
	fmt.Println("Database is connected successfully.")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users/register", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.Login)
	router.Run()
	/*
	input 			=> dari user
	handler 		=> mapping input dari user -> struct input
	service 		=> melakukan mapping dari struct input ke struct User
	repository	=> terdapat fungsi untuk CRUD ke Database
	db
	*/
}