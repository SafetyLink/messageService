package scyllaDB

import (
	internal "github.com/SafetyLink/messageService/internal/infra"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"go.uber.org/zap"
)

func NewScyllaProvider(logger *zap.Logger, config *internal.Config) gocqlx.Session {
	cluster := gocql.NewCluster(config.Scylla.ConnectionUrl)
	cluster.Keyspace = config.Scylla.KeySpace

	session, err := gocqlx.WrapSession(cluster.CreateSession())

	if err != nil {
		logger.Panic("Failed to connect to ScyllaDB", zap.Error(err))
	}
	logger.Info("Connected to ScyllaDB")
	return session
}
