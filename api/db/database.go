package db

type DB struct {
	// JSON database
	MemoryDB *MemoryDB
}

func NewWithMemoryDB(path string) (*DB, error) {
	memorydb, err := Open(path)
	if err != nil {
		return nil, err
	}
	return &DB{memorydb}, nil
}
