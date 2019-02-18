package controllers

import (
	"net/http"

	"MVC-Go.v2/models"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func GetRouter() *gin.Engine {
	return router
}

func GetHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"noti": "Load trang thành công",
	})
}

func Login(c *gin.Context) {
	models.OpenConnection()
	defer models.CloseConnection()

	var acc models.Account
	c.BindJSON(&acc)

	if models.Login(acc) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"noti": "Đăng nhập thành công",
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"noti": "Đăng nhập thất bại",
		})
	}
}
