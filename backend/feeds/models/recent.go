package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pranay999000/feeds/configs"
)

type Recent struct {
	gorm.Model
	FeedId		int64	`json:"feed_id"`
	Feed		Feed	`json:"feed"`
}

var readRecentDB *gorm.DB
var writeRecentDB *gorm.DB

func init() {
	configs.WriteConnect()
	configs.ReadConnect()

	writeRecentDB = configs.GetWriteDB()
	readRecentDB = configs.GetReadDB()

	writeRecentDB.AutoMigrate(&Recent{})
	readRecentDB.AutoMigrate(&Recent{})
}

func CreateRecent(feed_id int64) {
	var count int64
	readRecentDB.Model(&Recent{}).Count(&count)

	if count >= 8 {
		var recent Recent
		readRecentDB.Order("created_at").Limit(1).First(&recent)
		writeRecentDB.Where("ID=?", recent.ID).Delete(&Recent{})
	}
	newRecent := &Recent{FeedId: feed_id}
	writeRecentDB.NewRecord(newRecent)
	writeRecentDB.Create(newRecent)
}

func GetRecent() []Recent {
	var recents []Recent
	readRecentDB.Preload("Feed").Find(&recents)
	return recents
}

