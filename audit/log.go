package audit

import (
	"aerialfarm/domain"
	"aerialfarm/storage"
	"time"
)

type Logger struct {
	Store *storage.Store
	clock func() time.Time
}

func New(s *storage.Store) *Logger { return &Logger{Store: s, clock: time.Now} }
func (l *Logger) Record(actor, action, rid string) error {
	return l.Store.PutAudit(domain.NewAudit("log-"+rid+"-"+action, actor, action, rid, l.clock()))
}
func (l *Logger) Event(rid, kind, detail string) error {
	return l.Store.PutEvent(domain.NewEvent("log-event-"+rid+"-"+kind, rid, kind, detail, l.clock()))
}
