cd ..
mockgen -source=bvc/domain.go -destination=bvc/mock_domain.go -package=bvc
mockgen -source=bvc/client.go -destination=bvc/mock_client.go -package=bvc
mockgen -source=bvc/persistence.go -destination=bvc/mock_persistence.go -package=bvc
