package mapps

import (
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

		if got == nil {
			t.Fatal("Expected to get an error")
		}

		assertError(t, got, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

func assertError( t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q ", got, want)
	}
}




func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		defination := "this is just a test"

		err := dictionary.Add(word, defination)

		assertError(t, err, nil)
		assertDefination(t, dictionary, word, defination)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictionary := Dictionary{word: defination}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefination(t, dictionary, word, defination)

	})

}


func assertDefination(t testing.TB, dictionary Dictionary, word, defination string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, defination)
}



func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictionary := Dictionary{word: defination}
		newDefination := "new defination"

		err := dictionary.Update(word, newDefination)

		assertError(t, err, nil)
		assertDefination(t, dictionary, word, newDefination)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, defination)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test defination"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted ", word)
	}
}