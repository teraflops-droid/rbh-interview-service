package entity

import "time"

type Status string

const (
	StatusPending Status = "Pending"
	StatusTodo    Status = "Todo"
	StatusDone    Status = "Done"
)

type Card struct {
	Id          uint      `gorm:"primaryKey;autoIncrement;column:id"`
	Title       string    `gorm:"index;column:title;type:varchar(100)"`
	Description string    `gorm:"column:description;type:text"`
	Status      string    `gorm:"column:status"`
	Comments    []Comment `gorm:"ForeignKey:CardId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CreatedBy   string    `gorm:"column:created_by;type:varchar(100)"`
	UpdatedBy   string    `gorm:"column:updated_by;type:varchar(100)"`
}

func (Card) TableName() string {
	return "tb_card"
}
