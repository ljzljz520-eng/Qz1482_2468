package storage

import (
	"aerialfarm/domain"
	"path/filepath"
	"testing"
	"time"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "farm.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := domain.NewRecord("persist", "f", "p", time.Now())
	if e = s.PutRecord(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.GetRecord("persist")
	if e != nil || got.ID != "persist" {
		t.Fatalf("%v %v", got, e)
	}
}
