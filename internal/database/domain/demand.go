package domain

type Demand struct {
	DemandId             int    `gorm:"column:id"`
	DemandName           string `gorm:"column:name"`
	DemandLink           string `gorm:"column:link"`
	DemandNote           string `gorm:"column:note"`
	DemandPriorityStatus int    `gorm:"column:priority_status"`
	DemandInfo           string `gorm:"column:info"`
	BusinessId           int    `gorm:"column:business_id"`
	PeoplePhone          string `gorm:"column:user_id"`
	ModelId				 int	`gorm:"column:model_id"`
}
