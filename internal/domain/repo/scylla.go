package repo

import (
	"context"
	"github.com/SafetyLink/commons/types"
	scyllaRepoType "github.com/SafetyLink/messageService/internal/infra/scyllaRepository/types"
)

type ScyllaRepository interface {
	CreateMessage(ctx context.Context, message *types.Message, chatID int64) error
	GetMessageByMessageID(ctx context.Context, messageID int64) (*scyllaRepoType.Message, error)
	GetMessageByChatID(ctx context.Context, chatID int64) ([]*scyllaRepoType.Message, error)
}
