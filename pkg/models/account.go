package models

import (
	"gorm.io/datatypes"
)

type AccountsEntry struct {
	AccountId string            `gorm:"primaryKey;column:accountId;type:varchar(255);not null" json:"accountId"`
	Timestamp int64             `gorm:"not null" json:"timestamp"`
	Data      datatypes.JSONMap `gorm:"type:json;not null" json:"data"`
}

func (AccountsEntry) TableName() string {
	return "accountsEntry"
}
