package scyllaRepository

import (
	"context"
	"fmt"
	"github.com/SafetyLink/commons/types"
	scyllaRepoType "github.com/SafetyLink/messageService/internal/infra/scyllaRepository/types"
)

func (sr *ScyllaRepository) CreateMessage(ctx context.Context, message *types.Message, chatID int64) error {
	messageTable := sr.messageTable.InsertQuery(sr.scyllaSession)
	messageTable.BindStruct(scyllaRepoType.Message{
		MessageID:              message.MessageID,
		ChatID:                 chatID,
		AuthorID:               message.AuthorID,
		Content:                message.Content,
		AttachmentID:           message.Attachment.AttachmentID,
		AttachmentFileName:     message.Attachment.FileName,
		AttachmentIsScreenshot: message.Attachment.IsScreenshot,
		AttachmentType:         message.Attachment.Type,
		AttachmentTTL:          message.Attachment.TTL.TTL,
		AttachmentAfterView:    message.Attachment.TTL.AfterView,
		AttachmentTimeLimit:    message.Attachment.TTL.TimeLimit,
		TTL:                    message.TTL.TTL,
		AfterView:              message.TTL.AfterView,
		TimeLimit:              message.TTL.TimeLimit,
		Edited:                 message.Edited,
		EditedAt:               message.EditedAt,
		Pinned:                 message.Pinned,
		PinnedAt:               message.PinnedAt,
		ReferencedMessageID:    message.ReferencedMessageID,
		IsViewed:               message.IsViewed,
		IsScreenshot:           message.IsScreenshot,
		CreatedAt:              message.CreatedAt,
	})
	if err := messageTable.ExecRelease(); err != nil {
		return err
	}

	return nil
}

func (sr *ScyllaRepository) GetMessageByMessageID(ctx context.Context, messageID, chatID int64) (*scyllaRepoType.Message, error) {
	messageTable := sr.messageTable.SelectQuery(sr.scyllaSession)
	fmt.Println(messageID)
	messageTable.BindStruct(&scyllaRepoType.Message{
		MessageID: messageID,
		ChatID:    chatID,
	})
	var messages []*scyllaRepoType.Message

	if err := messageTable.Select(&messages); err != nil {
		return nil, err
	}

	return messages[0], nil
}

func (sr *ScyllaRepository) GetMessageByChatID(ctx context.Context, chatID int64) ([]*scyllaRepoType.Message, error) {
	messageTable := sr.messageTable.SelectQuery(sr.scyllaSession)

	messageTable.BindStruct(&scyllaRepoType.Message{
		ChatID: chatID,
	})
	var messages []*scyllaRepoType.Message

	if err := messageTable.Select(&messages); err != nil {
		return nil, err
	}

	return messages, nil
}
