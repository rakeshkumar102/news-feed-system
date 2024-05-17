package services

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/feeds/configs"
	"github.com/pranay999000/feeds/models"
	"github.com/pranay999000/feeds/utils"
)

func GetFeeds() gin.HandlerFunc {
	return func(c *gin.Context) {
		limitStr := c.DefaultQuery("limit", "12")
		pageStr := c.DefaultQuery("page", "1")
		user_id := c.Param("user_id")

		followUrl, err := configs.EnvMap("follows")

		utils.FailOnError(err, "Unable to find follow url")

		req, err := http.NewRequest("GET", followUrl + "follow/" + user_id + "/following", nil)
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		utils.FailOnError(err, "Failed to create followers request")

		client := &http.Client{}

		res, err := client.Do(req)

		utils.FailOnError(err, "Failed to make followers request")

		body, err := io.ReadAll(res.Body)

		utils.FailOnError(err, "Failed to read respose")

		var user models.User

		
		if err := json.Unmarshal(body, &user); err != nil {
			utils.FailOnError(err, "Failed to unmarshal users data")
		}
		
		var user_ids []string

		for _, u := range user.Users {
			user_ids = append(user_ids, u.ID)
		}

		limit, _ := strconv.Atoi(limitStr)
		page, _ := strconv.Atoi(pageStr)


		feeds := models.GetFeeds(int64(limit), int64(page), user_ids)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"feeds": feeds,
		})
	}
}

func CreateFeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newFeed models.Feed
		c.Bind(&newFeed)

		if newFeed.Title != "" && len(newFeed.Body) > 250 && len(newFeed.Title) < 30 && newFeed.UserId != "" {
			feed := newFeed.CreateFeed()
			models.CreateRecent(int64(feed.ID))

			c.JSON(http.StatusCreated, gin.H{
				"success": true,
				"data": feed,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Either title is null or length of title is long ot length of body is short",
			})
		}
	}
}

func LikeFeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var like models.Like
		c.Bind(&like)

		if like.FeedId != 0 && like.UserId != "" {
			err := like.CreateLike()

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
			})
		}
	}
}

func GetRecents() gin.HandlerFunc {
	return func(c *gin.Context) {
		recents := models.GetRecent()

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": recents,
		})
	}
}

func UpdateView() gin.HandlerFunc {
	return func(c *gin.Context) {
		feedId := c.Query("feedId")
		f_id, _ := strconv.Atoi(feedId)
		channel := make(chan models.Feed, 1)
		go models.CheckFeed(int64(f_id), channel)

		feed := <-channel

		if reflect.ValueOf(feed).IsZero() {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "feed not found",
			})
			return
		} else {
			models.CreateView(int64(f_id), feed.ViewCount)
			c.JSON(http.StatusOK, gin.H{
				"success": true,
			})
		}
	}
}

func GetPopular() gin.HandlerFunc {
	return func (c *gin.Context) {
		feeds := models.GetPopular()

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": feeds,
		})
	}
}