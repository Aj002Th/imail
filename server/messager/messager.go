package messager

// 消息推送
type Messager interface {
	Push(msg string) error
}
