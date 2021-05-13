package leveling

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"tasquest.com/server/application/gamification/tasks"
	"tasquest.com/server/infra/events"
)

var IsSubscribersInstanced sync.Once
var subscribersInstance *Subscribers

type Subscribers struct {
	subscriber         events.Subscriber
	progressionService ProgressionService
}

func NewLevelingSubscribers(
	eventSubscriber events.Subscriber,
	service ProgressionService,
) *Subscribers {
	IsSubscribersInstanced.Do(func() {
		subscribersInstance = &Subscribers{
			subscriber:         eventSubscriber,
			progressionService: service,
		}
	})
	return subscribersInstance
}

func (subs Subscribers) ToTaskCompletedEvent() (events.Subscription, error) {
	subscription, err := subs.subscriber.Subscribe(tasks.AdventurerTaskTopic, subs.progressionService.AwardExperience)

	if err != nil {
		log.Error("leveling.Subscribers.toTaskCompletedEvent failed to subscribe! Cause: " + err.Error())
	}

	return subscription, err
}
