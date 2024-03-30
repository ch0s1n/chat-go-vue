package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"testing"
)

func TestAddUser(t *testing.T) {
	user := User{

		Account:  "asdfjkh",
		Password: "asdf",
		Nickname: "",
	}
	InitDB()
	err, a := SearchUserList(&user)
	if err == nil {
		println("111klsajfd")
		println(len(a))
		println(a[0].Password)
	}

}

func TestInitDB(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	mydb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Print(err)
	}
	//var user []*User
	//mydb.Raw(`select * from user`).Scan(&user)
	//for _, item := range user {
	//	println(item)
	//}
	mydb.Exec(`UPDATE user_status set status = 1 where user_id = 31`)
}
