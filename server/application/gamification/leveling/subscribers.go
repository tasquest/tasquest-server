package leveling

import (
	log "github.com/sirupsen/logrus"

	"tasquest.com/server/application/gamification/tasks"
	"tasquest.com/server/infra/events"
)

type Subscribers struct {
	subscriber         events.Subscriber
	progressionService ProgressionService
}

func (subs Subscribers) toTaskCompletedEvent() (events.Subscription, error) {
	subscription, err := subs.subscriber.Subscribe(tasks.AdventurerTaskTopic, subs.progressionService.AwardExperience)

	if err != nil {
		log.Error("leveling.Subscribers.toTaskCompletedEvent failed to subscribe! Cause: " + err.Error())
	}

	return subscription, err
}
