package kafka

type Config struct {
	Brokers    []string `mapstructure:"brokers"`
	GroupID    string   `mapstructure:"groupID"`
	InitTopics bool     `mapstructure:"initTopics"`
}

// TopicConfig kafka topic config
type TopicConfig struct {
	TopicName         string `mapstructure:"topicName"yaml:"topicName"`
	Partitions        int    `mapstructure:"partitions"yaml:"partitions"`
	ReplicationFactor int    `mapstructure:"replicationFactor"yaml:"replicationFactor"`
}
