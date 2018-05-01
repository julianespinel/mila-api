package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

func initializeBVC(db *gorm.DB) BVCDomain {
	client := BVCClient{}
	persistence := BVCPersistence{db: db}
	bvc := BVCDomain{
		client: client,
		persistence: persistence,
	}
	return bvc
}

func main() {
	time.Sleep(15* time.Second)
	db, err := gorm.Open("mysql", "usertest:passwordtest@tcp(db:3306)/miladb?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	bvcApi := initializeBVC(db)
	gocron.Every(1).Day().At("23:59").Do(bvcApi.UpdateDailyStocks)
}