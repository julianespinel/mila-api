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

func getVCRRecorder(cassetteName string) *recorder.Recorder {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	rec, err := recorder.NewAsMode(
		cassetteName,
		recorder.ModeReplaying,
		transport,
	)
	if err != nil {
		log.Fatal(err)
	}
	return rec
}

func Test_BVCClient_getStocksClosingDataByDate_success(t *testing.T) {
	cassetteName := "fixtures/Test_BVCClient_getStocksClosingDataByDate_success"
	rec := getVCRRecorder(cassetteName)
	defer rec.Stop()
	httpClient := &http.Client{Transport: rec}
	bvcClient := InitClient(httpClient)
	date := time.Date(2018, time.April, 30, 0, 0, 0, 0, time.UTC)

	stocks := bvcClient.getStocksClosingDataByDate(date)
	assert.NotNil(t, stocks)
	assert.NotZero(t, len(stocks))
}
