package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// todo add file rotation (max size + age + backup count)
// todo add file mutex based on file lock (create mutex package)

type FileLogTarget struct {
	BaseLogTarget
	FilePath   string
	MaxSize    int
	MaxBackups int
}

func (t *FileLogTarget) Log(message MessageData) error {

	if err := createFileDir(t.FilePath); err != nil {
		log.Fatalf("error creating directory: %v", err)

		return err
	}

	f, err := openFile(t.FilePath)

	if err != nil {
		log.Fatalf("error opening file: %v", err)

		return err
	}

	defer f.Close()

	logger := log.New(
		io.MultiWriter(os.Stderr, f),
		composeLogPrefix(message.Level(),
			message.Category()),
		log.Ldate|log.Ltime)
	logger.Printf("%s\n %s\n", message, stack())

	return nil
}

func composeLogPrefix(level string, category string) string {
	logPrefix := ""

	if level != "" {
		logPrefix = level
	}

	if category != "" {
		if logPrefix != "" {
			logPrefix += " / "
		}

		logPrefix += category
	}

	if logPrefix != "" {
		return logPrefix + " "
	}

	return ""
}

func openFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	return f, err
}

func createFileDir(filePath string) error {
	fileDir := filepath.Dir(filePath)

	if err := os.MkdirAll(fileDir, os.ModePerm); err != nil {
		return err
	}

	return nil
}

//TODO подумать над этим делом.. выглядит как костыль, но иначе прилетает и ненужный мусор
func stack() string {
	buf := make([]byte, 1<<16)
	stackSize := runtime.Stack(buf, true)

	return string(buf[524:stackSize])
}
