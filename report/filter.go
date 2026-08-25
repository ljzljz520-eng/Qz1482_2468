package report

import "aerialfarm/domain"

func PendingOnly(rs []domain.Record) []domain.Record {
	out := []domain.Record{}
	for _, r := range rs {
		if r.Status == "pending" {
			out = append(out, r)
		}
	}
	return out
}
func ProcessedOnly(rs []domain.Record) []domain.Record {
	out := []domain.Record{}
	for _, r := range rs {
		if r.Status == "processed" {
			out = append(out, r)
		}
	}
	return out
}
func GroupByField(rs []domain.Record) map[string][]domain.Record {
	out := map[string][]domain.Record{}
	for _, r := range rs {
		out[r.FieldID] = append(out[r.FieldID], r)
	}
	return out
}
