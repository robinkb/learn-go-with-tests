package poker

import (
	"os"
	"testing"
)

var testDatabase = `
	[
		{"name": "Cleo", "wins": 10},
		{"name": "Chris", "wins": 33}
	]
`

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, testDatabase)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		AssertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, testDatabase)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		AssertScoreEquals(t, store.GetPlayerScore("Chris"), 33)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, testDatabase)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, testDatabase)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		store.RecordWin("Mike")

		got := store.GetPlayerScore("Mike")
		want := 1

		AssertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, testDatabase)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		got := store.GetLeague()
		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		AssertLeague(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
