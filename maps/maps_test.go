package maps

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new wrod", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("new wrod", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}

func TestDictionaryErr(t *testing.T) {
	t.Run("error with method Error", func(t *testing.T) {
		dictionaryErr := DictionaryErr("some error")

		if dictionaryErr.Error() != "some error" {
			t.Errorf("got %s want %s", dictionaryErr.Error(), "some error")
		}
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should finc added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func ExampleDictionary() {
	dictionary := Dictionary{"test": "test definition"}

	fmt.Print(dictionary)
	// Output: map[test:test definition]
}

func ExampleDictionary_Search() {
	dictionary := Dictionary{"test": "test definition"}
	definition, err := dictionary.Search("test")

	if err != nil {
		fmt.Print("word not found")
	}

	fmt.Print(definition)
	// Output: test definition
}

func ExampleDictionary_Add() {
	dictionary := Dictionary{}
	err := dictionary.Add("test", "test definition")

	if err != nil {
		fmt.Print("word already exists")
	}

	definition, _ := dictionary.Search("test")

	fmt.Print(definition)
	// Output: test definition
}

func ExampleDictionary_Update() {
	dictionary := Dictionary{"test": "test definition"}

	err := dictionary.Update("test", "new definition")

	if err != nil {
		fmt.Print("word does not exist")
	}

	definition, _ := dictionary.Search("test")
	fmt.Print(definition)
	// Output: new definition
}

func ExampleDictionary_Delete() {
	dictionary := Dictionary{"test": "test definition"}
	dictionary.Delete("test")

	// Throw an error if the word does not exist
	definition, _ := dictionary.Search("test")

	fmt.Print(definition)
	// Output:
}
