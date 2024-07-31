package model

import (
	"go-practice/go_zero/shop/internal/types"
	"gorm.io/gorm"
	"reflect"
)

// 其他Model的基类
type BaseModel struct {
	Id        uint            `gorm:"primary_key"`
	CreatedAt *types.JSONTime `json:",omitempty"`
	UpdatedAt *types.JSONTime `json:",omitempty"`
	PerPage   int64           `gorm:"-" json:",omitempty"`
	Page      int64           `gorm:"-" json:",omitempty"`
	OrderBy   string          `gorm:"-" json:",omitempty"`
	Selects   []string        `gorm:"-" json:",omitempty"`
	Omits     []string        `gorm:"-" json:",omitempty"`
	Withs     []string        `gorm:"-" json:",omitempty"`
	DB        *gorm.DB        `gorm:"-" json:"-"`
	//CreatedId        uint            `gorm:"-"`
	//UpdatedId        uint            `gorm:"-"`
	ModelDescription string       `gorm:"-" orm:"-" json:"-"`
	ModelObject      interface{}  `gorm:"-" orm:"-" json:"-"`
	ModelObjects     reflect.Type `gorm:"-" orm:"-" json:"-"`
	StartPage        int64        `gorm:"-" json:"startPage,omitempty"` //导出开始页
	EndPage          int64        `gorm:"-" json:"endPage,omitempty"`   //导出结束页
	QueryAll         bool         `gorm:"-" json:"queryAll,omitempty"`  //查询所有数据
}
