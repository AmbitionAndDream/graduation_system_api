package domain

type RequestDemand struct {
	DemandName           string              `json:"name" binding:"required"`
	DemandLink           string              `json:"link" binding:"required"`
	DemandPriorityStatus int                 `json:"priority_status" binding:"required"`
	DemandNote           string              `json:"note" binding:"required"`
	BusinessId           int                 `json:"business_id" binding:"required"`
	PeoplePhone          string              `json:"user_id" binding:"required"` //需求负责人(创建)
	DemandNodeInfo       []RequestDemandInfo `json:"info" binding:"required"`
}

type RequestDemandInfo struct {
	ItemId      int    `json:"item_id" binding:"required"`
	ItemType    int    `json:"item_type" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Avatar      string `json:"avatar" binding:"required"`
	PeoplePhone string `json:"info_user_id" binding:"required"` //需求处理人（解决）
	PeopleName  string `json:"user_name" binding:"required"`
	Time        int64  `json:"time" binding:"required"`
	Status      int64  `json:"status" binding:"required"`
}

type RequestDemandItem struct {
	ItemId          int    `json:"item_id" binding:"required"`
	DemandId        int    `json:"demand_id" binding:"required"`
	NodePeoplePhone string `json:"user_id" binding:"required"`
}

type RequestDemandSetTime struct {
	ItemId          int    `json:"item_id" binding:"required"`
	DemandId        int    `json:"demand_id" binding:"required"`
	NodePeoplePhone string `json:"user_id" binding:"required"`
	Time            int64  `json:"time" binding:"required"`
}

type RequestDeleteDemandId struct {
	DemandId []int `json:"demand_id" binding:"required"`
}

type ResponseDemandList struct {
	Total      int64            `json:"total"`
	Limit      int              `json:"limit"`
	Offset     int              `json:"offset"`
	DemandList []ResponseDemand `json:"list"`
}

type ResponseDemand struct {
	DemandId             int                      `json:"demand_id"`
	DemandName           string                   `json:"name"`
	DemandLink           string                   `json:"link"`
	DemandPriorityStatus int                      `json:"priority_status"`
	DemandNote           string                   `json:"note"`
	BusinessId           int                      `json:"business_id"`
	PeoplePhone          string                   `json:"user_id"`
	DemandNodeInfo       []ResponseDemandNodeInfo `json:"info"`
}

type ResponseDemandNodeInfo struct {
	ItemId      int    `json:"item_id"`
	ItemType    int    `json:"item_type"`
	Title       string `json:"title"`
	Avatar      string `json:"avatar"`
	PeoplePhone string `json:"info_user_id" `
	PeopleName  string `json:"user_name"`
	Time        int64  `json:"time"`
	Status      int64  `json:"status"`
}

type ResponseDemandPoolList struct {
	ReviewPool      *ResponseReviewPool      `json:"review_pool"`
	DevelopmentPoll *ResponseDevelopmentPoll `json:"development_poll"`
	TestPoll        *ResponseTestPoll        `json:"test_poll"`
	AcceptancePoll  *ResponseAcceptancePoll  `json:"acceptance_poll"`
	CompletePoll    *ResponseCompletePoll    `json:"complete_poll"`
}

type ResponseReviewPool struct {
	ResponseDemandList
}
type ResponseDevelopmentPoll struct {
	ResponseDemandList
}
type ResponseTestPoll struct {
	ResponseDemandList
}
type ResponseAcceptancePoll struct {
	ResponseDemandList
}
type ResponseCompletePoll struct {
	ResponseDemandList
}
