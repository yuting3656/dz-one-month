package models

import (
	"time"
)

type Team struct {
	Uuid       string    `gorm:"AUTO_INCREMENT" json:"uuid"`
	Name       string    `gorm:"column:name" json:"name"`
	Memo       string    `gorm:"column:memo" json:"memo"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

func (Team) TableName() string {
	return "team"
}
