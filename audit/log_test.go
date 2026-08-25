package audit

import (
	"aerialfarm/domain"
	"aerialfarm/storage"
	"path/filepath"
	"testing"
)

func TestAuditLogger(t *testing.T) {
	s, e := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	l := New(s)
	if e = l.Record("u", "review", "r"); e != nil {
		t.Fatal(e)
	}
	if !ValidActor(domain.Audit{Actor: "u", RecordID: "r"}) {
		t.Fatal("actor")
	}
}
