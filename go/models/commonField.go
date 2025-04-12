package models

import (
	"time"

	"gorm.io/gorm"
)

// 共通フィールド
type CommonFields struct {
	InsTimestamp time.Time  `gorm:"column:ins_timestamp;autoCreateTime" json:"ins_timestamp"`
	InsUserID    int        `gorm:"column:ins_user_id" json:"ins_user_id"`
	InsAction    string     `gorm:"column:ins_action" json:"ins_action"`
	UpdTimestamp time.Time  `gorm:"column:upd_timestamp;autoUpdateTime" json:"upd_timestamp"`
	UpdUserID    int        `gorm:"column:upd_user_id" json:"upd_user_id"`
	UpdAction    string     `gorm:"column:upd_action" json:"upd_action"`
	DelTimestamp gorm.DeletedAt `gorm:"column:del_timestamp" json:"del_timestamp"`
	DelUserID    *int       `gorm:"column:del_user_id" json:"del_user_id"`
	DelAction    *string    `gorm:"column:del_action" json:"del_action"`
}