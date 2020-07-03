package cassandra

import (
	"github.com/gocql/gocql"
)

const (
	Keyspace = "oauth"
	Host     = "127.0.0.1"
)

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster(Host)
	cluster.Keyspace = Keyspace
	cluster.Consistency = gocql.Quorum
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
