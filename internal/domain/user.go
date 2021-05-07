package domain

type ResponsePeopleList struct {
	Total  int64            `json:"total"`
	Limit  int              `json:"limit"`
	Offset int              `json:"offset"`
	User   []ResponsePeople `json:"user"`
}
type ResponsePeople struct {
	PeopleID    int    `json:"people_id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	RoleType    int    `json:"role_type"`
	IsAdmin     bool   `json:"is_admin"`
	Avatar 		string `json:"avatar"`
}

type RequestPeople struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	PassWord string `json:"pass_word" binding:"required"`
	IsAdmin  int    `json:"is_admin"`
	RoleType int    `json:"role_type" binding:"required"`
}

type RequestPeoplePhone struct {
	Phone string `json:"phone" binding:"required"`
}
