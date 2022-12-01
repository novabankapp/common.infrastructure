package Cassandra

import "time"

type Config struct {
	Username          string
	Password          string
	Addresses         []string
	Keyspace          string
	ProtoVersion      int
	ReplicationFactor int
	Timeout           time.Duration
}
