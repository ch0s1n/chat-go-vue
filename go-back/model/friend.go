package model

type Friend struct {
	Id             int    `json:"id"`
	UserId         int    `json:"user_id"`
	FriendId       int    `json:"friend_id"`
	FriendAccount  string `json:"friend_account"`
	FriendNickname string `json:"friend_nickname"`
}

func AddFriend(f *Friend) error {
	return MyDb.Table("friend").Create(&f).Error
}

func GetFriendListByUserId(userId int) ([]*Friend, error) {
	var friends []*Friend
	err := MyDb.Raw(`select * from friend where user_id = ? order by id`, userId).Scan(&friends).Error
	return friends, err
}
