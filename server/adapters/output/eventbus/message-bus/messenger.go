package message_bus

import (
	"sync"

	messagebus "github.com/vardius/message-bus"

	"tasquest.com/server/infra/events"
)

var IsMessengerInstanced sync.Once
var messengerInstance *Messenger

type Messenger struct {
	bus messagebus.MessageBus
}

type MessageSubscription struct {
	bus   messagebus.MessageBus
	topic string
	exec  interface{}
}

func NewMessenger() *Messenger {
	IsMessengerInstanced.Do(func() {
		messengerInstance = &Messenger{bus: messagebus.New(100)}
	})
	return messengerInstance
}

func (m Messenger) Publish(topic string, event interface{}) (interface{}, error) {
	m.bus.Publish(topic, event)
	return event, nil
}

func (m Messenger) Subscribe(topic string, exec interface{}) (events.Subscription, error) {
	err := m.bus.Subscribe(topic, exec)
	return MessageSubscription{
		bus:   m.bus,
		topic: topic,
		exec:  exec,
	}, err
}

func (ms MessageSubscription) Unsubscribe() error {
	return ms.bus.Unsubscribe(ms.topic, ms.exec)
}
