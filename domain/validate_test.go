package domain

import "testing"

func TestValidation(t *testing.T) {
	if ValidateRecord(Record{}) == nil {
		t.Fatal("expected error")
	}
	if !TransitionAllowed("processing", "processed") {
		t.Fatal("transition")
	}
}
