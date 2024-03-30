package model

import "log"

type User struct {
	Id       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}
type UserStatus struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	//0代表离线 1代表上线
	Status int `json:"status"`
}

func AddUser(user *User) error {
	begin := MyDb.Begin()
	err1 := begin.Table("user").Create(&user).Error
	if err1 != nil {
		log.Print(err1)
		begin.Rollback()
		return err1
	}
	status := UserStatus{
		UserId: user.Id,
		Status: 0,
	}
	err2 := begin.Table("user_status").Create(&status).Error
	if err2 != nil {
		log.Println(err2)
		begin.Rollback()
		return err2
	}
	begin.Commit()
	return nil
}

func CheckUniqueAccount(account string) (bool, error) {
	var result []*User
	err := MyDb.Raw(`select * from user where account = ?`, account).Scan(&result).Error
	return len(result) == 0, err
}

func CheckUserPassword(user *User) ([]*User, error) {
	var result []*User
	err := MyDb.Raw(`select * from user where account = ? and password = ?`, user.Account, user.Password).Scan(&result).Error
	return result, err
}

func SearchUserList(user *User) (error, []*User) {
	var result []*User
	var err error
	if user.Nickname != "" {
		err = MyDb.Raw(`select id,account,nickname from user where nickname = ?`, user.Nickname).Scan(&result).Error
	} else {
		err = MyDb.Raw(`select id,account,nickname from user where account = ?`, user.Account).Scan(&result).Error

	}
	return err, result
}

func ChangeStatus(userStatus *UserStatus) error {
	err := MyDb.Exec(`update user_status set status = ? where user_id = ?`, userStatus.Status, userStatus.UserId).Error
	return err
}
