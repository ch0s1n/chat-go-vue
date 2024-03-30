package router

import (
	"github.com/gin-gonic/gin"
	"go-back/controller"
	"go-back/utils"
	"log"
	"net/http"
)

func Start() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})
	//登陆注册
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/change-status", controller.ChangeUserStatus)
	//好友
	r.POST("/add-friend", utils.JWTMiddleware(), controller.DoAddFriend)
	r.POST("/get-friends", utils.JWTMiddleware(), controller.GetFriendList)
	r.POST("/search-friend", utils.JWTMiddleware(), controller.SearchFriend)
	err := r.Run()
	if err != nil {
		log.Println("======================>>启动失败")
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
