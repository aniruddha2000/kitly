package urlshortner

import "testing"

func TestStore(t *testing.T) {
	t.Run("Store", func(t *testing.T) {
		store := NewStore()
		store.Short("github.com")

		if _, ok := store.Data["github.com"]; !ok {
			t.Errorf("func -> %v | Data not being inserted", t.Name())
		}
	})
}
