package redis

type Config struct {
	Addr     string `mapstructure:"addr"yaml:"addr"`
	Password string `mapstructure:"password"yaml:"password"`
	DB       int    `mapstructure:"db"yaml:"db"`
	PoolSize int    `mapstructure:"poolSize"yaml:"PoolSize"`
}
