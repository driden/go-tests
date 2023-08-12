package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound      = DictionaryErr("could not find the word you are looking for")
	ErrWordExists    = DictionaryErr("word already exists")
	ErrWordNotExists = DictionaryErr("word doesn't exist")
)

func (d Dictionary) Search(word string) (error, string) {
	meaning, present := d[word]

	if !present {
		return ErrNotFound, ""
	}

	return nil, meaning
}

func (d Dictionary) Add(word, definition string) error {
	err, _ := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, definition string) error {
	err, _ := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordNotExists
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (err DictionaryErr) Error() string {
	return string(err)
}
