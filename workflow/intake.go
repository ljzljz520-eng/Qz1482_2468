package workflow

import (
	"aerialfarm/domain"
	"aerialfarm/storage"
	"context"
	"fmt"
	"time"
)

type Service struct {
	Store *storage.Store
	clock func() time.Time
}

func NewService(s *storage.Store) *Service { return &Service{Store: s, clock: time.Now} }
func (s *Service) Register(ctx context.Context, id, field, pilot string) (domain.Record, error) {
	select {
	case <-ctx.Done():
		return domain.Record{}, ctx.Err()
	default:
	}
	r := domain.NewRecord(id, field, pilot, s.clock())
	if e := domain.ValidateRecord(r); e != nil {
		return r, e
	}
	if e := s.Store.PutRecord(r); e != nil {
		return r, e
	}
	return r, nil
}
func (s *Service) Begin(ctx context.Context, id string) (domain.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if !domain.TransitionAllowed(r.Status, "processing") {
		return r, fmt.Errorf("invalid transition")
	}
	r.Status = "processing"
	e = s.Store.PutRecord(r)
	return r, e
}
func (s *Service) Archive(id string) (domain.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if !domain.TransitionAllowed(r.Status, "archived") {
		return r, fmt.Errorf("cannot archive")
	}
	r.Status = "archived"
	e = s.Store.PutRecord(r)
	return r, e
}
