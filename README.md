# mila-api

An API to provide MILA (Mercado Integrado Latinoamericano) stock data.

Status: [![CircleCI](https://circleci.com/gh/julianespinel/mila-api.svg?style=svg&circle-token=65f92f8bd064930b35681d97f699bf2707f10d3e)](https://circleci.com/gh/julianespinel/mila-api)

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
