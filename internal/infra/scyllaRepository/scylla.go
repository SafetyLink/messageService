package scyllaRepository

import (
	"github.com/SafetyLink/messageService/internal/domain/repo"
	internal "github.com/SafetyLink/messageService/internal/infra"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
	"go.uber.org/zap"
)

type ScyllaRepository struct {
	scyllaSession gocqlx.Session
	logger        *zap.Logger
	config        *internal.Config
	messageTable  *table.Table
}

func NewScyllaRepository(logger *zap.Logger, cfg *internal.Config, scyllaSession gocqlx.Session) repo.ScyllaRepository {
	err := scyllaSession.ExecStmt(`CREATE TABLE IF NOT EXISTS message (
	message_id bigint,
	chat_id bigint,
	author_id bigint,
	content text,
	attachment_id bigint,
	attachment_file_name varchar,
	attachment_is_screenshot boolean,
	attachment_type varchar,
	attachment_ttl boolean,
	attachment_after_view boolean,
	attachment_time_limit timestamp,
	ttl boolean,
	after_view  boolean,
	time_limit timestamp,
	edited boolean,
	edited_at timestamp,
	pinned boolean,
	pinned_at timestamp,
	referenced_message_id bigint,
	is_viewed boolean,
	is_screenshot boolean,
	created_at timestamp,
	PRIMARY KEY (message_id,chat_id)
)`)
	if err != nil {
		logger.Error("Error creating table", zap.Error(err))
	}

	messageMetadata := table.Metadata{
		Name: "message",
		Columns: []string{"message_id", "chat_id", "author_id", "content", "attachment_id",
			"attachment_file_name", "attachment_is_screenshot", "attachment_type", "attachment_ttl",
			"attachment_after_view", "attachment_time_limit", "ttl", "after_view", "time_limit", "edited",
			"edited_at", "pinned", "pinned_at", "referenced_message_id", "is_viewed", "is_screenshot",
			"created_at"},
		PartKey: []string{"chat_id"},
		SortKey: []string{"message_id", "created_at"},
	}

	messageTable := table.New(messageMetadata)

	return &ScyllaRepository{
		scyllaSession: scyllaSession,
		logger:        logger,
		config:        cfg,
		messageTable:  messageTable,
	}

}
