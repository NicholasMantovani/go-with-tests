package poker

import (
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league sorted", func(t *testing.T) {

		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Nicholas", "Wins": 3},
			{"Name": "Goku", "Wins": 4},
			{"Name": "Tylor", "Wins": 5}]
			`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{Name: "Tylor", Wins: 5},
			{Name: "Goku", Wins: 4},
			{Name: "Nicholas", Wins: 3},
		}
		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Nicholas", "Wins": 3},
			{"Name": "Goku", "Wins": 4},
			{"Name": "Tylor", "Wins": 5}]
			`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		got := store.GetPlayerScore("Nicholas")

		want := 3

		assertScoreEquals(t, got, want)

	})

	t.Run("store win for existing player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Nicholas", "Wins": 3},
			{"Name": "Goku", "Wins": 4},
			{"Name": "Tylor", "Wins": 5}]
			`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Nicholas")

		got := store.GetPlayerScore("Nicholas")

		want := 4

		assertScoreEquals(t, got, want)
	})

	t.Run("store win for new player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Nicholas", "Wins": 3},
			{"Name": "Goku", "Wins": 4},
			{"Name": "Tylor", "Wins": 5}]
			`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)

		store.RecordWin("Nuovo")

		got := store.GetPlayerScore("Nuovo")

		want := 1

		assertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {

		database, clean := createTempFile(t, "")
		defer clean()

		_, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
	})
}
