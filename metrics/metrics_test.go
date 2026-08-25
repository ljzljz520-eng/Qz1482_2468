package metrics

import (
	"aerialfarm/domain"
	"testing"
)

func TestMetricsHealth(t *testing.T) {
	var c Counter
	c.Observe(domain.Record{Status: "pending"})
	if c.Health() != "attention" {
		t.Fatal(c.Health())
	}
}
