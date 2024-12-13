package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
    ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
    StudentPhone    string         `gorm:"unique;type:varchar(64);not null" json:"student_phone"`
    StudentEmail    string         `gorm:"unique;type:varchar(255);not null" json:"student_email"`
    StudentID       string         `gorm:"unique;type:varchar(30);not null" json:"student_id"`
    Password        string         `gorm:"type:varchar(255);not null" json:"-"`  // 密码不应该在JSON中返回
    StudentName     string         `gorm:"type:varchar(255);not null" json:"student_name"`
    Gender          int            `gorm:"type:tinyint" json:"gender"`
    Nickname        string         `gorm:"type:varchar(20)" json:"nickname"`
    StudentAvatarID uint           `json:"student_avatar_id"`
    CollegeID       uint           `json:"college_id"`
    CreatedAt       time.Time      `json:"created_at"`
    UpdatedAt       time.Time      `json:"updated_at"`
    DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
