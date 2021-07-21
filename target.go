package logger

import (
	"github.com/thoas/go-funk"
)

type Target interface {
	Log(message MessageData) error
	CanLog(message MessageData) bool
}

type BaseLogTarget struct {
	Levels     []string
	Categories []string
}

func (t *BaseLogTarget) CanLog(message MessageData) bool {
	if len(t.Levels) > 0 && !funk.Contains(t.Levels, message.Level()) {
		return false
	}

	if len(t.Categories) > 0 && !funk.Contains(t.Categories, message.Category()) {
		return false
	}

	return true
}
