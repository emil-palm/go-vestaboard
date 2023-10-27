package layout

import (
	"errors"
	"fmt"
)

// PrintableChars is a string of all of the Vestaboard accepted chrs
const PrintableChars = " ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$() - +&=;: '\"%,.  /? °"

var (
	ErrInvalidCharacter = errors.New("invalid character")

	charNumbers map[string]int
)

func init() {
	charNumbers = make(map[string]int)
	for i, c := range PrintableChars {
		if _, ok := charNumbers[string(c)]; ok {
			// skip spaces.
			continue
		}
		charNumbers[string(c)] = i
	}
}

func CharToCode(c string) (int, error) {
	i, ok := charNumbers[c]
	if !ok {
		return -1, ErrInvalidCharacter
	}

	return i, nil
}

func ValidText(t string, newlineAccepted bool) error {
	for i, c := range t {
		if newlineAccepted && c == '\n' {
			continue
		}

		if _, err := CharToCode(string(c)); err != nil {
			return fmt.Errorf("invalid character %q at position %d, %w", string(c), i, err)
		}
	}
	return nil
}
