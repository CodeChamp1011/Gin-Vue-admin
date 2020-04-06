// 自动生成模板{{.StructName}}
package {{.PackageName}}

import (
	"gin-vue-admin/global"
    "github.com/jinzhu/gorm"
    "github.com/pkg/errors"
)

type {{.StructName}} struct {
      gorm.Model {{range .Fields}}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}"{{if .ColumnName}} gorm:"column:{{.ColumnName}}"{{end}}`{{ end }}
}

// 创建{{.StructName}}
func ({{.Abbreviation}} *{{.StructName}})Create{{.StructName}}()(err error){
        err = global.GVA_DB.Create({{.Abbreviation}}).Error
        return err
}

// 删除{{.StructName}}
func ({{.Abbreviation}} *{{.StructName}})Delete{{.StructName}}()(err error){
        err = global.GVA_DB.Delete({{.Abbreviation}}).Error
        return err
}

// 更新{{.StructName}}
func ({{.Abbreviation}} *{{.StructName}})Update{{.StructName}}()(err error, re{{.Abbreviation}} {{.StructName}}){
        err = global.GVA_DB.Save({{.Abbreviation}}).Error
        return err, *{{.Abbreviation}}
}

// 根据ID查看单条{{.StructName}}
func ({{.Abbreviation}} *{{.StructName}})FindById()(err error,re{{.Abbreviation}} {{.StructName}}){
    err = global.GVA_DB.Where("id = ?",{{.Abbreviation}}.ID).First(&re{{.Abbreviation}}).Error
    return err,re{{.Abbreviation}}
}

// 分页获取{{.StructName}}
func ({{.Abbreviation}} *{{.StructName}})GetInfoList(info PageInfo)(err error, list interface{}, total int){
    	limit := info.PageSize
        offset := info.PageSize * (info.Page - 1)
        db:=global.GVA_DB
    	if err != nil {
    		return
    	} else {
    		var re{{.StructName}}List []{{.StructName}}
    		err = db.Limit(limit).Offset(offset).Find(&re{{.StructName}}List).Error
    		return err, re{{.StructName}}List, total
    	}
}