package maps

import "errors"

var ErrNotFound = errors.New("EXPECTED TO GET AN ERROR")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// An interesting property of maps is
// that you can modify them without passing as an address to it
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
