package models

import (
	"fmt"
	"math"

	"github.com/jinzhu/gorm"
	"github.com/pranay999000/feeds/configs"
)

type Popular struct {
	gorm.Model
	FeedId		int64	`json:"feed_id"`
	Feed		Feed	`json:"feed"`
}

var readPopularDB *gorm.DB
var writePopularDB *gorm.DB
var transactionPopularDB *gorm.DB

func init() {
	configs.WriteConnect()
	configs.ReadConnect()
	configs.TransactionConnect()

	writePopularDB = configs.GetWriteDB()
	readPopularDB = configs.GetReadDB()
	transactionPopularDB = configs.GetTransactionDB()

	writePopularDB.AutoMigrate(&Popular{})
	readPopularDB.AutoMigrate(&Popular{})
}

func CreateView(feed_id int64, view_count int64) {

	transactionPopularDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Feed{}).Where("ID=?", feed_id).Update("view_count", view_count + 1).Error; err != nil {
			return err
		}

		var popular []Popular
		readPopularDB.Preload("Feed").Find(&popular)

		needUpdate := true

		if len(popular) < 3 {
			for _, v := range popular {
				if v.Feed.ID == uint(feed_id) {
					needUpdate = false
					break
				}
			}

			if needUpdate {
				newPopular := &Popular{FeedId: int64(feed_id)}
				fmt.Println("1 error")
				if err := tx.Create(newPopular).Error; err != nil {
					return err
				}
			}
		} else {
			var min int64 = math.MaxInt64
			minFeed := Feed{}
			popularId := 0

			for _, v := range popular {
				if v.Feed.ViewCount < min {
					min = v.Feed.ViewCount
					minFeed = v.Feed
					popularId = int(v.ID)
				}
			}

			fmt.Println("2 error")
			if minFeed.ViewCount <= view_count + 1 {
				if err := tx.Model(&Popular{}).Where("ID=?", popularId).Update("feed_id", feed_id).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func UpdatePopular(feed Feed) {
	var popular []Popular
	readPopularDB.Preload("Feed").Find(&popular)

	if len(popular) < 3 {
		for _, v := range popular {
			if v.Feed.ID == feed.ID {
				return
			}
		}
		newPopular := &Popular{FeedId: int64(feed.ID)}
		writePopularDB.NewRecord(newPopular)
		writePopularDB.Create(newPopular)
	} else {
		var min int64 = math.MaxInt64
		minFeed := Feed{}
		popularId := 0
		for _, v := range popular {
			if v.Feed.ViewCount < min {
				min = v.Feed.ViewCount
				minFeed = v.Feed
				popularId = int(v.ID)
			}
		}

		if minFeed.ViewCount <= feed.ViewCount {
			writePopularDB.Model(&Popular{}).Where("ID=?", popularId).Update("feed_id", feed.ID)
		}
	}
}

func GetPopular() []Feed {
	var feeds []Feed
	readPopularDB.Preload("Feed").Find(&feeds)
	return feeds
}