package repository_test

import (
	"depot/pkg/repository"
	"testing"
)

func Test_FlushStore(t *testing.T) {

	t.Run(t.Name(), func(t *testing.T) {
		inMemStore := &repository.InMemoryStore{}

		inMemStore.FlushStore()
		storeData := inMemStore.GetAllStoreData()
		if len(storeData) != 0 {
			t.Fatalf("After setValue call, store count not correct got %d and want %d", len(storeData), 0)
		}
	})
}
func Test_AddItemFromStore_OneItem(t *testing.T) {
	t.Run(t.Name(), func(t *testing.T) {

		want := map[string]string{"Item1": "Item1Val"}
		inMemStore := repository.NewInMemoryStore(make(map[string]string))

		inMemStore.AddItemToStore("Item1", "Item1Val")
		got := inMemStore.GetAllStoreData()

		assertStoreDataCount(t, want, got)
	})
}
func Test_AddItemFromStore_More(t *testing.T) {
	t.Run(t.Name(), func(t *testing.T) {

		inMemStore := repository.NewInMemoryStore(make(map[string]string))
		inMemStore.AddItemToStore("Item1", "Item1Val")
		inMemStore.AddItemToStore("Item2", "Item1Val")
		inMemStore.AddItemToStore("Item3", "Item1Val")

		storeData := inMemStore.GetAllStoreData()
		want := 3

		if len(storeData) != want {
			t.Errorf("Not correct AddItemToStore count %d, should be %d", len(storeData), want)
		}
	})
}
func Test_GetItemFromStore(t *testing.T) {
	t.Run(t.Name(), func(t *testing.T) {

		cases := []struct {
			name    string
			want    string
			keyWord string
		}{
			{
				name:    "GetItemFromStore case #1",
				want:    "Belgium",
				keyWord: "Brussels",
			},
			{
				name:    "GetItemFromStore case #2",
				want:    "France",
				keyWord: "Paris",
			}, {
				name:    "GetItemFromStore case #3",
				want:    "Turkey",
				keyWord: "",
			},
		}

		initData := map[string]string{
			"Amsterdam": "Netherlands",
			"Brussels":  "Belgium",
			"Paris":     "France",
			"Madrid":    "Spain",
		}

		inMemStore := repository.NewInMemoryStore(initData)

		for _, cs := range cases {
			t.Run(cs.name, func(t *testing.T) {
				got, err := inMemStore.GetItemFromStore(cs.keyWord)
				if got != cs.want {
					t.Errorf("Not correct get item from store, got %q, want %q", got, cs.want)
				}

				if err != nil {
					t.Errorf("Not expected error case, got %s", err)
				}

			})
		}

	})
}

func assertStoreDataCount(t testing.TB, want, got map[string]string) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("Not correct response body, got %q, want %q", got, want)
	}
}
