package maps

import "errors"

type Dictionary map[string]string

var ErrKeyNotPresent = errors.New("key not present in the dictionary")

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if ok {
		return val, nil
	}
	return "", ErrKeyNotPresent
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
