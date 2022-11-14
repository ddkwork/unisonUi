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
	"runtime"

	"github.com/richardwilkes/unison/internal/skia"
)

// Ink holds a color, pattern, or gradient to draw with.
type Ink interface {
	Paint(canvas *Canvas, rect Rect, style PaintStyle) *Paint
}

// PaintStyle holds the type of painting to do.
type PaintStyle byte

// Possible values for PaintStyle.
const (
	Fill PaintStyle = iota
	Stroke
	StrokeAndFill
)

// StrokeCap holds the style for rendering the endpoint of a stroked line.
type StrokeCap byte

// Possible values for StrokeCap.
const (
	ButtCap StrokeCap = iota
	RoundCap
	SquareCap
)

// StrokeJoin holds the method for drawing the junction between connected line segments.
type StrokeJoin byte

// Possible values for StrokeJoin.
const (
	MiterJoin StrokeJoin = iota
	RoundJoin
	BevelJoin
)

// BlendMode holds the mode used for blending pixels.
type BlendMode byte

// Possible values for BlendMode.
const (
	ClearBlendMode BlendMode = iota
	SrcBlendMode
	DstBlendMode
	SrcOverBlendMode
	DstOverBlendMode
	SrcInBlendMode
	DstInBlendMode
	SrcOutBlendMode
	DstOutBlendMode
	SrcAtopBlendMode
	DstAtopBlendMode
	XorBlendMode
	PlusBlendMode
	ModulateBlendMode
	ScreenBlendMode
	OverlayBlendMode
	DarkenBlendMode
	LightenBlendMode
	ColorDodgeBlendMode
	ColorBurnBlendMode
	HardLightBlendMode
	SoftLightBlendMode
	DifferenceBlendMode
	ExclusionBlendMode
	MultiplyBlendMode
	HueBlendMode
	SaturationBlendMode
	ColorBlendMode
	LuminosityBlendMode
)

// FilterQuality holds the image filtering level. Lower settings draw faster, while higher settings look better when the
// image is scaled.
type FilterQuality byte

// Possible values for FilterQuality.
const (
	NoQuality FilterQuality = iota
	LowQuality
	MediumQuality
	HighQuality
)

// Paint controls options applied when drawing.
type Paint struct {
	paint skia.Paint
}

func newPaint(paint skia.Paint) *Paint {
	p := &Paint{paint: paint}
	runtime.SetFinalizer(p, func(obj *Paint) {
		ReleaseOnUIThread(func() {
			skia.PaintDelete(obj.paint)
		})
	})
	return p
}

// NewPaint creates a new Paint.
func NewPaint() *Paint {
	p := newPaint(skia.PaintNew())
	p.SetAntialias(true)
	return p
}

func (p *Paint) paintOrNil() skia.Paint {
	if p == nil {
		return nil
	}
	return p.paint
}

// Clone the Paint.
func (p *Paint) Clone() *Paint {
	return newPaint(skia.PaintClone(p.paint))
}

// Equivalent returns true if these Paint objects are equivalent.
func (p *Paint) Equivalent(other *Paint) bool {
	if p == nil {
		return other == nil
	}
	if other == nil {
		return false
	}
	return skia.PaintEquivalent(p.paint, other.paint)
}

// Reset the Paint back to its default state.
func (p *Paint) Reset() {
	skia.PaintReset(p.paint)
}

// Antialias returns true if pixels on the active edges of a path may be drawn with partial transparency.
func (p *Paint) Antialias() bool {
	return skia.PaintIsAntialias(p.paint)
}

// SetAntialias sets whether pixels on the active edges of a path may be drawn with partial transparency.
func (p *Paint) SetAntialias(enabled bool) {
	skia.PaintSetAntialias(p.paint, enabled)
}

// Dither returns true if color error may be distributed to smooth color transition.
func (p *Paint) Dither() bool {
	return skia.PaintIsDither(p.paint)
}

// SetDither sets whether color error may be distributed to smooth color transition.
func (p *Paint) SetDither(enabled bool) {
	skia.PaintSetDither(p.paint, enabled)
}

// Color returns the current color.
func (p *Paint) Color() Color {
	return Color(skia.PaintGetColor(p.paint))
}

// SetColor sets the color.
func (p *Paint) SetColor(color Color) {
	skia.PaintSetColor(p.paint, skia.Color(color))
}

// Style returns the current PaintStyle.
func (p *Paint) Style() PaintStyle {
	return PaintStyle(skia.PaintGetStyle(p.paint))
}

// SetStyle sets the PaintStyle.
func (p *Paint) SetStyle(style PaintStyle) {
	skia.PaintSetStyle(p.paint, skia.PaintStyle(style))
}

// StrokeWidth returns the current stroke width.
func (p *Paint) StrokeWidth() float32 {
	return skia.PaintGetStrokeWidth(p.paint)
}

