package logger

import (
	"github.com/getsentry/sentry-go"
)

//var mu sync.Mutex
//
//func init() {
//	mu = sync.Mutex{}
//}

type SentryLogTarget struct {
	BaseLogTarget
}

func (l *SentryLogTarget) Log() error {

	sentry.CurrentHub().CaptureMessage("message")
	sentry.CurrentHub().CaptureMessage("message 1")

	sentry.Flush(0)

	sentry.Flush(0)

	return nil
}
