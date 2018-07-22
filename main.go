package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julianespinel/mila-api/admin"
	"github.com/julianespinel/mila-api/bvc"
	"github.com/julianespinel/mila-api/core"
	"github.com/julianespinel/mila-api/models"
	"github.com/kataras/iris"
	"github.com/robfig/cron"
)

func initializeBVCClient() bvc.MilaClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	return bvc.InitClient(httpClient)
}

func initializeCore(db *gorm.DB) core.API {
	bvcClient := initializeBVCClient()
	persistence := core.InitPersistence(db)
	domain := core.InitDomain(bvcClient, persistence)
	api := core.InitAPI(domain)
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

func initializeDBConnection(dbConfig models.DatabaseConfig) *gorm.DB {
	dbURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
		dbConfig.Charset,
	)
	db, err := gorm.Open(dbConfig.Dialect, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func readConfigFromFile(configFilePath string) models.Config {
	var config models.Config
	if _, err := toml.DecodeFile(configFilePath, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

func main() {
	configFilePath := flag.String(
		"config",                               // flag name
		"config/development.toml",              // default value
		"File path of the configuration file.", // usage
	)
	config := readConfigFromFile(*configFilePath)

	db := initializeDBConnection(config.Database)
	milaAPI := initializeCore(db)
	gocron.Every(1).Day().At("23:00").Do(milaAPI.UpdateDailyStocks, time.Now())

	// Start all the pending jobs
	<-gocron.Start()

	// Initialize Iris app
	irisApp := initializeIrisApp()
	// Define Iris routes
	milaAdminRoutes := irisApp.Party("/mila/admin", logURLAndIP)
	milaAdminRoutes = admin.AddRoutes(milaAdminRoutes)

	milaAPIRoutes := irisApp.Party("/mila/api", logURLAndIP)
	milaAPIRoutes = milaAPI.AddRoutes(milaAPIRoutes)

	// Listen for incoming HTTP/1.x & HTTP/2 clients.
	irisAddress := fmt.Sprintf(":%d", config.Web.Port)
	irisApp.Run(
		iris.Addr(irisAddress),
		iris.WithCharset(config.Web.Charset),
		iris.WithoutVersionChecker,
	)
}
