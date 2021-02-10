package db

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

type MemoryDB struct {
	path string
	lock sync.Mutex
	file *os.File
}

func Open(path string) (*MemoryDB, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return &MemoryDB{file: f, path: path}, nil
}

// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func (m *MemoryDB) Load(path string, v interface{}) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}

// Save saves a representation of v to the file at path.
func (m *MemoryDB) Copy(v interface{}) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(m.file, r)
	return err
}

func (m *MemoryDB) WriteFile(v interface{}) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(m.path, b, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(m.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	m.file = f
	return nil

}

// Marshal is a function that marshals the object into an
// io.Reader.
// By default, it uses the JSON marshaller.
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
