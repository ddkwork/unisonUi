// Copyright ©2021-2022 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package unison

import (
	"embed"
	"path"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/richardwilkes/toolbox"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/unison/internal/skia"
)

// DefaultSystemFamilyName is the default system font family name and will be used as a fallback where needed.
const DefaultSystemFamilyName = "Roboto"

//go:embed resources/fonts
var fontFS embed.FS

var (
	internalFontLock sync.RWMutex
	internalFonts         = make(map[string]*internalFont)
	_                Font = &fontImpl{}
)

// Pre-defined fonts
var (
	SystemFont                = &IndirectFont{}
	EmphasizedSystemFont      = &IndirectFont{}
	SmallSystemFont           = &IndirectFont{}
	EmphasizedSmallSystemFont = &IndirectFont{}
	LabelFont                 = &IndirectFont{}
	FieldFont                 = &IndirectFont{}
	KeyboardFont              = &IndirectFont{}
)

// Font holds a realized FontFace of a specific size that can be used to render text.
type Font interface {
	// Face returns the FontFace this Font belongs to.
	Face() *FontFace
	// Size returns the size of the font. This is the value that was passed to FontFace.Font() when creating the font.
	Size() float32
	// Metrics returns a copy of the FontMetrics for this font.
	Metrics() FontMetrics
	// Baseline returns the number of logical pixels to the bottom of characters without descenders.
	Baseline() float32
	// LineHeight returns the recommended line height of the font.
	LineHeight() float32
	// RuneToGlyph converts a rune into a glyph. Missing glyphs will have a value of 0.
	RuneToGlyph(r rune) uint16
	// RunesToGlyphs converts the runes into glyphs. Missing glyphs will have a value of 0.
	RunesToGlyphs(r []rune) []uint16
	// GlyphWidth returns the width for the glyph. This does not do font fallback for missing glyphs.
	GlyphWidth(glyph uint16) float32
	// GlyphWidths returns the widths for each glyph. This does not do font fallback for missing glyphs.
	GlyphWidths(glyphs []uint16) []float32
	// SimpleWidth returns the width of a string. It does not do font fallback, nor does it consider tabs or line
	// endings.
	SimpleWidth(str string) float32
	// Descriptor returns a FontDescriptor for this Font.
	Descriptor() FontDescriptor
	skiaFont() skia.Font
}

type internalFont struct {
	family string
	faces  []*FontFace
}

// FontHinting holds the type of font hinting to use.
type FontHinting byte

// Possible values for FontHinting.
const (
	FontHintingNone FontHinting = iota
	FontHintingSlight
	FontHintingNormal
	FontHintingFull
)

// FontMetrics flags
const (
	UnderlineThicknessIsValidFontMetricsFlag = 1 << iota
	UnderlinePositionIsValidFontMetricsFlag
	StrikeoutThicknessIsValidFontMetricsFlag
	StrikeoutPositionIsValidFontMetricsFlag
	BoundsInvalidFontMetricsFlag
)

// FontMetrics holds various metrics about a font.
type FontMetrics = skia.FontMetrics

type fontImpl struct {
	size    float32
	face    *FontFace
	font    skia.Font
	metrics FontMetrics
}

func (f *fontImpl) Face() *FontFace {
	return f.face
}

func (f *fontImpl) Size() float32 {
	return f.size
}

func (f *fontImpl) Metrics() FontMetrics {
	return f.metrics
}

func (f *fontImpl) Baseline() float32 {
	return f.metrics.Descent + f.size
}

func (f *fontImpl) LineHeight() float32 {
	return f.size + f.metrics.Descent*2
}

func (f *fontImpl) RuneToGlyph(r rune) uint16 {
	return skia.FontRuneToGlyph(f.font, r)
}

func (f *fontImpl) RunesToGlyphs(r []rune) []uint16 {
	if len(r) == 0 {
		return nil
	}
	return skia.FontRunesToGlyphs(f.font, r)
}

func (f *fontImpl) GlyphWidth(glyph uint16) float32 {
	return skia.FontGlyphWidths(f.font, []uint16{glyph})[0]
}

