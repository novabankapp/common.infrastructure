package Cassandra

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/psanford/memfs"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/migrate"
	"io/fs"
	"log"
	"os"
	"strconv"
)

func NewCassandraSession(config Config) (*gocqlx.Session, error) {
	cluster := gocql.NewCluster(config.Addresses...)
	sess, e := cluster.CreateSession()
	if e != nil {
		log.Println(e)
		return nil, e
	}
	createKeyspace(config.Keyspace, sess, config.ReplicationFactor)
	cluster.Keyspace = config.Keyspace
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = config.ProtoVersion
	cluster.Timeout = config.Timeout
	//cluster.ConnectTimeout = config.Timeout
	//cluster.Authenticator = gocql.PasswordAuthenticator{Username: config.Username,
	//	Password: config.Password}
	session, err := gocqlx.WrapSession(sess, e)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &session, nil
}

func createKeyspace(keyspace string, session *gocql.Session, replicationFactor int) {
	query := "CREATE KEYSPACE IF NOT EXISTS " + keyspace + " WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : " + strconv.FormatInt(int64(replicationFactor), 10) + " };"
	if err := session.Query(query).Exec(); err != nil {
		log.Fatal("--FATAL--", err, query, strconv.FormatInt(int64(replicationFactor), 10))
	}
	log.Println("INFO", "Configuration keyspace: "+keyspace)
}

func MigrateFromFile(context context.Context, session gocqlx.Session, filename string) {
	f := memfs.New()
	content, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	f.WriteFile(fmt.Sprint("query", ".cql"), content, fs.ModePerm)
	migrate.FromFS(context, session, f)
}
func MigrateFromQueryString(context context.Context, session gocqlx.Session, text string) {
	f := memfs.New()
	f.WriteFile(fmt.Sprint("query", ".cql"), []byte(text), fs.ModePerm)
	migrate.FromFS(context, session, f)

}
