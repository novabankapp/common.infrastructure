package kafka

type Config struct {
	Brokers     []string    `mapstructure:"brokers"`
	GroupID     string      `mapstructure:"groupID"`
	InitTopics  bool        `mapstructure:"initTopics"`
	KafkaTopics KafkaTopics `mapstructure:"kafkaTopics"`
	Try         string
}

// TopicConfig kafka topic config
type TopicConfig struct {
	TopicName         string `mapstructure:"topicName" yaml:"topicName"`
	Partitions        int    `mapstructure:"partitions" yaml:"partitions"`
	ReplicationFactor int    `mapstructure:"replicationFactor" yaml:"replicationFactor"`
}

type KafkaTopics struct {
	UserCreated         TopicConfig `mapstructure:"userCreated"`
	UserUpdated         TopicConfig `mapstructure:"userUpdated"`
	UserDeleted         TopicConfig `mapstructure:"userDeleted"`
	UserPasswordChanged TopicConfig `mapstructure:"userPasswordChanged"`
	UserLocked          TopicConfig `mapstructure:"userLocked"`
}