func (f *fontImpl) GlyphWidths(glyphs []uint16) []float32 {
	if len(glyphs) == 0 {
		return nil
	}
	return skia.FontGlyphWidths(f.font, glyphs)
}

func (f *fontImpl) SimpleWidth(str string) float32 {
	if str == "" {
		return 0
	}
	return skia.FontMeasureText(f.font, str)
}

func (f *fontImpl) skiaFont() skia.Font {
	return f.font
}

func (f *fontImpl) Descriptor() FontDescriptor {
	weight, spacing, slant := f.face.Style()
	return FontDescriptor{
		Family:  f.face.Family(),
		Size:    f.size,
		Weight:  weight,
		Spacing: spacing,
		Slant:   slant,
	}
}

func init() {
	const fontDir = "resources/fonts"
	entries, err := fontFS.ReadDir(fontDir)
	if err != nil {
		jot.Error(errs.NewWithCause("unable to read embedded file system", err))
		return
	}
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			name := entry.Name()
			lower := strings.ToLower(name)
			if strings.HasSuffix(lower, ".otf") || strings.HasSuffix(lower, ".ttf") {
				var data []byte
				if data, err = fontFS.ReadFile(path.Join(fontDir, name)); err != nil {
					jot.Error(errs.NewWithCausef(err, "unable to read font %s", name))
				} else if _, err = RegisterFont(data); err != nil {
					jot.Error(errs.NewWithCause(name, err))
				}
			}
		}
	}

	baseSize := float32(10)
	SystemFont.Font = MatchFontFace(DefaultSystemFamilyName, MediumFontWeight, StandardSpacing, NoSlant).Font(baseSize)
	EmphasizedSystemFont.Font = MatchFontFace(DefaultSystemFamilyName, BoldFontWeight, StandardSpacing, NoSlant).Font(baseSize)
	SmallSystemFont.Font = MatchFontFace(DefaultSystemFamilyName, MediumFontWeight, StandardSpacing, NoSlant).Font(baseSize - 1)
	EmphasizedSmallSystemFont.Font = MatchFontFace(DefaultSystemFamilyName, BoldFontWeight, StandardSpacing, NoSlant).Font(baseSize - 1)
	LabelFont.Font = MatchFontFace(DefaultSystemFamilyName, NormalFontWeight, StandardSpacing, NoSlant).Font(baseSize)
	FieldFont.Font = MatchFontFace(DefaultSystemFamilyName, NormalFontWeight, StandardSpacing, NoSlant).Font(baseSize)
	keyboardFamilyName := DefaultSystemFamilyName
	if runtime.GOOS == toolbox.MacOS {
		// This is a special font on macOS. Ideally, I'd find a source for an equivalent font and embed it so that the
		// same font could be used on all platforms.
		keyboardFamilyName = ".Keyboard"
	}
	KeyboardFont.Font = MatchFontFace(keyboardFamilyName, MediumFontWeight, StandardSpacing, NoSlant).Font(baseSize)
}

// RegisterFont registers a font with the font manager.
func RegisterFont(data []byte) (*FontDescriptor, error) {
	f := CreateFontFace(data)
	if f == nil {
		return nil, errs.New("unable to load font")
	}
	weight, spacing, slant := f.Style()
	fd := &FontDescriptor{
		Family:  f.Family(),
		Size:    9, // Arbitrary
		Weight:  weight,
		Spacing: spacing,
		Slant:   slant,
	}
	internalFontLock.Lock()
	defer internalFontLock.Unlock()
	if info, ok := internalFonts[fd.Family]; ok {
		add := true
		for _, one := range info.faces {
			weight2, spacing2, slant2 := one.Style()
			if weight == weight2 && spacing == spacing2 && slant == slant2 {
				add = false
				break
			}
		}
		if add {
			info.faces = append(info.faces, f)
			sort.Slice(info.faces, func(i, j int) bool {
				return info.faces[i].Less(info.faces[j])
			})
		}
	} else {
		internalFonts[fd.Family] = &internalFont{
			family: fd.Family,
			faces:  []*FontFace{f},
		}
	}
	return fd, nil
}
