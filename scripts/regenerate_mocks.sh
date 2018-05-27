cd ..
mockgen -source=core/domain.go -destination=core/mock_domain.go -package=core
mockgen -source=core/persistence.go -destination=core/mock_persistence.go -package=core
mockgen -source=bvc/client.go -destination=bvc/mock_client.go -package=bvc
