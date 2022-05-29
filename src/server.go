package main

import (
	Controller "github.com/BenjaminLee-boop/golang-OAuthServer/src/controller"
	"github.com/BenjaminLee-boop/golang-OAuthServer/src/db"
	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.Init()
	h := Controller.New(DB)
	r := gin.Default()

	r.GET("/", h.GetAllUsers)
	r.GET("/{id}", h.GetUser)
	r.POST("/", h.AddUser)
	r.PUT("/{id}", h.UpdateUser)
	r.DELETE("/{id}", h.DeleteUser)

	r.Run(":9001")
}
