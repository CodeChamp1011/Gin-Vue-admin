package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/url"
	"os"
)

// @Tags AutoCode
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// TODO @Param data body model.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func CreateTemp(c *gin.Context) {
	var a model.AutoCodeStruct
	_ = c.ShouldBindJSON(&a)
	if err := utils.Verify(a, utils.AutoCodeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if a.AutoCreateApiToSql {
		apiList := [6]model.SysApi{
			{
				Path:        "/" + a.Abbreviation + "/" + "create" + a.StructName,
				Description: "新增" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "POST",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName,
				Description: "删除" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "DELETE",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName + "ByIds",
				Description: "批量删除" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "DELETE",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "update" + a.StructName,
				Description: "更新" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "PUT",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "find" + a.StructName,
				Description: "根据ID获取" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "GET",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "get" + a.StructName + "List",
				Description: "获取" + a.Description + "列表",
				ApiGroup:    a.Abbreviation,
				Method:      "GET",
			},
		}
		for _, v := range apiList {
			if err := service.AutoCreateApi(v); err != nil {
				global.GVA_LOG.Error("自动化创建失败!请自行清空垃圾数据!", zap.Any("err", err))
				c.Writer.Header().Add("success", "false")
				c.Writer.Header().Add("msg", url.QueryEscape("自动化创建失败!请自行清空垃圾数据!"))
				return
			}
		}
	}
	err := service.CreateTemp(a)
	if err != nil {
		c.Writer.Header().Add("success", "false")
		c.Writer.Header().Add("msg", url.QueryEscape(err.Error()))
		_ = os.Remove("./ginvueadmin.zip")
	} else {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Writer.Header().Add("success", "true")
		c.File("./ginvueadmin.zip")
		_ = os.Remove("./ginvueadmin.zip")
	}
}

// @Tags AutoCode
// @Summary 获取当前数据库所有表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getTables [get]
func GetTables(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	err, tables := service.GetTables(dbName)
	if err != nil {
		global.GVA_LOG.Error("查询table失败!", zap.Any("err", err))
		response.FailWithMessage("查询table失败", c)
	} else {
		response.OkWithDetailed(gin.H{"tables": tables}, "获取成功", c)
	}
}

// @Tags AutoCode
// @Summary 获取当前所有数据库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getDatabase [get]
func GetDB(c *gin.Context) {
	if err, dbs := service.GetDB(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"dbs": dbs}, "获取成功", c)
	}
}

// @Tags AutoCode
// @Summary 获取当前表所有字段
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /autoCode/getColumn [get]
func GetColumn(c *gin.Context) {
	dbName := c.DefaultQuery("dbName", global.GVA_CONFIG.Mysql.Dbname)
	tableName := c.Query("tableName")
	if err, columns := service.GetColumn(tableName, dbName); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"columns": columns}, "获取成功", c)
	}
}
