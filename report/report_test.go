package report

import (
	"aerialfarm/domain"
	"testing"
)

func TestSummary(t *testing.T) {
	s := Summarize([]domain.Record{{Status: "processed"}, {Status: "pending"}})
	if s.Processed != 1 || s.Pending != 1 {
		t.Fatal(s)
	}
	if StatusLabel("processed") != "已处理" {
		t.Fatal("label")
	}
}
