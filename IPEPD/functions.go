package main

import (
	"errors"
	"fmt"
)

// // 1
// func SafeDivide(a, b int, out *int) error {
// 	defer func() {
// 		fmt.Println("Done func")
// 	}()
// 	if out == nil {
// 		return errors.New("Pointer = nil")
// 	} else if b == 0 {
// 		return errors.New("It cannot be devide by 0")
// 	}
// 	*out = a / b
// 	return nil
// }

// 2
type Saver interface {
	Save(key, value string) error
	Get(key string) (string, error)
}

type MemoryStore struct {
	data map[string]string
}

func (m *MemoryStore) Save(key, value string) error {
	if m == nil {
		return errors.New("MemoryStore is nil")
	}
	if key == "" {
		return errors.New("Empty string")
	}
	m.data[key] = value
	return nil
}

func NewMempryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]string),
	}
}

func (m *MemoryStore) Get(key string) (string, error) {
	if m == nil {
		return "", errors.New("MemoryStore is nil")
	}
	if key == "" {
		return "", errors.New("Empty key")
	}
	if m.data == nil {
		return "", errors.New("store is not initialized")
	}

	v, ok := m.data[key]
	if !ok {
		return "", fmt.Errorf("Key %q not found", key)
	}
	return v, nil

}

func Upsert(s Saver, key, value string) error {
	if s == nil {
		return errors.New("Saver is nil")
	}
	return s.Save(key, value)

}
