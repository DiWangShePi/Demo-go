package router

import (
	"Demo/middleware"
	"Demo/services"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (adminRouter AdminRouter) InitAdminRouter(Router *gin.Engine) {
	initAdminRouter := Router.Group("admin").Use(middleware.JWTAuth())
	{
		initAdminRouter.POST("update", services.AdminUpdateAdmin)
		initAdminRouter.POST("createUser", services.AdminCreateUser)
		// initAdminRouter.POST("updateUser", services.AdminUpdateUser)
		initAdminRouter.POST("deleteUser", services.AdminDeleteUser)
		initAdminRouter.GET("findUser", services.AdminFindUser)
		initAdminRouter.GET("findUserList", services.AdminFindUserList)
	}
	publicAdminRouter := Router.Group("padmin")
	{
		publicAdminRouter.POST("login", services.AdminLogin)
	}
}
