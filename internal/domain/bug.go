package domain

type RequestBugSolve struct {
	BugId     int `json:"bug_id" binding:"required"`
	SolveType int `json:"solve_type" binding:"required"`
	Status    int `json:"status" binding:"required"`
}

type RequestBug struct {
	SystemId       int    `json:"system_id" binding:"required"`
	DemandId       int    `json:"demand_id" binding:"required"`
	BugName        int    `json:"title" binding:"required"`
	PriorityStatus int    `json:"priority_status" binding:"required"`
	ReporterId     string `json:"reporter_id" binding:"required"`
	HandlerId      string `json:"handler_id" binding:"required"`
	Type           int    `json:"type" binding:"required"`
	Opportunity    int    `json:"opportunity" binding:"required"`
	SolveType      int    `json:"solve_type" binding:"required"`
	Desc           string `json:"desc" binding:"required"`
}

type BugList struct {
	Limit          int
	Offset         int
	Status         int
	BeginTime      int64
	EndTime        int64
	SolveType      int
	Type           int
	Opportunity    int
	IsAssign       int
	PriorityStatus int
	PeoplePhone    string
}

type ResponseBug struct {
	BugId          int    `json:"bug_id"`
	SystemId       int    `json:"system_id" `
	DemandId       int    `json:"demand_id"`
	BugName        int    `json:"title"`
	PriorityStatus int    `json:"priority_status" `
	ReporterId     string `json:"reporter_id" `
	HandlerId      string `json:"handler_id" `
	Type           int    `json:"type"`
	Opportunity    int    `json:"opportunity"`
	BeginTime      int64  `json:"begin_time"`
	SolveType      int    `json:"solve_type"`
}
