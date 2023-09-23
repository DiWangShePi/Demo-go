package router

import (
	"Demo/middleware"
	"Demo/services"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (userRouter UserRouter) InitUserRouter(Router *gin.Engine) {
	initUserRouter := Router.Group("user").Use(middleware.JWTAuth())
	{
		initUserRouter.POST("update", services.UserUpdateUser)
		initUserRouter.POST("delete", services.UserDeleteUser)
		initUserRouter.GET("find", services.UserFindUser)
		initUserRouter.POST("create", services.UserCreateUser)
	}
	publicUserRouter := Router.Group("puser")
	{
		publicUserRouter.POST("login", services.UserLogin)
	}
}
