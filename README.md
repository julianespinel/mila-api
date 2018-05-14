# mila-api

An API to provide MILA (Mercado Integrado Latinoamericano) stock data.

Status: [![CircleCI](https://circleci.com/gh/julianespinel/mila-api.svg?style=svg&circle-token=65f92f8bd064930b35681d97f699bf2707f10d3e)](https://circleci.com/gh/julianespinel/mila-api)

## Testing

Run the tests: `go test ./...`

Create new mocks: `mockgen -source=file_name.go -destination=mock_file_name.go -package=destination_package_name`

Regenerate mocks:
1. `cd scripts`
1. `sh regenerate_mocks.sh`
