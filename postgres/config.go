package postgres

import "fmt"

type Config struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"post"`
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	SslMode  bool   `yaml:"sslMode"`
	Timezone string `yaml:"timezone"`
}

func GetDSN(config Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.Database, config.Port, config.SslMode, config.Timezone)
}
