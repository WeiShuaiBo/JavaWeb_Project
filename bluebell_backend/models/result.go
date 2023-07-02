package models

type Result struct {
	id uint64 `gorm:"id"`
	*Project
	*ProjectDetail
}
