package report

import (
	"aerialfarm/domain"
	"aerialfarm/storage"
	"sort"
)

type Summary struct{ Total, Processed, Pending, Archived int }

func Summarize(rs []domain.Record) Summary {
	var s Summary
	for _, r := range rs {
		s.Total++
		switch r.Status {
		case "processed":
			s.Processed++
		case "pending":
			s.Pending++
		case "archived":
			s.Archived++
		}
	}
	return s
}
func FieldReport(rs []domain.Record, field string) Summary {
	return Summarize(storage.FilterRecords(rs, field, ""))
}
func Latest(rs []domain.Record, n int) []domain.Record {
	sort.Slice(rs, func(i, j int) bool { return rs[i].CapturedAt.After(rs[j].CapturedAt) })
	if n > len(rs) {
		n = len(rs)
	}
	return rs[:n]
}
func CompletionRate(s Summary) float64 {
	if s.Total == 0 {
		return 0
	}
	return float64(s.Processed) / float64(s.Total)
}
