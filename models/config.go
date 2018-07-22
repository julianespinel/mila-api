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
	DBName   string `toml:"db_name"`
	Charset  string
}
