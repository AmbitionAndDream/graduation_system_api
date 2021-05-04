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
	Finish    = 1
	NotFinish = 0
)

//item type
const (
	ReviewPoll = iota
	DevelopmentPoll
	TestPoll
	AcceptancePoll
	CompletePoll
)

const (
	ReporterId = 0
	HandlerId  = 1
)
