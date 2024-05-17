package routes

import (
	version1 "blog/interfaces/version1"
	"blog/middleware"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRounter() {
	//In Gin, modes typically control the level of logging among other
	//behaviors, with common modes being "debug", "release", and "test".
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	// 后端管理路由接口
	back := r.Group("interfaces/version1")
	back.Use(middleware.JwtToken())
	{
		// 用户模块路由接口
		back.PUT("user/:id", version1.EditUser)
		back.DELETE("user/:id", version1.DeleteUser)

		// 分类模块路由接口
		back.POST("category/add", version1.AddCategory)
		back.PUT("category/:id", version1.EditCate)
		back.DELETE("category/:id", version1.DeleteCate)

		// 文章模块路由接口
		back.POST("acticle/add", version1.AddArticle)
		back.PUT("artical/:id", version1.EditArt)
		back.DELETE("article/:id", version1.DeleteArt)

		// 上传文件路由接口
		back.POST("upload",version1.UpLoad)

	}

	// 前端展示接口
	front := r.Group("interfaces/version1")
	{
		front.POST("user/add", version1.AddUser)
		front.GET("users", version1.GetUsers)

		front.GET("category", version1.GetCate)

		front.GET("article", version1.GetArt)
		front.GET("article/list/:id", version1.GetCateArt)
		front.GET("article/info/:id", version1.GetArtInfo)

		front.POST("login", version1.Login)
	}

	_ = r.Run(utils.HttpPort)
}
