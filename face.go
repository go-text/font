package font

import "github.com/benoitkugler/textlayout/fonts"

type Face interface {
	// NominalGlyph returns the glyph used to represent the given rune,
	// or false the rune is not supported.
	NominalGlyph(r rune) (fonts.GID, bool)
}
