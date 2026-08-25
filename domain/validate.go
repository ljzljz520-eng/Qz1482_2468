package domain

import "fmt"

func ValidateRecord(r Record) error {
	if r.ID == "" {
		return fmt.Errorf("record id required")
	}
	if r.FieldID == "" {
		return fmt.Errorf("field id required")
	}
	if r.Pilot == "" {
		return fmt.Errorf("pilot required")
	}
	return nil
}
func ValidateUser(u User) error {
	if u.ID == "" || u.Name == "" {
		return fmt.Errorf("user identity required")
	}
	if u.Role == "" {
		return fmt.Errorf("role required")
	}
	return nil
}
func AllowedStatus(s string) bool {
	switch s {
	case "received", "processing", "processed", "pending", "archived":
		return true
	default:
		return false
	}
}
func TransitionAllowed(from, to string) bool {
	if !AllowedStatus(from) || !AllowedStatus(to) {
		return false
	}
	if from == "received" && (to == "processing" || to == "pending") {
		return true
	}
	if from == "processing" && (to == "processed" || to == "pending") {
		return true
	}
	return from == to || to == "archived"
}
