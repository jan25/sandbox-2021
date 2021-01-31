package memory

// Record represensts a record for in-memory storage.
type Record struct {
	id    int64
	value string
}

// NewRecord created new record for in-memory storage.
func NewRecord(id int64, value string) Record {
	return Record{id: id, value: value}
}

// GetID gets ID for record.
func (m Record) GetID() int64 {
	return m.id
}

// GetValue gets value for record.
func (m Record) GetValue() string {
	return m.value
}
