package pong

import (
	"image/color"
	_ "log"
	"math"
)

// Color used in game
type RGBA color.RGBA

// Defines the z order of different Drawable
type ZIndex int

// Methods required to draw something
type Drawable interface {

	// Computes the color with the given baseColor
	ColorAt(position float64, baseColor RGBA) RGBA

	// The ZIndex of this Drawable thing
	ZIndex() ZIndex

	// Move this Drawable forward in time by dt, returns keepAlive
	Animate(dt float64) (keepAlive bool)
}

// Line that can be drawn on the board
type Line struct {
	// Extends of the line
	leftEdge, rightEdge float64

	color RGBA

	zindex ZIndex
}

var testLine Drawable = Line{}

// Construct a Line
func NewLine(leftEdge, rightEdge float64, color RGBA, zindex ZIndex) Line {
	return Line{
		leftEdge:  leftEdge,
		rightEdge: rightEdge,
		color:     color,
		zindex:    zindex,
	}
}

// Returns the color at position blended on top of baseColor
func (line Line) ColorAt(position float64, baseColor RGBA) RGBA {

	if line.leftEdge < position && position < line.rightEdge {
		return line.color
	}

	return baseColor
}

// ZIndex of line
func (line Line) ZIndex() ZIndex {
	return line.zindex
}

// Animate line
func (line Line) Animate(dt float64) bool {
	return true
}

// Represents a background animation of a sinusoid moving forward
type Sinusoid struct {
	scale float64

	originalColor RGBA

	// offset related to time passing
	offset float64

	zindex ZIndex
}

var testSinusoid Drawable = &Sinusoid{}

// Construct a Sinusoid
func NewSinusoid(scale float64, originalColor RGBA, zindex ZIndex) *Sinusoid {
	return &Sinusoid{
		scale:         scale,
		originalColor: originalColor,
		zindex:        zindex,
	}
}

// Returns the color at position blended on top of baseColor
func (this *Sinusoid) ColorAt(position float64, baseColor RGBA) RGBA {

	sine := (math.Sin((position+this.offset)*this.scale) + 1.0) / 2.0

	return RGBA{
		uint8(float64(this.originalColor.R) * sine),
		uint8(float64(this.originalColor.G) * sine),
		uint8(float64(this.originalColor.B) * sine),
		this.originalColor.A,
	}
}

// ZIndex of line
func (this *Sinusoid) ZIndex() ZIndex {
	return this.zindex
}

// Animate line
func (this *Sinusoid) Animate(dt float64) bool {

	this.offset += dt

	return true
}

// Represents a background animation of moving through the HSL color space
type HSLWheel struct {

	// Hwo to scale lumniosity so that it's event spread across all points
	scale float64

	// goes from 0 to 1 and then wraps
	hue float64

	zindex ZIndex
}

var testHSLWheel Drawable = &HSLWheel{}

// Construct an HSLWheel
func NewHSLWheel(field *GameField, zindex ZIndex) *HSLWheel {
	return &HSLWheel{
		scale:  float64(field.width),
		hue:    0.0,
		zindex: zindex,
	}
}

// Returns the color at position blended on top of baseColor
func (this *HSLWheel) ColorAt(position float64, baseColor RGBA) RGBA {

	luminosity := position / this.scale

	// shift it up because we don't care much about the very dark colors
	luminosity = luminosity*0.8 + 0.2

	return hslToRGB(this.hue, 1.0, luminosity)
}

// Convert HSL to RGB, based on http://mjijackson.com/2008/02/rgb-to-hsl-and-rgb-to-hsv-color-model-conversion-algorithms-in-javascript
func hslToRGB(hue, saturation, luminosity float64) RGBA {

	var red, green, blue uint8

	if saturation == 0 {
		red = 255
		green = 255
		blue = 255
	} else {

		hueToRGB := func(p, q, t float64) float64 {
			if t < 0.0 {
				t += 1.0
			}
			if t > 1.0 {
				t -= 1.0
			}

			if t < 1.0/6.0 {
				return p + (q-p)*6.0*t
			}
			if t < 1.0/2.0 {
				return q
			}
			if t < 2.0/3.0 {
				return p + (q-p)*(2.0/3.0-t)*6.0
			}
			return p
		}

		var q float64
		if luminosity < 0.5 {
			q = luminosity * (1.0 + saturation)
		} else {
			q = luminosity + saturation - luminosity*saturation
		}

		p := 2*luminosity - q
		red = uint8(hueToRGB(p, q, hue+1.0/3.0) * 255.0)
		green = uint8(hueToRGB(p, q, hue) * 255.0)
		blue = uint8(hueToRGB(p, q, hue-1.0/3.0) * 255.0)
	}

	//log.Print("Converted ", hue, saturation, luminosity, " to ", red, green, blue)

	return RGBA{red, green, blue, 255}
}

// ZIndex of line
func (this *HSLWheel) ZIndex() ZIndex {
	return this.zindex
}

// Animate line
func (this *HSLWheel) Animate(dt float64) bool {

	this.hue += dt * 0.1

	if this.hue > 1.0 {
		this.hue -= 1.0
	}

	//log.Print("Hue", this.hue)

	return true
}
