package domain

type Bug struct {
	SystemId       int    `json:"column:system_id" `
	BugId          int    `json:"column:bug_id" `
	DemandId       int    `json:"column:demand_id"`
	BugName        int    `json:"column:bug_name"`
	PriorityStatus int    `json:"column:priority_status" `
	ReporterId     string `json:"column:reporter_id" `
	HandlerId      string `json:"column:handler_id" `
	Type           int    `json:"column:type"`
	Opportunity    int    `json:"column:opportunity"`
	BeginTime      int64  `json:"column:begin_time"`
	SolveType      int    `json:"column:solve_type"`
	Desc           string `json:"column:bug_desc"`
}
