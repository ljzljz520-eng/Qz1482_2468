package storage

import (
	"aerialfarm/domain"
	"strings"
)

func FilterRecords(rs []domain.Record, field, status string) []domain.Record {
	out := make([]domain.Record, 0)
	for _, r := range rs {
		if field != "" && r.FieldID != field {
			continue
		}
		if status != "" && r.Status != status {
			continue
		}
		out = append(out, r)
	}
	return out
}
func SearchNotes(rs []domain.Record, term string) []domain.Record {
	term = strings.ToLower(term)
	out := []domain.Record{}
	for _, r := range rs {
		if strings.Contains(strings.ToLower(r.Notes), term) {
			out = append(out, r)
		}
	}
	return out
}
