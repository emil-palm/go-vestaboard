package layout

import "errors"

// Color represents constants for supported colors.
type Color int

var ErrInvalidColor = errors.New("invalid color")

const (
	PoppyRed Color = iota + 63
	Orange
	Yellow
	Green
	ParisBlue
	Violet
	White
)

func validateColor(c Color) error {
	if c < PoppyRed || c > White {
		return ErrInvalidColor
	}
	return nil
}
