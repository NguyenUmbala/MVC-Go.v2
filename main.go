package main

import (
	"MVC-Go.v2/controllers"
)

func main() {
	router := controllers.GetRouter()
	router.LoadHTMLGlob("views/*")

	router.GET("/", controllers.GetHTML)
	router.POST("/", controllers.Login)

	router.Run(":8080")
}
