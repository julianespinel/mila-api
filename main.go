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
	"github.com/kataras/iris"
)

func initializeBVC(db *gorm.DB) bvc.API {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	client := bvc.InitClient(httpClient)
	persistence := bvc.InitPersistence(db)
	domain := bvc.InitDomain(client, persistence)
	api := bvc.InitAPI(domain)
	return api
}

func logURLAndIP(ctx iris.Context) {
	ctx.Application().Logger().Infof("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())
	ctx.Next()
}

func initializeIrisApp() *iris.Application {
	app := iris.New()
	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		ctx.Next()
	})
	return app
}

func main() {
	time.Sleep(15 * time.Second)
	// TODO: get real values from a config file.
	db, err := gorm.Open("mysql", "usertest:passwordtest@tcp(db:3306)/miladb?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	bvcAPI := initializeBVC(db)
	gocron.Every(1).Day().At("23:00").Do(bvcAPI.UpdateDailyStocks, time.Now())

	// Start all the pending jobs
	<-gocron.Start()

	// Initialize Iris app
	irisApp := initializeIrisApp()
	// Define Iris routes
	milaAPIRoutes := irisApp.Party("/mila/api", logURLAndIP)
	milaAPIRoutes = bvcAPI.AddRoutes(milaAPIRoutes)

	// Listen for incoming HTTP/1.x & HTTP/2 clients on localhost port 8080.
	// TODO: get port from config file.
	irisApp.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"), iris.WithoutVersionChecker)
}
