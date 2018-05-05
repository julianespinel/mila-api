package main

import (
	"log"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julianespinel/mila-api/bvc"
)

func initializeBVC(db *gorm.DB) bvc.Domain {
	client := bvc.Client{}
	persistence := bvc.InitPersistence(db)
	domain := bvc.InitDomain(client, persistence)
	return domain
}

func main() {
	time.Sleep(15 * time.Second)
	db, err := gorm.Open("mysql", "usertest:passwordtest@tcp(db:3306)/miladb?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	bvcApi := initializeBVC(db)
	gocron.Every(1).Day().At("23:59").Do(bvcApi.UpdateDailyStocks)

	// Start all the pending jobs
	<-gocron.Start()
}
