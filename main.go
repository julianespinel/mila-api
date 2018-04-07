package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"
	"./bvc"
	"log"
)

func initializeBVC(db *gorm.DB) bvc.BVCDomain {
	client := bvc.BVCClient{}
	persistence := bvc.BVCPersistence{db: db}
	bvc := bvc.BVCDomain{
		client: client,
		persistence: persistence,
	}
	return bvc
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	bvcApi := initializeBVC(db)
	gocron.Every(1).Day().At("23:59").Do(bvcApi.UpdateDailyStocks)
}