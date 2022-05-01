package Cassandra

import "time"

type Config struct {
	Username     string
	Password     string
	Addresses    []string
	Keyspace     string
	ProtoVersion int
	Timeout      time.Duration
}
