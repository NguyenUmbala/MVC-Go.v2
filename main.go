package main

import (
	"MVC-Go.v2/controllers"
)

func main() {
	router := controllers.GetRouter()
	router.LoadHTMLGlob("views/*")

	router.GET("/", controllers.GetHTML)
	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.SignUp)

	router.Run(":8080")
}
