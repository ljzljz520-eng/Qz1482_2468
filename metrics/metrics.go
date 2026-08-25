package metrics

import "aerialfarm/domain"

type Counter struct{ Received, Processed, Pending int }

func (c *Counter) Observe(r domain.Record) {
	switch r.Status {
	case "received", "processing":
		c.Received++
	case "processed":
		c.Processed++
	case "pending":
		c.Pending++
	}
}
func (c Counter) Total() int { return c.Received + c.Processed + c.Pending }
func (c Counter) Health() string {
	if c.Pending > 0 {
		return "attention"
	}
	return "healthy"
}
