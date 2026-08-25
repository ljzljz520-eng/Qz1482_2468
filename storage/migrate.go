package storage

import (
	"aerialfarm/domain"
	"go.etcd.io/bbolt"
)

func (s *Store) SeedUser(u domain.User) error {
	if err := domain.ValidateUser(u); err != nil {
		return err
	}
	return s.PutUser(u)
}
func (s *Store) CountRecords() (int, error) { rs, e := s.ListRecords(); return len(rs), e }
func (s *Store) RemoveRecord(id string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(recordsBucket).Delete([]byte(id)) })
}
