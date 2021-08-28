// Package font provides supports for parsing
// several font formats (postscript, bitmap and truetype)
// and provides a common API.
//
// It provides low-level metric information used by shaping
// engines and a way of rendering outline and bitmap glyphs.
package font

// Face is the common interface to represent various font formats.
//
// A Face object is the in-memory representation of a font file:
// as such it does not hold scaling information.
//
// Complex font features may require an extension interface :
// see opentype.FaceOpentype for an example.
type Face interface {
	FaceMetrics
	FaceRenderer
}

// FaceMetrics exposes the information about a font file
// required by a layout engine.
// All the returned metrics are expressed in font units
// and should be scaled.
type FaceMetrics interface {
	// Upem returns the units per em of the font file.
	// If not found, should return 1000 as fallback value.
	Upem() uint16

	// GlyphName returns the name of the given glyph, or an empty
	// string if the glyph is invalid or has no name.
	GlyphName(glyph GID) string

	// FontHExtents returns the extents of the font for horizontal text in font units,
	// or false if not available.
	// `varCoords` (in normalized coordinates) is only useful for variable fonts.
	FontHExtents(varCoords []float32) (FontExtents, bool)

	// FontVExtents is the same as `FontHExtents`, but for vertical text.
	FontVExtents(varCoords []float32) (FontExtents, bool)

	// NominalGlyph returns the glyph used to represent the given rune,
	// or false if not found.
	NominalGlyph(ch rune) (GID, bool)

	// HorizontalAdvance returns the horizontal advance in font units.
	// When no data is available but the glyph index is valid, this method
	// should return a default value (the upem number for example).
	// If the glyph is invalid it should return 0.
	// `varCoords` is used by variable fonts, and is specified in normalized coordinates.
	HorizontalAdvance(glyph GID, varCoords []float32) float32

	// VerticalAdvance is the same as `HorizontalAdvance`, but for vertical advance.
	VerticalAdvance(glyph GID, varCoords []float32) float32

	// GlyphHOrigin fetches the (X,Y) coordinates of the origin (in font units) for a glyph ID,
	// for horizontal text segments.
	// Returns `false` if not available.
	GlyphHOrigin(glyph GID, varCoords []float32) (x, y int32, found bool)

	// GlyphVOrigin is the same as `GlyphHOrigin`, but for vertical text segments.
	GlyphVOrigin(glyph GID, varCoords []float32) (x, y int32, found bool)

	// GlyphExtents retrieve the extents for a specified glyph, of false, if not available.
	// `varCoords` is used by variable fonts, and is specified in normalized coordinates.
	// For bitmap glyphs, the closest resolution to `xPpem` and `yPpem` is selected.
	GlyphExtents(glyph GID, varCoords []float32, xPpem, yPpem uint16) (GlyphExtents, bool)

	// NormalizeVariations should normalize the given design-space coordinates. The minimum and maximum
	// values for the axis are mapped to the interval [-1,1], with the default
	// axis value mapped to 0.
	// This should be a no-op for non-variable fonts.
	NormalizeVariations(varCoords []float32) []float32
}

// FaceRenderer implements the drawing of glyphs.
type FaceRenderer interface {
	// TODO: add position information and output/destination
	Draw(glyph GID)
}

// GID is used to identify glyphs in a font.
// It is mostly internal to the font and should not be confused with
// Unicode code points.
type GID uint32

// FontExtents exposes font-wide extent values, measured in font units.
// Note that typically ascender is positive and descender negative in coordinate systems that grow up.
type FontExtents struct {
	Ascender  float32 // Typographic ascender.
	Descender float32 // Typographic descender.
	LineGap   float32 // Suggested line spacing gap.
}

// GlyphExtents exposes extent values, measured in font units.
// Note that height is negative in coordinate systems that grow up.
type GlyphExtents struct {
	XBearing float32 // Left side of glyph from origin
	YBearing float32 // Top side of glyph from origin
	Width    float32 // Distance from left to right side
	Height   float32 // Distance from top to bottom side
}
