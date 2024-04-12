package routes

import (
	"go-talk/src/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {

	// test
	engine.GET("/ping", controller.PONG)

	// user
	engine.POST("/register", controller.UserRegisterHandler)
	engine.POST("/login", controller.UserLoginHandler)
	engine.POST("/create-post", controller.UserCreatePostHandler)
	engine.GET("/get-online-users", controller.UserGetOnlineFriendsHandler)

	// admin
	engine.POST("/admin-login", controller.AdminLoginHandler)
	engine.GET("/delete-user", controller.AdminDeleteUserHandler)
}
