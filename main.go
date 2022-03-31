package main

import (
	"fmt"
	"log"

	"go-es/config"
	"go-es/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	client, err := config.CreateElasticConnection()
	if err != nil {
		log.Println("Connection failed : ", err.Error())
	}
	config.Client = client
	fmt.Println("config client: ", config.Client)
	router := gin.Default()
	router.GET("/users/:id", controller.FindUserById)
	router.POST("/users", controller.SaveUser)
	router.PUT("/users", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
	router.Run(":8080")

}
