package types

import "time"

type Message struct {
	MessageID              int64     `json:"message_id"`
	ChatID                 int64     `json:"chat_id"`
	AuthorID               int64     `json:"author_id"`
	Content                string    `json:"content"`
	AttachmentID           int64     `json:"attachment_id"`
	AttachmentFileName     string    `json:"attachment_file_name"`
	AttachmentIsScreenshot bool      `json:"attachment_is_screenshot"`
	AttachmentType         string    `json:"attachment_type"`
	AttachmentTTL          bool      `json:"attachment_ttl"`
	AttachmentAfterView    bool      `json:"attachment_after_view"`
	AttachmentTimeLimit    time.Time `json:"attachment_time_limit"`
	TTL                    bool      `json:"ttl"`
	AfterView              bool      `json:"after_view"`
	TimeLimit              time.Time `json:"time_limit"`
	Edited                 bool      `json:"edited"`
	EditedAt               time.Time `json:"edited_at"`
	Pinned                 bool      `json:"pinned"`
	PinnedAt               time.Time `json:"pinned_at"`
	ReferencedMessageID    int64     `json:"referenced_message_id"`
	IsViewed               bool      `json:"is_viewed"`
	IsScreenshot           bool      `json:"is_screenshot"`
	CreatedAt              time.Time `json:"created_at"`
}
