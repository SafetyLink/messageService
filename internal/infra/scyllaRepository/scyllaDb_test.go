package scyllaRepository

import (
	"context"
	"fmt"
	"github.com/SafetyLink/commons/config"
	"github.com/SafetyLink/commons/logger"
	"github.com/SafetyLink/commons/snowflake"
	"github.com/SafetyLink/commons/types"
	internal "github.com/SafetyLink/messageService/internal/infra"
	"github.com/SafetyLink/messageService/internal/infra/adapters/scyllaDB"
	"testing"
	"time"
)

var ChatID int64 = 110334622035607552

func TestScyllaCreateMessage(t *testing.T) {
	log := logger.InitLogger()
	cfg, err := config.ReadConfigInTest[internal.Config]()
	if err != nil {
		t.Error(err)
	}

	session := scyllaDB.NewScyllaProvider(log, cfg)
	ScyllaRepo := NewScyllaRepository(log, cfg, session)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	generateMessageID := snowflake.GenerateSnowflakeID()
	authorID := 1

	messagePayload := &types.Message{
		MessageID: generateMessageID,
		AuthorID:  int64(authorID),
		Content:   "Hello world!",
		Attachment: &types.Attachment{
			AttachmentID: 0,
			FileName:     "",
			IsScreenshot: false,
			Type:         "",
			TTL: &types.AttachmentTTL{
				TTL:       false,
				AfterView: false,
				TimeLimit: time.Time{},
			},
		},
		TTL: &types.TTL{
			TTL:       false,
			AfterView: false,
			TimeLimit: time.Time{},
		},
		Edited:              false,
		EditedAt:            time.Time{},
		Pinned:              false,
		PinnedAt:            time.Time{},
		ReferencedMessageID: 0,
		IsViewed:            false,
		IsScreenshot:        false,
		CreatedAt:           time.Now(),
	}

	err = ScyllaRepo.CreateMessage(ctx, messagePayload, ChatID)
	if err != nil {
		t.Error(err)
	}

	queryMessage, err := ScyllaRepo.GetMessageByMessageID(ctx, generateMessageID)
	if err != nil {
		t.Error(err)
	}

	t.Logf(fmt.Sprintf("Your message >> %v", queryMessage))
	return
}

func TestGetMessageByChatID(t *testing.T) {
	log := logger.InitLogger()
	cfg, err := config.ReadConfigInTest[internal.Config]()
	if err != nil {
		t.Error(err)
	}

	session := scyllaDB.NewScyllaProvider(log, cfg)
	ScyllaRepo := NewScyllaRepository(log, cfg, session)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	queryMessage, err := ScyllaRepo.GetMessageByChatID(ctx, ChatID)
	if err != nil {
		t.Error(err)
	}

	for _, message := range queryMessage {
		t.Logf(fmt.Sprintf("Your message >> %v", message))
	}
	return

}