// SetStrokeWidth sets the stroke width.
func (p *Paint) SetStrokeWidth(width float32) {
	skia.PaintSetStrokeWidth(p.paint, width)
}

// StrokeMiter returns the current stroke miter limit for sharp corners.
func (p *Paint) StrokeMiter() float32 {
	return skia.PaintGetStrokeMiter(p.paint)
}

// SetStrokeMiter sets the miter limit for sharp corners.
func (p *Paint) SetStrokeMiter(miter float32) {
	skia.PaintSetStrokeMiter(p.paint, miter)
}

// StrokeCap returns the current StrokeCap.
func (p *Paint) StrokeCap() StrokeCap {
	return StrokeCap(skia.PaintGetStrokeCap(p.paint))
}

// SetStrokeCap sets the StrokeCap.
func (p *Paint) SetStrokeCap(strokeCap StrokeCap) {
	skia.PaintSetStrokeCap(p.paint, skia.StrokeCap(strokeCap))
}

// StrokeJoin returns the current StrokeJoin.
func (p *Paint) StrokeJoin() StrokeJoin {
	return StrokeJoin(skia.PaintGetStrokeJoin(p.paint))
}

// SetStrokeJoin sets the StrokeJoin.
func (p *Paint) SetStrokeJoin(strokeJoin StrokeJoin) {
	skia.PaintSetStrokeJoin(p.paint, skia.StrokeJoin(strokeJoin))
}

// BlendMode returns the current BlendMode.
func (p *Paint) BlendMode() BlendMode {
	return BlendMode(skia.PaintGetBlendMode(p.paint))
}

// SetBlendMode sets the BlendMode.
func (p *Paint) SetBlendMode(blendMode BlendMode) {
	skia.PaintSetBlendMode(p.paint, skia.BlendMode(blendMode))
}

// Shader returns the current Shader.
func (p *Paint) Shader() *Shader {
	return newShader(skia.PaintGetShader(p.paint))
}

// SetShader sets the Shader.
func (p *Paint) SetShader(shader *Shader) {
	skia.PaintSetShader(p.paint, shader.shaderOrNil())
}

// ColorFilter returns the current ColorFilter.
func (p *Paint) ColorFilter() *ColorFilter {
	return newColorFilter(skia.PaintGetColorFilter(p.paint))
}

// SetColorFilter sets the ColorFilter.
func (p *Paint) SetColorFilter(filter *ColorFilter) {
	skia.PaintSetColorFilter(p.paint, filter.filterOrNil())
}

// MaskFilter returns the current MaskFilter.
func (p *Paint) MaskFilter() *MaskFilter {
	return newMaskFilter(skia.PaintGetMaskFilter(p.paint))
}

// SetMaskFilter sets the MaskFilter.
func (p *Paint) SetMaskFilter(filter *MaskFilter) {
	skia.PaintSetMaskFilter(p.paint, filter.filterOrNil())
}

// ImageFilter returns the current ImageFilter.
func (p *Paint) ImageFilter() *ImageFilter {
	return newImageFilter(skia.PaintGetImageFilter(p.paint))
}

// SetImageFilter sets the ImageFilter.
func (p *Paint) SetImageFilter(filter *ImageFilter) {
	skia.PaintSetImageFilter(p.paint, filter.filterOrNil())
}

// PathEffect returns the current PathEffect.
func (p *Paint) PathEffect() *PathEffect {
	return newPathEffect(skia.PaintGetPathEffect(p.paint))
}

// SetPathEffect sets the PathEffect.
func (p *Paint) SetPathEffect(effect *PathEffect) {
	skia.PaintSetPathEffect(p.paint, effect.effectOrNil())
}

// FillPath returns a path representing the path if it was stroked. resScale determines the precision used. Values >1
// increase precision, while those <1 reduce precision to favor speed and size. If hairline returns true, the path
// represents a hairline, otherwise it represents a fill.
func (p *Paint) FillPath(path *Path, resScale float32) (result *Path, hairline bool) {
	result = NewPath()
	isFill := skia.PaintGetFillPath(p.paint, path.path, result.path, nil, resScale)
	return result, !isFill
}

// FillPathWithCull returns a path representing the path if it was stroked. cullRect will prune any parts outside of the
// rect. resScale determines the precision used. Values >1 increase precision, while those <1 reduce precision to favor
// speed and size. If hairline returns true, the path represents a hairline, otherwise it represents a fill.
func (p *Paint) FillPathWithCull(path *Path, cullRect Rect, resScale float32) (result *Path, hairline bool) {
	result = NewPath()
	isFill := skia.PaintGetFillPath(p.paint, path.path, result.path, skia.RectToSkRect(&cullRect), resScale)
	return result, !isFill
}
