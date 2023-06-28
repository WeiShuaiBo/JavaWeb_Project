package models

import "time"

type Community struct {
	CommunityID   uint64 `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
}

type CommunityDetail struct {
	CommunityID   uint64    `json:"community_id" gorm:"column:community_id"`
	CommunityName string    `json:"community_name" gorm:"column:community_name"`
	Introduction  string    `json:"introduction,omitempty" gorm:"column:introduction"`
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
}

func (c Community) TableName() string {
	return "community"
}
func (c CommunityDetail) TableName() string {
	return "community"
}
