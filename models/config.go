package models

type Config struct {
	Web      WebConfig
	Database DatabaseConfig
}

type WebConfig struct {
	Port    int
	Charset string
}

type DatabaseConfig struct {
	Dialect  string
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
	Charset  string
}
