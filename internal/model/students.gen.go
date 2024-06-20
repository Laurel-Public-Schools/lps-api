// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameStudent = "students"

// Student mapped from table <students>
type Student struct {
	StudentEmail string `gorm:"column:studentEmail;not null" json:"studentEmail"`
	StudentName  string `gorm:"column:studentName;not null" json:"studentName"`
	ClassroomID  string `gorm:"column:classroomId;not null" json:"classroomId"`
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Status       string `gorm:"column:status;not null;default:default" json:"status"`
}

// TableName Student's table name
func (*Student) TableName() string {
	return TableNameStudent
}