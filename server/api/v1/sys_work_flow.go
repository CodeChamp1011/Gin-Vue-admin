package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags workflow
// @Summary 注册工作流
// @Produce  application/json
// @Param data body model.SysWorkflow true "注册工作流接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /workflow/createWorkFlow [post]
func CreateWorkFlow(c *gin.Context) {
	var wk model.SysWorkflow
	_ = c.ShouldBindJSON(&wk)
	err := service.Create(wk)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), c)
	} else {
		response.OkWithMessage("获取成功", c)
	}
}
