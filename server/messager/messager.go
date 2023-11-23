package messager

// Messager 消息推送
type Messager interface {
	Push(topic, msg string) error
}
