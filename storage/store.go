package storage

import (
	"aerialfarm/domain"
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"sync"
)

var recordsBucket = []byte("records")
var usersBucket = []byte("users")
var eventsBucket = []byte("events")
var auditsBucket = []byte("audits")

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range [][]byte{recordsBucket, usersBucket, eventsBucket, auditsBucket} {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func put[T any](db *bbolt.DB, b []byte, key string, v T) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(b).Put([]byte(key), data) })
}
func get[T any](db *bbolt.DB, b []byte, key string, out *T) error {
	return db.View(func(tx *bbolt.Tx) error {
		raw := tx.Bucket(b).Get([]byte(key))
		if raw == nil {
			return fmt.Errorf("not found: %s", key)
		}
		return json.Unmarshal(raw, out)
	})
}
func (s *Store) PutRecord(r domain.Record) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return put(s.db, recordsBucket, r.ID, r)
}
func (s *Store) GetRecord(id string) (domain.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var r domain.Record
	e := get(s.db, recordsBucket, id, &r)
	return r, e
}
func (s *Store) PutUser(u domain.User) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return put(s.db, usersBucket, u.ID, u)
}
func (s *Store) GetUser(id string) (domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var u domain.User
	e := get(s.db, usersBucket, id, &u)
	return u, e
}
func (s *Store) PutEvent(v domain.Event) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return put(s.db, eventsBucket, v.ID, v)
}
func (s *Store) PutAudit(v domain.Audit) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return put(s.db, auditsBucket, v.ID, v)
}
func (s *Store) ListRecords() ([]domain.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var out []domain.Record
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(recordsBucket).ForEach(func(_, v []byte) error {
			var r domain.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
