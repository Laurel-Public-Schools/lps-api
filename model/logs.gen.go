// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameLog = "logs"

// Log mapped from table <logs>
type Log struct {
	ID     string `gorm:"column:id;primaryKey" json:"id"`
	User   string `gorm:"column:user" json:"user"`
	Type   string `gorm:"column:type;not null;default:error" json:"type"`
	Action string `gorm:"column:action;not null" json:"action"`
}

// TableName Log's table name
func (*Log) TableName() string {
	return TableNameLog
}
