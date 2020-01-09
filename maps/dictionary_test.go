package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known words", func(t *testing.T) {
		e := "this is just a test"
		assertDefinition(t, dictionary, "test", e)
	})

	t.Run("unknown words", func(t *testing.T) {
		_, err := dictionary.Search("whatever")

		assertError(t, ErrNotFound, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("a new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		e := "this is just a test"
		assertDefinition(t, dictionary, "test", e)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		err := dictionary.Add("test", "something else")

		assertError(t, ErrWordExists, err)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: definition}

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
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, expected string) {
	t.Helper()
	a, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find word:", word)
	}

	if a != expected {
		t.Errorf("expected: %q actual: %q", expected, a)
	}
}

func assertError(t *testing.T, e, a error) {
	t.Helper()
	if a != e {
		t.Errorf("Expected error: %q actual: %q", e, a)
	}
}
