package query

import (
	"aerialfarm/domain"
	"aerialfarm/storage"
	"path/filepath"
	"testing"
	"time"
)

func TestQueryFiltering(t *testing.T) {
	s, e := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	s.PutRecord(domain.NewRecord("q1", "a", "p", time.Now()))
	s.PutRecord(domain.NewRecord("q2", "b", "p", time.Now()))
	q := New(s)
	rs, e := q.Find("a", "")
	if e != nil || len(rs) != 1 {
		t.Fatal(rs, e)
	}
}
