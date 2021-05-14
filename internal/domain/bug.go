package domain

type RequestBugSolve struct {
	BugId     int `json:"bug_id" binding:"required"`
	SolveType int `json:"solve_type" binding:"required"`
	Status    int `json:"status" binding:"required"`
}

type RequestBug struct {
	SystemId       int    `json:"system_id" binding:"required"`
	DemandId       int    `json:"demand_id" binding:"required"`
	PriorityStatus int    `json:"priority_status" binding:"required"`
	BugName        string `json:"title" binding:"required"`
	ReporterId     string `json:"reporter_id" binding:"required"`
	HandlerId      string `json:"handler_id" binding:"required"`
	Type           int    `json:"type" binding:"required"`
	Opportunity    int    `json:"opportunity" binding:"required"`
	SolveType      int    `json:"solve_type"`
	Desc           string `json:"desc"`
}

type BugList struct {
	Limit          int    `json:"limit" binding:"required"`
	Offset         int	  `json:"offset" binding:"required"`
	Status         int	  `json:"status"`
	BeginTime      int64  `json:"begin_time"`
	EndTime        int64  `json:"end_time"`
	SolveType      int	  `json:"solve_type"`
	Type           int    `json:"type"`
	Opportunity    int	  `json:"opportunity"`
	IsAssign       int	  `json:"is_assign"`
	PriorityStatus int    `json:"priority_status"`
	SystemId       int    `json:"system_id"`
	PeoplePhone    string `json:"phone"`
}

type ResponseBugList struct {
	Total   int64         `json:"total"`
	Limit   int           `json:"limit"`
	Offset  int           `json:"offset"`
	BugList []ResponseBug `json:"bug_list"`
}

type ResponseBug struct {
	BugId          int    `json:"bug_id"`
	SystemId       int    `json:"system_id" `
	DemandId       int    `json:"demand_id"`
	BugName        string `json:"title"`
	PriorityStatus int    `json:"priority_status" `
	ReporterId     string `json:"reporter_id" `
	HandlerId      string `json:"handler_id" `
	Type           int    `json:"type"`
	Opportunity    int    `json:"opportunity"`
	BeginTime      int64  `json:"begin_time"`
	SolveType      int    `json:"solve_type"`
	Status		   int    `json:"status"`
	Desc 		   string `json:"desc"`
}
