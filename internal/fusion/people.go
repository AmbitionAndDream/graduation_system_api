package fusion

import (
	"graduation_system_api/internal/errors"
	// "graduation_system_api/internal/global"
	"graduation_system_api/internal/db"
	"log"
)

// type User struct {
// 	Id int
// 	Phone string
// 	Name string
// 	Pwd string
// 	Is_admin int
// 	Role_type int
// 	Avatar string
// }

func peopleDel(phone string) (error) {
	conn := db.GetDb()
	defer conn.Close()
	dbInstance := conn.Table("user").Where("phone = ?", phone).Delete(nil)
	log.Println(dbInstance, "db")
	if err := dbInstance.Error; err != nil {
		return errors.New(errors.ServerError, "")
	}else if dbInstance.RowsAffected == 0 {
		return errors.New(errors.ServerError, "查无此人")
	}else {
		return nil
	}
}