package logger

import (
	"github.com/liana-go/queue"
	"github.com/thoas/go-funk"
)

// todo uid to target
// todo make except by target uid
// todo add uid to every message
// todo add previous message uid

const LevelDebug = "debug"
const LevelInfo = "info"
const LevelWarning = "warning"
const LevelError = "error"
const LevelCritical = "critical"

var LogsQueueName string
var logger *Logger
var worker queue.QueueWorker

func init() {
	LogsQueueName = "logs"
	logger = &Logger{
		Targets: make([]Target, 0),
	}
}

func New(consumingThreads int, messagesMaxCount int) *Logger {
	if consumingThreads <= 0 {
		consumingThreads = 1
	}
	if messagesMaxCount <= 0 {
		messagesMaxCount = 10
	}

	worker = queue.QueueWorker{
		QueueName:  LogsQueueName,
		Broker:     queue.NewMemoryBroker(messagesMaxCount),
		Callable:   logger.logQueueMessage,
		IsInfinite: false,
	}

	go worker.Run()

	return logger
}

// Logger TODO add FlushInterval and batch logging
type Logger struct {
	Targets []Target
}

func (l *Logger) Log(data interface{}, extraData map[string]interface{}, level string, category string) {
	message := &Message{
		data:      data,
		extraData: extraData,
		level:     level,
		category:  category,
	}

	_ = worker.Broker.Publish(LogsQueueName, message, nil)
}

func (l *Logger) Debug(data interface{}, extraData map[string]interface{}, category string) {
	l.Log(data, extraData, LevelDebug, category)
}

func (l *Logger) Info(data interface{}, extraData map[string]interface{}, category string) {
	l.Log(data, extraData, LevelInfo, category)
}

func (l *Logger) Warning(data interface{}, extraData map[string]interface{}, category string) {
	l.Log(data, extraData, LevelWarning, category)
}

func (l *Logger) Error(data interface{}, extraData map[string]interface{}, category string) {
	l.Log(data, extraData, LevelError, category)
}

func (l *Logger) Critical(data interface{}, extraData map[string]interface{}, category string) {
	l.Log(data, extraData, LevelCritical, category)
}

func (l *Logger) logMessage(message MessageData) {
	for k, target := range l.Targets {
		if !funk.Contains(message.Except(), k) && target.CanLog(message) {
			if err := target.Log(message); err != nil {
				l.logMessage(composeFailLoggingMessage(message, err, k))
			}
		}
	}
}

func (l *Logger) AddTarget(target *Target) {
	l.Targets = append(l.Targets, *target)
}

func (l *Logger) logQueueMessage(data queue.MessageData) {
	// TODO add error on fail interface convertation
	message := data.Data().(MessageData)

	l.logMessage(message)
}

func composeFailLoggingMessage(message MessageData, err error, except int) MessageData {
	// todo make except by target uid
	message.AddExceptTarget(except)

	return &Message{
		data: map[string]interface{}{
			"errorMessage":    "Error on trying to log data",
			"error":           err,
			"originalMessage": message.Data(),
		},
		level:    LevelCritical,
		except:   message.Except(),
		previous: []MessageData{message},
	}
}
