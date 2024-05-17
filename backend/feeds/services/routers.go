package services

import "github.com/gin-gonic/gin"

func FeedRouters(r *gin.Engine) {
	r.GET("/feeds/:user_id", GetFeeds())
	r.POST("/create", CreateFeed())
	r.PUT("/like", LikeFeed())
	r.GET("/recents", GetRecents())
	r.PUT("/update/view", UpdateView())
	r.GET("/popular", GetPopular())
}