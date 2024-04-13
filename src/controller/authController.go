package controller

import (
	"go-talk/src/database"
	"go-talk/src/models"
	// "go-talk/src/util"
	"log"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// register user
func UserRegisterHandler(c *gin.Context) {

  var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

  log.Println("user Details: ", user)
  user.Password = string(hashedPassword) 
  log.Println("user Details: ", user)

  if err := database.DB.Create(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
    return
  }

  c.JSON(http.StatusOK, map[string]string{"response": "user creation successful"})
}

func UserLoginHandler(c *gin.Context) {

}

func AdminLoginHandler(c *gin.Context) {

}
