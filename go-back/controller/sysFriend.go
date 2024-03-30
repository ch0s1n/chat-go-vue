package controller

import (
	"github.com/gin-gonic/gin"
	"go-back/model"
	"go-back/utils"
	"log"
	"net/http"
	"strconv"
)

func DoAddFriend(context *gin.Context) {
	var friend model.Friend
	_ = context.ShouldBind(&friend)
	friends, err := model.GetFriendListByUserId(friend.UserId)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, item := range friends {
		if item.FriendId == friend.FriendId {
			context.JSON(http.StatusBadRequest, utils.HandleResponse("已是好友,请勿重复添加", nil))
			return
		}

	}
	err = model.AddFriend(&friend)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.HandleResponse("添加好友失败,请重试", nil))
		log.Println(err)
		return
	}
	context.JSON(http.StatusOK, utils.HandleResponse("添加成功", nil))
}

func GetFriendList(context *gin.Context) {
	var m map[string]string
	_ = context.BindJSON(&m)
	userid, _ := strconv.Atoi(m["userid"])
	friends, err := model.GetFriendListByUserId(userid)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.HandleResponse("获取好友列表失败,请重试", nil))
		log.Println(err)
		return
	}
	context.JSON(http.StatusOK, utils.HandleResponse("获取好友列表成功", friends))
}

func SearchFriend(context *gin.Context) {
	var user *model.User
	_ = context.ShouldBind(&user)
	err, users := model.SearchUserList(user)
	if err == nil {
		if len(users) == 0 {
			context.JSON(http.StatusAccepted, utils.HandleResponse("未找到好友", nil))
			return
		}
		context.JSON(http.StatusOK, utils.HandleResponse("搜索好友为", users))
		return
	}
	log.Print(err)
	context.JSON(http.StatusBadRequest, utils.HandleResponse("搜索错误", nil))

}
