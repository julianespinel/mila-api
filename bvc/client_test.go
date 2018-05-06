package bvc

import (
	"crypto/tls"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/assert"
)

func Test_BVCClient_GetStocksClosingDataByDate_success(t *testing.T) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	recording, err := recorder.NewAsMode("fixtures/bvc", recorder.ModeReplaying, transport)
	if err != nil {
		log.Fatal(err)
	}
	defer recording.Stop()
	httpClient := &http.Client{Transport: recording}
	bvcClient := InitClient(httpClient)
	date := time.Date(2018, time.April, 30, 0, 0, 0, 0, time.UTC)

	stocks := bvcClient.GetStocksClosingDataByDate(date)
	assert.NotNil(t, stocks)
	assert.NotZero(t, len(stocks))
}
