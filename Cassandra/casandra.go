package Cassandra

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"log"
)

func NewCassandraSession(config Config) (*gocqlx.Session, error) {
	cluster := gocql.NewCluster(config.Addresses...)
	cluster.Keyspace = config.Keyspace
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = config.ProtoVersion
	cluster.ConnectTimeout = config.Timeout
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: config.Username,
		Password: config.Password}
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &session, nil
}
