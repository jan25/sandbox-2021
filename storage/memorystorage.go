package main

// MemoryRecord represensts a record for in-memory storage.
type MemoryRecord struct {
	id    int64
	value string
}

// NewMemoryRecord created new record for in-memory storage.
func NewMemoryRecord(id int64, value string) MemoryRecord {
	return MemoryRecord{id: id, value: value}
}

func (m MemoryRecord) getID() int64 {
	return m.id
}

func (m MemoryRecord) getValue() string {
	return m.value
}

// MemoryStorage represents a in-memory storage.
type MemoryStorage struct {
	kv map[int64]string
}

// NewMemoryStorage creates a memory storage.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{kv: make(map[int64]string)}
}

func (s *MemoryStorage) read(id int64) (Record, error) {
	if value, ok := s.kv[id]; ok {
		return MemoryRecord{id: id, value: value}, nil
	}
	return nil, ErrRecordNotFound
}

func (s *MemoryStorage) write(r Record) error {
	s.kv[r.getID()] = r.getValue()
	return nil
}
