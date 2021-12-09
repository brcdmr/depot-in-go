package repository

// InMemoryStore is a thread safe store
type InMemoryStore struct {
	store chan func(map[string]string)
}

func NewInMemoryStore(initialStoreData map[string]string) *InMemoryStore {
	s := &InMemoryStore{
		store: make(chan func(map[string]string)),
	}

	go s.loopItems(initialStoreData)
	return s
}

func (st *InMemoryStore) loopItems(initialStoreData map[string]string) {
	items := initialStoreData
	for op := range st.store {
		op(items)
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
