package query

import (
	"aerialfarm/domain"
	"aerialfarm/report"
	"aerialfarm/storage"
)

type Service struct{ Store *storage.Store }

func New(s *storage.Store) *Service { return &Service{Store: s} }
func (q *Service) Find(field, status string) ([]domain.Record, error) {
	rs, e := q.Store.ListRecords()
	if e != nil {
		return nil, e
	}
	return storage.FilterRecords(rs, field, status), nil
}
func (q *Service) Dashboard(field string) (report.Summary, error) {
	rs, e := q.Find(field, "")
	if e != nil {
		return report.Summary{}, e
	}
	return report.Summarize(rs), nil
}
func (q *Service) Search(term string) ([]domain.Record, error) {
	rs, e := q.Store.ListRecords()
	if e != nil {
		return nil, e
	}
	return storage.SearchNotes(rs, term), nil
}
