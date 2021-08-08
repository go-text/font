// Package opentype provides an implementation of
// font.Face for Opentype (Truetype) fonts.
//
// It supports advanced Opentype and AAT layout features.
package opentype

import "github.com/go-text/font"

// var _ FaceOpentype = (*Face)(nil)

// FaceOpentype is a specialization of font.Face for Opentype
// font files.
type FaceOpentype interface {
	font.Face

	// TablesLayout fetches the advanced Opentype and AAT layout tables of the font.
	TablesLayout() TablesLayout

	// Variations returns the variations for the font,
	// or an empty table.
	Variations() TableFvar

	// IsGraphite returns true if the font has Graphite capabilities.
	// The returned Face will be used to load Graphite tables.
	// Overide this method to disable Graphite functionalities.
	IsGraphite() (*Face, bool)
}

// Face is the in-memory representation of a font file (.ttf, .otf)
// or an element of a font collection (.ttc, .otc, .dfont)
type Face struct{}

// TableFvar is the font variations table
// (https://docs.microsoft.com/typography/opentype/spec/fvar)
type TableFvar struct{}

// TablesLayout exposes advanced layout tables.
// All the fields are optionals, since a font may only provide a subset of these tables.
type TablesLayout struct {
	GDEF TableGDEF // An absent table has a nil Class
	Trak TableTrak
	Ankr TableAnkr
	Feat TableFeat
	Morx TableMorx
	Kern TableKernx
	Kerx TableKernx
	GSUB TableGSUB // An absent table has a nil slice of lookups
	GPOS TableGPOS // An absent table has a nil slice of lookups
}

type (
	// https://docs.microsoft.com/typography/opentype/spec/gdef
	TableGDEF struct{}
	// https://developer.apple.com/fonts/TrueType-Reference-Manual/RM06/Chap6trak.html
	TableTrak struct{}
	// https://developer.apple.com/fonts/TrueType-Reference-Manual/RM06/Chap6ankr.html
	TableAnkr struct{}
	// https://developer.apple.com/fonts/TrueType-Reference-Manual/RM06/Chap6feat.html
	TableFeat struct{}
	// https://developer.apple.com/fonts/TrueType-Reference-Manual/RM06/Chap6morx.html
	TableMorx struct{}
	// https://developer.apple.com/fonts/TrueType-Reference-Manual/RM06/Chap6kerx.html
	TableKernx struct{}
	// https://docs.microsoft.com/typography/opentype/spec/gsub
	TableGSUB struct{}
	// https://docs.microsoft.com/typography/opentype/spec/gpos
	TableGPOS struct{}
)
