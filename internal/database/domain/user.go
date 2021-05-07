package domain

type User struct {
	Id          int
	Name        string
	PhoneNumber string `gorm:"column:phone"`
	PassWord    string `gorm:"column:pwd"`
	RoleType    int
	IsAdmin     int
	Avatar 		string
}
