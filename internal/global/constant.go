package global

//role
const (
	//管理员
	Administrator = 1
	//普通用户
	NonAdministrator = 0
)

//action
const (
	Create             = "add"
	Select             = "list"
	Delete             = "del"
	SelectDemandMyList = "my_list"
	Solve              = "solve"
	SetTime            = "set_time"
	Item               = "item"
	SelectPoolList     = "pool_list"
	Detail             = "detail"
)

//kind
const (
	PeoPle   = "people"
	Business = "business"
	Demand   = "demand"
	Bug      = "bug"
)

//item status
const (
	Finish    = 2
	NotFinish = 1
)

//item type
const (
	ReviewPoll = 1
	DevelopmentPoll = 2
	TestPoll = 3
	AcceptancePoll = 4
	//CompletePoll
)

const (
	ReporterId = 0
	HandlerId  = 1
)
