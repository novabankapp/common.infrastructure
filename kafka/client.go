package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func NewKafkaClient(ctx context.Context, config Config) (*kafka.Conn, error) {
	return kafka.DialContext(ctx, "tcp", config.Brokers[0])
}
