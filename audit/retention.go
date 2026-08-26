package audit

import (
	"aerialfarm/domain"
	"time"
)

func Recent(events []domain.Event, since time.Time) []domain.Event {
	out := []domain.Event{}
	for _, e := range events {
		if !e.At.Before(since) {
			out = append(out, e)
		}
	}
	return out
}
func ActionCount(a []domain.Audit) map[string]int {
	out := map[string]int{}
	for _, v := range a {
		out[v.Action]++
	}
	return out
}
func ValidActor(a domain.Audit) bool { return a.Actor != "" && a.RecordID != "" }
