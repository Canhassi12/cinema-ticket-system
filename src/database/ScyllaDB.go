package database

import (
	"context"

	"github.com/scylladb/scylla-go-driver"
)

type ScyllaConn struct{}

type ScyllaConnInterface interface {
	Conn() *scylla.Session
}

func (ScyllaConn) Conn() *scylla.Session {
	ctx := context.Background()

	cfg := scylla.DefaultSessionConfig("go", "127.0.0.1:8000")
	session, err := scylla.NewSession(ctx, cfg)

	if err != nil {
		panic("DEU ERRO AUQIO CNHASS PAMIGAOSD")
	}

	return session
}
