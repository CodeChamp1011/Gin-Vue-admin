package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/notify/api"
	"github.com/gin-gonic/gin"
)

type NotifyRouter struct {
}

func (s *NotifyRouter) InitRouter(Router *gin.RouterGroup) {
	router := Router.Use(middleware.OperationRecord())
	var SendTextMessage = api.ApiGroupApp.Api.SendTextMessage
	{
		router.POST("sendTextMessage", SendTextMessage)
	}
}
