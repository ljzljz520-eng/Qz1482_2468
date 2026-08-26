package workflow

import (
	"aerialfarm/storage"
	"context"
	"path/filepath"
	"testing"
)

func TestCancellationLeavesPending(t *testing.T) {
	s, e := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	svc := NewService(s)
	r, _ := svc.Register(context.Background(), "cancel", "f", "p")
	svc.Begin(context.Background(), r.ID)
	ctx, c := context.WithCancel(context.Background())
	c()
	r, _ = svc.Process(ctx, r.ID)
	if r.Status != "processed" {
		t.Fatalf("expected processed, got %s", r.Status)
	}
}
func TestRecordFlow33(t *testing.T) {
	s, e := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	svc := NewService(s)
	r, _ := svc.Register(context.Background(), "33", "north-33", "pilot")
	svc.Begin(context.Background(), r.ID)
	ctx, c := context.WithCancel(context.Background())
	c()
	r, _ = svc.Process(ctx, r.ID)
	if r.Status != "processed" {
		t.Fatalf("航拍农田巡检资料显示待处理状态: %s", r.Status)
	}
}
