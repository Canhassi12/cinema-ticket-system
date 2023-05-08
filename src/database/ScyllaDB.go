package database

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type ScyllaConn struct{}

type ScyllaConnInterface interface {
	Conn() *gocqlx.Session
}

func (ScyllaConn) Conn() *gocqlx.Session {
	cluster := gocql.NewCluster("127.0.0.1:8000")

	session, err := gocqlx.WrapSession(cluster.CreateSession())

	if err != nil {
		panic(err.Error())
	}

	return &session
}
