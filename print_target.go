package logger

import "fmt"

type PrintLogTarget struct {
	BaseLogTarget
}

func (l *PrintLogTarget) Log(message MessageData) error {
	fmt.Printf("Call log message in PrintLogTarget \n message: %v \n extraData %s  \n level %v \n category %v \n",
		message.Data(), message.ExtraData(), message.Level(), message.Category(),
	)

	return nil
}
