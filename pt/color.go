package pt

import (
	"image/color"
	"math"
)

var (
	Black = Color{0, 0, 0, 1}
	White = Color{1, 1, 1, 1}
)

type Color struct {
	R, G, B, A float64
}

func HexColor(x int) Color {
	a := float64((x>>24)&0xff) / 255
	r := float64((x>>16)&0xff) / 255
	g := float64((x>>8)&0xff) / 255
	b := float64((x>>0)&0xff) / 255
	return Color{r, g, b, a}.Pow(2.2)
}

func Kelvin(K float64) Color {
	var red, green, blue float64
	// red
	if K >= 6600 {
		a := 351.97690566805693
		b := 0.114206453784165
		c := -40.25366309332127
		x := K/100 - 55
		red = a + b*x + c*math.Log(x)
	} else {
		red = 255
	}
	// green
	if K >= 6600 {
		a := 325.4494125711974
		b := 0.07943456536662342
		c := -28.0852963507957
		x := K/100 - 50
		green = a + b*x + c*math.Log(x)
	} else if K >= 1000 {
		a := -155.25485562709179
		b := -0.44596950469579133
		c := 104.49216199393888
		x := K/100 - 2
		green = a + b*x + c*math.Log(x)
	} else {
		green = 0
	}
	// blue
	if K >= 6600 {
		blue = 255
	} else if K >= 2000 {
		a := -254.76935184120902
		b := 0.8274096064007395
		c := 115.67994401066147
		x := K/100 - 10
		blue = a + b*x + c*math.Log(x)
	} else {
		blue = 0
	}
	red = math.Min(1, red/255)
	green = math.Min(1, green/255)
	blue = math.Min(1, blue/255)
	return Color{red, green, blue, 1}
}

func NewColor(c color.Color) Color {
	r, g, b, _ := c.RGBA()
	return Color{float64(r) / 65535, float64(g) / 65535, float64(b) / 65535, 1}
}

func (c Color) RGBA() color.RGBA {
	r := uint8(math.Max(0, math.Min(255, c.R*255)))
	g := uint8(math.Max(0, math.Min(255, c.G*255)))
	b := uint8(math.Max(0, math.Min(255, c.B*255)))
	a := uint8(math.Max(0, math.Min(255, c.A*255)))
	return color.RGBA{r, g, b, a}
}

func (c Color) RGBA64() color.RGBA64 {
	r := uint16(math.Max(0, math.Min(65535, c.R*65535)))
	g := uint16(math.Max(0, math.Min(65535, c.G*65535)))
	b := uint16(math.Max(0, math.Min(65535, c.B*65535)))
	a := uint16(math.Max(0, math.Min(65535, c.A*65535)))
	return color.RGBA64{r, g, b, a}
}

func (c Color) Add(b Color) Color {
	return Color{c.R + b.R, c.G + b.G, c.B + b.B, c.A + b.A}
}

func (c Color) Sub(b Color) Color {
	return Color{c.R - b.R, c.G - b.G, c.B - b.B, c.A - b.A}
}

func (c Color) Mul(b Color) Color {
	return Color{c.R * b.R, c.G * b.G, c.B * b.B, c.A * b.A}
}

func (c Color) MulScalar(b float64) Color {
	return Color{c.R * b, c.G * b, c.B * b, c.A * b}
}

func (c Color) DivScalar(b float64) Color {
	return Color{c.R / b, c.G / b, c.B / b, c.A / b}
}

func (c Color) Min(b Color) Color {
	return Color{math.Min(c.R, b.R), math.Min(c.G, b.G), math.Min(c.B, b.B), math.Min(c.A, b.A)}
}

func (c Color) Max(b Color) Color {
	return Color{math.Max(c.R, b.R), math.Max(c.G, b.G), math.Max(c.B, b.B), math.Min(c.A, b.A)}
}

func (c Color) MinComponent() float64 {
	return math.Min(math.Min(c.R, c.G), c.B)
}

func (c Color) MaxComponent() float64 {
	return math.Max(math.Max(c.R, c.G), c.B)
}

func (c Color) Pow(b float64) Color {
	return Color{math.Pow(c.R, b), math.Pow(c.G, b), math.Pow(c.B, b), math.Pow(c.A, b)}
}

func (c Color) Mix(b Color, pct float64) Color {
	c = c.MulScalar(1 - pct)
	b = b.MulScalar(pct)
	return c.Add(b)
}
