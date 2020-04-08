package model

// 初始版本自动化代码工具
type AutoCodeStruct struct {
	StructName   string  `json:"structName"`
	PackageName  string  `json:"packageName"`
	Abbreviation string  `json:"abbreviation"`
	Fields       []Field `json:"fields"`
}

type Field struct {
	FieldName  string `json:"fieldName"`
	FieldType  string `json:"fieldType"`
	FieldJson  string `json:"fieldJson"`
	ColumnName string `json:"columnName"`
}