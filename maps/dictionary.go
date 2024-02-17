package maps

// This one I very like it is quite useful to have "static" set of errors
// https://dave.cheney.net/2016/04/07/constant-errors
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

const (
	ErrKeyNotPresent    DictionaryErr = DictionaryErr("key not present in the dictionary")
	ErrKeyAlreadyExists DictionaryErr = DictionaryErr("key is already present in the dictionary")
	ErrCantUpdate       DictionaryErr = DictionaryErr("key not present, can't update")
	ErrNoDeleteKey      DictionaryErr = DictionaryErr("key not present, can't delete")
)

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if ok {
		return val, nil
	}
	return "", ErrKeyNotPresent
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	if err == nil {
		return ErrKeyAlreadyExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	// I don't really know why but I don't like switch in imperative languages
	// switch statements or pattern matching in Haskell or even Rust feels much better
	switch err {
	case ErrKeyNotPresent:
		return ErrCantUpdate

	case nil:
		d[key] = value

	default:
		return err
	}

	return nil
}

// This delete return an error, in tutorial it was said that it compicates API
// But i think that now our API is more flexible with error
// When user does not need to know if item was present just dont use error
func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	if err == ErrKeyNotPresent {
		return ErrNoDeleteKey
	}
	delete(d, key)
	return nil
}
