package domain

type Bug struct {
	SystemId       int    `gorm:"column:system_id"`
	BugId          int    `gorm:"column:id"`
	DemandId       int    `gorm:"column:demand_id"`
	BugName        string `gorm:"column:bug_name"`
	PriorityStatus int    `gorm:"column:priority_status" `
	ReporterId     string `gorm:"column:reporter_id" `
	HandlerId      string `gorm:"column:handler_id" `
	Type           int    `gorm:"column:type"`
	Opportunity    int    `gorm:"column:opportunity"`
	BeginTime      int64  `gorm:"column:begin_time"`
	SolveType      int    `gorm:"column:solve_type"`
	Desc           string `gorm:"column:bug_desc"`
	Status		   int    `gorm:"column:status"`
}
