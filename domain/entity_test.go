package domain

import (
	"testing"
	"time"
)

func TestRecordConstruction(t *testing.T) {
	r := NewRecord("r1", "f1", "p1", time.Unix(1, 0))
	if r.Status != "received" {
		t.Fatal(r)
	}
}
