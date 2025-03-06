package services

type EventBus interface {
	Publish(topic string, data interface{}) error
	Subscribe(topic string, handler func(data interface{})) error
}
