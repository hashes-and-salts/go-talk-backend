package controller

import (
	"go-talk/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(c *gin.Context) {

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

func UserLoginHandler(c *gin.Context) {

}

func UserCreatePostHandler(c *gin.Context) {

}

func UserGetOnlineFriendsHandler(c *gin.Context) {

}

func AdminLoginHandler(c *gin.Context) {

}

func AdminDeleteUserHandler(c *gin.Context) {

}

func PONG(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
