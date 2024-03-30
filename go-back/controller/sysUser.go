package controller

import (
	"github.com/gin-gonic/gin"
	"go-back/model"
	"go-back/utils"
	"log"
	"net/http"
	"strconv"
)

func Register(context *gin.Context) {
	var user model.User
	_ = context.ShouldBind(&user)

	if user.Password == "" || user.Account == "" {
		return
	}
	isUnique, err := model.CheckUniqueAccount(user.Account)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.HandleResponse("注册用户失败,请重试", nil))
		log.Println(err)
		return
	}
	if !isUnique {
		context.JSON(http.StatusBadRequest, utils.HandleResponse("用户名重复,请重试", nil))
		return
	}
	err = model.AddUser(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.HandleResponse("注册用户失败,请重试", nil))
		log.Println(err)
		return
	}
	context.JSON(http.StatusOK, utils.HandleResponse("注册成功", nil))
}

func Login(context *gin.Context) {
	var user model.User
	_ = context.ShouldBind(&user)
	checkedUserList, err := model.CheckUserPassword(&user)
	if err != nil || len(checkedUserList) != 1 {
		log.Println(err)
		context.JSON(http.StatusBadRequest, utils.HandleResponse("用户名或密码错误", nil))
		return
	}

	token, err := utils.GenToken(checkedUserList[0].Nickname, checkedUserList[0].Account, checkedUserList[0].Id)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, utils.HandleResponse("登陆失败,请重试", nil))
		return
	}
	result := map[string]string{"token": token, "userid": strconv.Itoa(checkedUserList[0].Id), "nickname": checkedUserList[0].Nickname}
	context.JSON(http.StatusOK, utils.HandleResponse("登陆成功", result))
}

func ChangeUserStatus(context *gin.Context) {
	var userStatus model.UserStatus
	err := context.ShouldBind(&userStatus)
	if err != nil {
		log.Println(err)
		return
	}
	err = model.ChangeStatus(&userStatus)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, utils.HandleResponse("服务器错误", nil))
		return
	}
	context.JSON(http.StatusOK, utils.HandleResponse("更改状态成功", nil))
}
