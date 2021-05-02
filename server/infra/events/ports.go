package events

type Publisher interface {
	Publish(topic string, event interface{}) (interface{}, error)
}

type Subscription interface {
	Unsubscribe() error
}

type Subscriber interface {
	Subscribe(topic string, exec func()) (Subscription, error)
}
