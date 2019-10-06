package maps

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Search(word string) (string, error) {
	if def, ok := d[word]; ok {
		return def, nil
	}

	return "", ErrNotFound
}

// This is neat. I really like this.
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
