package main

import (
	"MVC-Go.v2/controllers"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	router := controllers.GetRouter()
	router.LoadHTMLGlob("views/*")

	router.GET("/", controllers.GetHTML)
	router.POST("/", controllers.Login)

	router.Run(":8080")
}
