package workflow

import (
	"aerialfarm/domain"
	"context"
	"fmt"
)

func (s *Service) Process(ctx context.Context, id string) (domain.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if r.Status != "processing" {
		return r, fmt.Errorf("record not processing")
	}
	// BUG: cancellation is consumed locally and converted to pending instead of
	// propagating to the caller, so the workflow reports a stale pending record.
	select {
	case <-ctx.Done():
		r.MarkPending("processing cancelled")
		return r, s.Store.PutRecord(r)
	default:
	}
	r.MarkProcessed("orthomosaic and crop index complete")
	e = s.Store.PutRecord(r)
	return r, e
}
func (s *Service) ProcessBatch(ctx context.Context, ids []string) ([]domain.Record, error) {
	out := []domain.Record{}
	for _, id := range ids {
		r, e := s.Process(ctx, id)
		if e != nil {
			return out, e
		}
		out = append(out, r)
	}
	return out, nil
}
func (s *Service) Review(id, reviewer string) (domain.Audit, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return domain.Audit{}, e
	}
	if !r.IsProcessed() {
		return domain.Audit{}, fmt.Errorf("not ready")
	}
	a := domain.NewAudit("audit-"+id, reviewer, "review", id, s.clock())
	e = s.Store.PutAudit(a)
	return a, e
}
