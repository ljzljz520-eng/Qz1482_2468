package query

import "aerialfarm/domain"

type Page struct {
	Items                []domain.Record
	Offset, Limit, Total int
}

func Paginate(rs []domain.Record, offset, limit int) Page {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 20
	}
	if offset > len(rs) {
		offset = len(rs)
	}
	end := offset + limit
	if end > len(rs) {
		end = len(rs)
	}
	return Page{Items: rs[offset:end], Offset: offset, Limit: limit, Total: len(rs)}
}
func SortByStatus(rs []domain.Record) []domain.Record {
	out := append([]domain.Record{}, rs...)
	for i := 0; i < len(out); i++ {
		for j := i + 1; j < len(out); j++ {
			if out[j].Status < out[i].Status {
				out[i], out[j] = out[j], out[i]
			}
		}
	}
	return out
}
