package routes

import (
	"go-talk/src/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
	// user
	engine.GET("/login", controller.UserLoginHandler)
	engine.GET("/create-post", controller.UserCreatePostHandler)
	engine.GET("/get-online-users", controller.UserGetOnlineFriendsHandler)

	// admin
	engine.GET("/admin-login", controller.AdminLoginHandler)
	engine.GET("/delete-user", controller.AdminDeleteUserHandler)
}
