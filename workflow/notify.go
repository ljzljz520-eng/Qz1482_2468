package workflow

import (
	"aerialfarm/domain"
	"fmt"
)

type Notification struct{ RecordID, Channel, Message string }

func BuildNotification(r domain.Record) Notification {
	channel := "dashboard"
	if r.Status == "pending" {
		channel = "alert"
	}
	return Notification{RecordID: r.ID, Channel: channel, Message: fmt.Sprintf("%s is %s", r.FieldID, r.Status)}
}
func (s *Service) Notify(r domain.Record) error {
	n := BuildNotification(r)
	return s.Store.PutEvent(domain.NewEvent("event-"+r.ID, r.ID, "notification", n.Message, s.clock()))
}
func NotificationPriority(r domain.Record) int {
	if r.Status == "pending" {
		return 3
	}
	if r.Status == "processed" {
		return 1
	}
	return 2
}
