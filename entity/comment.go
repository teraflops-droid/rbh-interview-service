package entity

import "time"

type Comment struct {
	Id          uint      `gorm:"primaryKey;autoIncrement;column:id"`
	Description string    `gorm:"column:description;type:text"`
	CreateBy    string    `gorm:"column:create_by;type:varchar(100)"`
	CardId      uint      `gorm:"column:card_id"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Comment) TableName() string {
	return "tb_comment"
}
