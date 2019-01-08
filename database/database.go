package db

import (
	"github.com/gocql/gocql"
)

var session *gocql.Session

func Connect() *gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "dev"
	cluster.Consistency = gocql.Quorum

	session, _ = cluster.CreateSession()

	return session
}

func GetSession() *gocql.Session {
	return session
}
