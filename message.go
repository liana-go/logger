package logger

type MessageData interface {
	Data() interface{}
	ExtraData() map[string]interface{}
	Level() string
	Category() string
	Except() []int
	AddExceptTarget(key int)
	Previous() []MessageData
	AddPrevious(MessageData)
}

type Message struct {
	data      interface{}
	extraData map[string]interface{}
	level     string
	category  string
	except    []int
	previous  []MessageData
}

func (l *Message) Data() interface{} {
	return l.data
}

func (l *Message) ExtraData() map[string]interface{} {
	return l.extraData
}

func (l *Message) Level() string {
	return l.level
}

func (l *Message) Category() string {
	return l.category
}

func (l *Message) Except() []int {
	return l.except
}

func (l *Message) AddExceptTarget(key int) {
	l.except = append(l.except, key)
}

func (l *Message) AddPrevious(message MessageData) {
	l.previous = append(l.previous, message)
}

func (l *Message) Previous() []MessageData {
	return l.previous
}
