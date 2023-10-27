package layout

import (
	"errors"
	"strings"
)

var (
	ErrMessageTruncated  = errors.New("message truncated")
	ErrInvalidCoordinate = errors.New("invalid coordinate")
)

type Layout [6][22]int

func NewLayout() Layout {
	var layout Layout
	return layout
}

func (l *Layout) ValidCoordinate(x, y int) error {
	if x < 0 || y < 0 || x > 5 || y > 21 {
		return ErrInvalidCoordinate
	}
	return nil
}

func (l *Layout) Print(sx, sy int, s string) error {
	if err := l.ValidCoordinate(sx, sy); err != nil {
		return err
	}

	s = strings.ToUpper(s)
	// Preflight the string before an invalid set.
	if err := ValidText(s, false); err != nil {
		return err
	}

	x, y := sx, sy
	for _, c := range s {
		if x > 5 {
			return ErrMessageTruncated
		}
		l[x][y], _ = CharToCode(string(c))
		y++
		if y == 22 {
			x++
			y = 0
		}
	}

	return nil
}

func (l *Layout) SetColor(x, y int, c Color) error {
	if err := l.ValidCoordinate(x, y); err != nil {
		return err
	}
	if err := validateColor(c); err != nil {
		return err
	}

	l[x][y] = int(c)
	return nil
}
