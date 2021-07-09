package model

import "time"

type BaseModel struct {
	ID uint `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	// MySQL的DATE/DATATIME类型可以对应Golang的time.Time
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// 有 DeletedAt(类型需要是 *time.Time) 即支持 gorm 软删除
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index"`
}
