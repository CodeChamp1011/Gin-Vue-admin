package model

import (
	"gin-vue-admin/global"
	"gorm.io/gorm"
	"time"
)


type GVA_Workflow interface {
	CreateWorkflow(WorkflowNode,GVA_Workflow)
	GetBusinessType()string
}

type WorkflowFunc struct {
	BusinessType string
}

func(w *WorkflowFunc)CreateWorkflow(node WorkflowNode,businessModel GVA_Workflow){

}

func(w *WorkflowFunc)GetBusinessType()(businessType string){
	return w.BusinessType
}

//定义clazz常量

const (
	USER_TASK     string = "userTask"
	SCRIPT_TASK   string = "scriptTask"
	RECEIVE_TASK  string = "receiveTask"
	MAIL_TASK     string = "mailTask"
	TIMER_START   string = "timerStart"
	MESSAGE_START string = "messageStart"
	GATEWAY       string = "gateway"
	FLOW          string = "flow"
	START         string = "start"
	END           string = "end"
	PROCESS       string = "process"
)

type  WorkflowMove struct {
	global.GVA_MODEL
	WorkflowProcessID string `json:"workflowProcessID" gorm:"comment:工作流模板ID"`
	BusinessType string `json:"businessType" gorm:"comment:业务标记"`
	BusinessID uint `json:"businessID" gorm:"comment:业务ID"`
	WorkflowMoveInfo []WorkflowMoveInfo `json:"workflowMoveInfo" gorm:"comment:当前流转详情"`
	Promoter uint `json:"Promoter" gorm:"comment:当前流转发起人"`
}


type WorkflowMoveInfo struct {
	global.GVA_MODEL
	WorkflowMoveID uint `json:"workflowMoveID" gorm:"comment:关联WorkflowMove"`
	WorkflowNodeID string `json:"workflowNodeID" gorm:"comment:流程节点ID"`
	Params   string `json:"params" gorm:"comment:流转参数"`
	Description string `json:"description" gorm:"comment:流转说明"`
	Status   int `json:"status" gorm:"comment:流转状态（来自字典 0为进行中）"`
}

type WorkflowProcess struct {
	ID          string `json:"id" form:"id" gorm:"comment:流程标识;primaryKey;unique;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Name        string         `json:"name" gorm:"comment:流程名称"`
	Category    string         `json:"category" gorm:"comment:分类"`
	Clazz       string         `json:"clazz" gorm:"comment:类型"`
	Label       string         `json:"label" gorm:"comment:流程标题"`
	HideIcon    bool           `json:"hideIcon" gorm:"comment:是否隐藏图标"`
	Description string         `json:"description" gorm:"comment:详细介绍"`
	View        string         `json:"view" gorm:"comment:前端视图文件"`
	Nodes       []WorkflowNode `json:"nodes"` // 流程节点数据
	Edges       []WorkflowEdge `json:"edges"` // 流程链接数据
}

type WorkflowNode struct {
	ID                string `json:"id" form:"id" gorm:"comment:节点id;primaryKey;unique;not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
	WorkflowProcessID string         `json:"-" gorm:"comment:流程标识"`
	Clazz             string         `json:"clazz" gorm:"comment:节点类型"`
	Label             string         `json:"label" gorm:"comment:节点名称"`
	Type              string         `json:"type" gorm:"comment:图标类型"`
	Shape             string         `json:"shape" gorm:"comment:形状"`
	Description       string         `json:"description" gorm:"comment:详细介绍"`
	View              string         `json:"view" gorm:"comment:前端视图文件"`
	X                 float64        `json:"y" gorm:"comment:x位置"`
	Y                 float64        `json:"x" gorm:"comment:y位置"`
	WaitState         string         `json:"waitState" gorm:"comment:等待属性"`
	StateValue        string         `json:"stateValue" gorm:"comment:等待值"`
	To                string         `json:"to" gorm:"comment:收件人"`
	Subject           string         `json:"subject" gorm:"comment:标题"`
	Content           string         `json:"content" gorm:"comment:内容"`
	Cycle             string         `json:"cycle" gorm:"comment:循环时间"`
	Duration          string         `json:"duration" gorm:"comment:持续时间"`
	HideIcon          bool           `json:"hideIcon" gorm:"comment:是否隐藏图标"`
	DueDate           time.Time      `json:"dueDate" gorm:"comment:到期时间"`
	AssignType        string         `json:"assignType" gorm:"comment:审批类型"`
	AssignValue       string         `json:"assignValue" gorm:"comment:审批类型值"`
	Success           bool           `json:"success" gorm:"comment:是否成功"`
}

type WorkflowEdge struct {
	ID                  string `json:"id" form:"id" gorm:"comment:唯一标识;primaryKey;unique;not null"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt     `json:"-" gorm:"index"`
	WorkflowProcessID   string             `json:"-" gorm:"comment:流程标识"`
	Clazz               string             `json:"clazz" gorm:"comment:类型（线）"`
	Source              string             `json:"source" gorm:"comment:起点节点"`
	Target              string             `json:"target" gorm:"comment:目标节点"`
	SourceAnchor        int                `json:"sourceAnchor" gorm:"comment:起点"`
	TargetAnchor        int                `json:"targetAnchor" gorm:"comment:目标点"`
	Description         string             `json:"description" gorm:"comment:详细介绍"`
	Shape               string             `json:"shape" gorm:"comment:形状"`
	StartPoint          WorkflowStartPoint `json:"startPoint"` // 起点信息
	EndPoint            WorkflowEndPoint   `json:"endPoint"`   // 终点信息
	Label               string             `json:"label" gorm:"comment:标题"`
	HideIcon            bool               `json:"hideIcon" gorm:"comment:隐藏图标"`
	ConditionExpression string             `json:"conditionExpression" gorm:"comment:条件标识"`
	Seq                 string             `json:"seq" gorm:"comment:序号"`
	Reverse             bool               `json:"reverse" gorm:"comment:是否反向"`
}

type WorkflowStartPoint struct {
	WorkflowEdgeID string
	WorkflowPoint
}

type WorkflowEndPoint struct {
	WorkflowEdgeID string
	WorkflowPoint
}

//
type WorkflowPoint struct {
	global.GVA_MODEL
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Index int     `json:"index"`
}
