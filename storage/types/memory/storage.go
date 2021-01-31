package memory

import (
	types "github.com/jan25/sandbox-2021/storage/types"
)

// Storage represents a in-memory storage.
type Storage struct {
	kv map[int64]string
}

// NewStorage creates a in-memory storage.
func NewStorage() *Storage {
	return &Storage{kv: make(map[int64]string)}
}

func (s *Storage) Read(id int64) (types.Record, error) {
	if value, ok := s.kv[id]; ok {
		return Record{id: id, value: value}, nil
	}
	return nil, types.ErrRecordNotFound
}

func (s *Storage) Write(r types.Record) error {
	s.kv[r.GetID()] = r.GetValue()
	return nil
}
