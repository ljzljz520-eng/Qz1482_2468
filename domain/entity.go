package domain

import "time"

type Record struct {
	ID, FieldID, Pilot string
	Status             string
	CapturedAt         time.Time
	Notes              string
}
type User struct {
	ID, Name, Role string
	Active         bool
}
type Event struct {
	ID, RecordID, Kind, Detail string
	At                         time.Time
}
type Audit struct {
	ID, Actor, Action, RecordID string
	At                          time.Time
}

func NewRecord(id, field, pilot string, at time.Time) Record {
	return Record{ID: id, FieldID: field, Pilot: pilot, Status: "received", CapturedAt: at}
}
func (r Record) IsProcessed() bool          { return r.Status == "processed" }
func (r *Record) MarkProcessed(note string) { r.Status = "processed"; r.Notes = note }
func (r *Record) MarkPending(reason string) { r.Status = "pending"; r.Notes = reason }
func NewEvent(id, rid, kind, detail string, at time.Time) Event {
	return Event{ID: id, RecordID: rid, Kind: kind, Detail: detail, At: at}
}
func NewAudit(id, actor, action, rid string, at time.Time) Audit {
	return Audit{ID: id, Actor: actor, Action: action, RecordID: rid, At: at}
}
