package workflow

import (
	"aerialfarm/storage"
	"context"
	"path/filepath"
	"testing"
)

func testService(t *testing.T) *Service {
	s, e := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { s.Close() })
	return NewService(s)
}
func TestWorkflowOne(t *testing.T) {
	s := testService(t)
	r, e := s.Register(context.Background(), "w1", "field", "pilot")
	if e != nil {
		t.Fatal(e)
	}
	if _, e = s.Begin(context.Background(), r.ID); e != nil {
		t.Fatal(e)
	}
	r, e = s.Process(context.Background(), r.ID)
	if e != nil || !r.IsProcessed() {
		t.Fatalf("%v %v", r, e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s := testService(t)
	r, _ := s.Register(context.Background(), "w2", "field", "pilot")
	s.Begin(context.Background(), r.ID)
	s.Process(context.Background(), r.ID)
	if _, e := s.Review(r.ID, "reviewer"); e != nil {
		t.Fatal(e)
	}
	if _, e := s.Archive(r.ID); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	s := testService(t)
	r, _ := s.Register(context.Background(), "w3", "field", "pilot")
	s.Begin(context.Background(), r.ID)
	s.Process(context.Background(), r.ID)
	if e := s.Notify(r); e != nil {
		t.Fatal(e)
	}
}
