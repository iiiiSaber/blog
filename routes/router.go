package routes

import (
	version1 "blog/interfaces/version1"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRounter(){
	gin.SetMode(utils.AppMode)
	r :=gin.Default()

	router := r.Group("interfaces/version1")
	{
		// 用户模块路由接口
		router.POST("user/add",version1.AddUser)
		router.GET("users",version1.GetUsers)
		router.PUT("user/:id",version1.EditUser)
		router.DELETE("user/:id",version1.DeleteUser)

		// 分类模块路由接口
		router.POST("category/add",version1.AddCategory)
		router.GET("category",version1.GetCate)
		router.PUT("category/:id",version1.EditCate)
		router.DELETE("category/:id",version1.DeleteCate)

		// 文章模块路由接口


	}
	r.Run(utils.HttpPort)
}