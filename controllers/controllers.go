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

func SignUp(c *gin.Context) {
	models.OpenConnection()
	defer models.CloseConnection()

	person := models.Person{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}

	models.SignUp(person)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"noti": "Đăng kí thành công",
	})
}

func Login(c *gin.Context) {
	models.OpenConnection()
	defer models.CloseConnection()

	person := models.Person{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}

	if models.Login(person) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"noti": "Đăng nhập thành công",
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"noti": "Đăng nhập thất bại",
		})
	}
}
