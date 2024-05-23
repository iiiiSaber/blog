package routes

import (
	version1 "blog/interfaces/version1"
	"blog/middleware"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRounter(){
	gin.SetMode(utils.AppMode)
	r :=gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	// 后端管理路由接口
	back := r.Group("interfaces/version1")
	back.Use(middleware.JwtToken())
	{
		// 用户模块路由接口
		back.GET("admin/users",version1.GetUsers)
		back.PUT("user/:id",version1.EditUser)
		back.DELETE("user/:id",version1.DeleteUser)
		// 修改密码
		back.PUT("admin/changePassword/:id",version1.ChangeUserPassword)

		// 分类模块路由接口
		back.GET("admin/category", version1.GetCate)
		back.POST("category/add",version1.AddCategory)
		back.PUT("category/:id",version1.EditCate)
		back.DELETE("category/:id",version1.DeleteCate)

		// 文章模块路由接口
		back.GET("admin/article/info/:id", version1.GetArtInfo)
		back.GET("admin/article", version1.GetArt)
		back.POST("article/add",version1.AddArticle)
		back.PUT("article/:id",version1.EditArt)
		back.DELETE("article/:id",version1.DeleteArt)

		// 上传文件路由接口
		back.POST("upload",version1.UpLoad)

	}

	// 前端展示接口
	front := r.Group("interfaces/version1")
	{
		front.POST("user/add",version1.AddUser)
		front.GET("users",version1.GetUsers)
		front.GET("user/:id",version1.GetUserInfo)

		front.GET("category",version1.GetCate)
		front.GET("category/:id", version1.GetCateInfo)

		front.GET("article",version1.GetArt)
		front.GET("article/list/:id",version1.GetCateArt)
		front.GET("article/info/:id",version1.GetArtInfo)

		front.POST("login",version1.Login)
	}
	
	_ = r.Run(utils.HttpPort)
}