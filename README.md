# mila-api

An API to provide MILA (Mercado Integrado Latinoamericano) stock data.

Status: [![CircleCI](https://circleci.com/gh/julianespinel/mila-api.svg?style=svg&circle-token=65f92f8bd064930b35681d97f699bf2707f10d3e)](https://circleci.com/gh/julianespinel/mila-api)

## Installation

Please follow these steps:

1. `git clone git@github.com:julianespinel/mila-api.git`
1. `cd mila-api/`
1. `docker-compose up --build`

Check the server is running:
```
curl http://localhost:3000/mila/admin/ping
```

If the server is up and running you should get JSON in the response body:
```json
{"message":"pong"}
```

## Testing

Run the tests:
1. `go test ./...`

Check code coverage:
1. Generate coverage report: `go test -coverprofile cover.out ./...`
1. Translate coverage report to HTML: `go tool cover -html=cover.out -o cover.html`
1. Open the file `cover.html` with a web browser.

Create new mocks:
1. `mockgen -source=file_name.go -destination=mock_file_name.go -package=destination_package_name`

Regenerate mocks:
1. `cd scripts`
1. `sh regenerate_mocks.sh`

## Usage

Currently the system performs the following operation:

1. Every day at 23:00 the system runs a job responsible for updating the stock data with the information of the last trading day from the Colombian stock market.

After the job runs you can retrieve the stock data with this request:
```
curl http://localhost:3000/mila/api/colombia
```

The response should look like this:
```json
{
  "date": "2018-07-22T22:43:19.640894256Z",
  "country": "colombia",
  "stocksData": [{
    "date": "2018-07-19T22:42:55Z",
    "country": "colombia",
    "symbol": "ECOPETROL ",
    "name": "ECOPETROL ",
    "currency": "cop",
    "open": "0",
    "high": "0",
    "low": "0",
    "close": "2950",
    "adjClose": "0",
    "volume": 5021416
  }, {
    "date": "2018-07-19T22:42:55Z",
    "country": "colombia",
    "symbol": "PFBCOLOM  ",
    "name": "PFBCOLOM  ",
    "currency": "cop",
    "open": "0",
    "high": "0",
    "low": "0",
    "close": "34540",
    "adjClose": "0",
    "volume": 281086
  }]
}
```
