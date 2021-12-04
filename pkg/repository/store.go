package repository

import (
	"errors"
	"sync"
)

type InMemoryStore struct {
	store map[string]string

	//A mutex is used to sync read/write access to the map
	lock sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		map[string]string{},
		sync.RWMutex{},
	}
}

func (st *InMemoryStore) AddItemToStore(key string, value string) {
	st.lock.Lock()
	defer st.lock.Unlock()
	st.store[key] = value
}

func (st *InMemoryStore) GetItemFromStore(key string) (string, error) {
	st.lock.RLock()
	defer st.lock.RUnlock()

	data := st.store[key]

	if data == "" {
		return data, errors.New(key + " does not found in the storage")
	}
	return data, nil
}
