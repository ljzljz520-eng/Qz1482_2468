package api

import (
	"aerialfarm/domain"
	"aerialfarm/query"
	"aerialfarm/workflow"
	"context"
	"encoding/json"
)

type Server struct {
	Workflow *workflow.Service
	Query    *query.Service
}

func New(w *workflow.Service, q *query.Service) *Server { return &Server{Workflow: w, Query: q} }
func (s *Server) Register(ctx context.Context, id, field, pilot string) ([]byte, error) {
	r, e := s.Workflow.Register(ctx, id, field, pilot)
	if e != nil {
		return nil, e
	}
	return json.Marshal(r)
}
func (s *Server) Process(ctx context.Context, id string) ([]byte, error) {
	r, e := s.Workflow.Begin(ctx, id)
	if e != nil {
		return nil, e
	}
	r, e = s.Workflow.Process(ctx, r.ID)
	if e != nil {
		return nil, e
	}
	return json.Marshal(r)
}
func (s *Server) Lookup(field, status string) ([]domain.Record, error) {
	return s.Query.Find(field, status)
}
