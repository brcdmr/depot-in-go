package repository

import (
	"fmt"
	"log"
	"time"
)

// InMemoryStore is a thread safe store
type InMemoryStore struct {
	store     chan func(map[string]string)
	intervals chan func(map[string]*time.Timer)
	fileSys   FileSystem
	interval  time.Duration
	fileName  string
}

func NewInMemoryStore(initialStoreData map[string]string, interval time.Duration, fPath string) *InMemoryStore {
	s := &InMemoryStore{
		store:     make(chan func(map[string]string)),
		intervals: make(chan func(map[string]*time.Timer)),
		fileSys:   FileSystem{Path: fPath},
		interval:  interval,
		fileName:  "",
	}

	go s.loopItems(initialStoreData)
	go s.loopIntervals()
	go s.startFileScheduler(interval)
	return s
}

func (st *InMemoryStore) loopItems(initialStoreData map[string]string) {
	items := initialStoreData
	for act := range st.store {
		act(items)
	}
}

func (st *InMemoryStore) loopIntervals() {
	inters := map[string]*time.Timer{}
	for act := range st.intervals {
		act(inters)
	}
}

func (st *InMemoryStore) AddItemToStore(key string, value string) {

	st.store <- func(items map[string]string) {
		items[key] = value
	}
}

// Get retrieves a store data at the specified key
func (st *InMemoryStore) GetItemFromStore(key string) string {

	data := make(chan string, 1)
	st.store <- func(items map[string]string) {
		data <- items[key]
	}

	return <-data
}

// Delete all data from store
func (st *InMemoryStore) FlushStore() {

	st.store <- func(items map[string]string) {
		for key := range items {
			delete(items, key)
		}
	}
}

func (st *InMemoryStore) GetAllStoreData() map[string]string {
	allData := make(chan map[string]string, 1)
	st.store <- func(items map[string]string) {
		p := map[string]string{}
		for key, value := range items {
			p[key] = value
		}

		allData <- p
	}

	return <-allData
}

func (st *InMemoryStore) startFileScheduler(duration time.Duration) {

	for range time.Tick(duration * time.Second) { // to do: change to min
		timeStamp := time.Now().Unix()

		newFileName := fmt.Sprintf("%d-data.json", timeStamp)

		log.Printf("%s - %s minutes have passed and called write file function!!", newFileName, duration.String())
		st.fileSys.WriteFile(st.GetAllStoreData(), newFileName)

		st.fileSys.RemoveFile(st.fileName)
		st.fileName = newFileName

	}

}
