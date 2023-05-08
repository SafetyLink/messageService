package scyllaDB

import (
	"fmt"
	"github.com/SafetyLink/commons/config"
	"github.com/SafetyLink/commons/logger"
	internal "github.com/SafetyLink/messageService/internal/infra"
	"testing"
)

func TestNewScyllaProvider(t *testing.T) {
	log := logger.InitLogger()
	cfg, err := config.ReadConfigInTest[internal.Config]()
	if err != nil {
		t.Error(err)
	}

	session := NewScyllaProvider(log, cfg)

	fmt.Println(session)
}
