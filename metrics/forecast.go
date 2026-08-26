package metrics

import "aerialfarm/domain"

func StatusVector(rs []domain.Record) map[string]int {
	out := map[string]int{}
	for _, r := range rs {
		out[r.Status]++
	}
	return out
}
func NeedsAttention(c Counter) bool { return c.Pending >= 3 || c.Processed == 0 && c.Received > 0 }
func EstimateHours(c Counter) float64 {
	if c.Received == 0 {
		return 0
	}
	return float64(c.Received-c.Processed) * 1.5
}
