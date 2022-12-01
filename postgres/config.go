package postgres

import "fmt"

type Config struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"post"`
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	SslMode  string `yaml:"sslMode"`
	Timezone string `yaml:"timezone"`
}

func GetDSN(config Config) string {
	fmt.Println(config.Timezone)
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		config.Host, config.User, config.Password, config.Database, config.Port, config.SslMode)
}
