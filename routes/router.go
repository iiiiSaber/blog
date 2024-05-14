package routes

import (
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRounter(){
	gin.SetMode(utils.AppMode)
	r :=gin.Default()

	router := r.Group("interfaces/version1.0")
	{
		router.GET("hello",func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK,gin.H{
				"msg":"OK",
			})
		})
	}
	r.Run(utils.HttpPort)
}