package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julianespinel/mila-api/bvc"
)

func initializeBVC(db *gorm.DB) bvc.Domain {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	client := bvc.InitClient(httpClient)
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
	gocron.Every(1).Day().At("23:00").Do(bvcApi.UpdateDailyStocks, time.Now())

	// Start all the pending jobs
	<-gocron.Start()
}
