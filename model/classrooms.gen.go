// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameClassroom = "classrooms"

// Classroom mapped from table <classrooms>
type Classroom struct {
	ID          string `gorm:"column:id;primaryKey" json:"id"`
	RoomNumber  string `gorm:"column:roomNumber;not null" json:"roomNumber"`
	TeacherName string `gorm:"column:teacherName;not null" json:"teacherName"`
	TeacherID   string `gorm:"column:teacherId" json:"teacherId"`
	Comment     string `gorm:"column:comment" json:"comment"`
}

// TableName Classroom's table name
func (*Classroom) TableName() string {
	return TableNameClassroom
}
