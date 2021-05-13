package application

import (
	"sync"

	"tasquest.com/server/application/gamification/leveling"
	"tasquest.com/server/infra/events"
)

var IsSubscriptionManagementInstanced sync.Once
var subscriptionManagementInstance *SubscriptionManagement

type SubscriptionManagement struct {
	levelingSubscribers *leveling.Subscribers
	subscriptions       []events.Subscription
}

func NewSubscriptionManagement(
	levelingSubscribers *leveling.Subscribers,
) *SubscriptionManagement {
	IsSubscriptionManagementInstanced.Do(func() {
		subscriptionManagementInstance = &SubscriptionManagement{
			levelingSubscribers: levelingSubscribers,
		}
		subscriptionManagementInstance.startSubscriptions()
	})
	return subscriptionManagementInstance
}

func (sm SubscriptionManagement) startSubscriptions() {
	if taskCompleteSubscription, err := sm.levelingSubscribers.ToTaskCompletedEvent(); err == nil {
		sm.subscriptions = append(sm.subscriptions, taskCompleteSubscription)
	}
}

func init() {
	subscriptionManagementInstance.startSubscriptions()
}
