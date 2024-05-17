package models

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/pranay999000/feeds/configs"
)

type Like struct {
	gorm.Model
	UserId		string		`json:"user_id"`
	FeedId		int64		`json:"feed_id"`
	Feed		Feed		`json:"feed"`
}

var readLikeDB *gorm.DB
var writeLikeDB *gorm.DB
var transactionDB *gorm.DB

func init() {
	configs.WriteConnect()
	configs.ReadConnect()
	configs.TransactionConnect()

	writeLikeDB = configs.GetWriteDB()
	readLikeDB = configs.GetReadDB()
	transactionDB = configs.GetTransactionDB()

	writeLikeDB.AutoMigrate(&Like{})
	readLikeDB.AutoMigrate(&Like{})
}

func (like *Like) CreateLike() error {

	channelLike := make(chan Like, 1)
	channelFeed := make(chan Feed, 1)

	go checkLike(like.UserId, like.FeedId, channelLike)
	go CheckFeed(like.FeedId, channelFeed)

	channelCheckLike := <- channelLike
	channelCheckFeed := <- channelFeed

	if reflect.ValueOf(channelCheckFeed).IsZero() {
		return fmt.Errorf("no feed found")
	}

	transactionDB.Transaction(func(tx *gorm.DB) error {
		if !reflect.ValueOf(channelCheckLike).IsZero() {
			if err := tx.Delete(&Like{}, channelCheckLike.ID).Error; err != nil {
				return err
			}

			if err := tx.Model(&Feed{}).Where("ID=?", like.FeedId).Update("like_count", channelCheckFeed.LikeCount - 1).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Create(like).Error; err != nil {
				return err
			}

			if err := tx.Model(&Feed{}).Where("ID=?", like.FeedId).Update("like_count", channelCheckFeed.LikeCount + 1).Error; err != nil {
				return err
			}
		}

		return nil
	})
	
	return nil
}

func GetLikeByFeed(feedId int64) []Like {
	var likes []Like
	readLikeDB.Where("feed_id=?", feedId).Preload("Feed").Find(&likes)
	return likes
}

func checkLike(user_id string, feed_id int64, channel chan Like) {
	var like Like
	readLikeDB.Where("user_id=? AND feed_id=?", user_id, feed_id).First(&like)

	channel <- like
}