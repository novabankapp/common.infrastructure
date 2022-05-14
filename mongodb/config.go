package mongodb

type Config struct {
	URI      string `mapstructure:"uri" yaml:"uri"`
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
	Db       string `mapstructure:"db" yaml:"db"`
}
