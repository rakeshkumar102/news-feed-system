package services

import "github.com/gin-gonic/gin"

func FollowRouters(r *gin.Engine) {
	r.GET("/connect/orientdb", ConnectOrientDB())
	r.GET("/follow/:user_id/:direction", GetFollows())
	r.GET("/create/user/:user_id", CreateUserVertex())
	r.GET("/create/follow/:user_id/:following_user_id", CreateFollowEdge())
	r.GET("/unfollow/:user_id/:following_user_id", DeleteEdge())
}