package maps

type Dictionary map[string]string

type DictionaryErr string

var (
	ErrNotFound         = DictionaryErr("No definition found.")
	ErrWordExists       = DictionaryErr("Word is already defined.")
	ErrWordDoesNotExist = DictionaryErr("Word does not exist.")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	def := d[word]

	if def == "" {
		return def, ErrNotFound
	}

	return def, nil
}

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

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
