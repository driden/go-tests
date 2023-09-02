package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		_, got := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		got, _ := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error.")
		}

		assertErr(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		assertDefinition(t, dictionary, "test", "this is just a test")
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add("test", "this is just a test2")

		assertErr(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update an existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertErr(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertErr(t, err, ErrWordNotExists)
	})
}
func TestDelete(t *testing.T) {
	t.Run("delete an existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		err, _ := dictionary.Search(word)

		if err != ErrNotFound {
			t.Errorf("Expected %q to be the deleted", word)
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertErr(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, want string) {
	t.Helper()

	err, got := dictionary.Search(word)

	if err != nil {
		t.Fatalf("Should've found word %q: %q", word, err)
	}

	assertStrings(t, got, want)
}
