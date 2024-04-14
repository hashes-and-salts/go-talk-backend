package controller

import (
	"errors"
	"go-talk/src/database"
	"go-talk/src/models"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// register user
func UserRegisterHandler(c *gin.Context) {

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "user creation successful"})
}

func UserLoginHandler(c *gin.Context) {

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	queryResult := database.DB.First(&user)
	if queryResult.RowsAffected == 0 || errors.Is(queryResult.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{"message": "invalid credentials"})
		return
	}

	// jwt if found
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    user.Email,
		"username": user.Username,
	})

	tokenString, err := claims.SignedString([]byte(os.Getenv("JWT_SIGN_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not sign jwt key"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

	c.SetCookie(
		"token",
		tokenString,
		10,
		"/",
		"localhost",
		true,
		true,
	)

}

func AdminLoginHandler(c *gin.Context) {

}
