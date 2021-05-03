package domain

type Business struct {
	BusinessID   int    `gorm:"column:id"`
	BusinessName string `gorm:"column:name"`
}
