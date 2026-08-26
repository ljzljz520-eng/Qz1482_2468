package workflow

import (
	"aerialfarm/domain"
	"context"
	"fmt"
)

type Queue struct{ Pending []string }

func NewQueue() *Queue { return &Queue{Pending: []string{}} }
func (q *Queue) Enqueue(id string) error {
	if id == "" {
		return fmt.Errorf("empty id")
	}
	q.Pending = append(q.Pending, id)
	return nil
}
func (q *Queue) Dequeue() (string, error) {
	if len(q.Pending) == 0 {
		return "", fmt.Errorf("queue empty")
	}
	id := q.Pending[0]
	q.Pending = q.Pending[1:]
	return id, nil
}
func (s *Service) ProcessQueued(ctx context.Context, q *Queue) ([]domain.Record, error) {
	out := []domain.Record{}
	for len(q.Pending) > 0 {
		id, _ := q.Dequeue()
		r, e := s.Process(ctx, id)
		if e != nil {
			return out, e
		}
		out = append(out, r)
	}
	return out, nil
}
