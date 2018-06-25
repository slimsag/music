package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func writePNG(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

func drawLabel(img draw.Image, x, y int, label string) {
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.White),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	bounds, _ := d.BoundString(label)
	d.Dot.Y -= (bounds.Min.Y - bounds.Max.Y)

	wx := int(bounds.Max.X-bounds.Min.X) / 64
	wy := int(bounds.Max.Y-bounds.Min.Y) / 64
	drawRect(img, x, y, wx, wy, color.Black)

	d.DrawString(label)
}

func drawLineX(img draw.Image, x, y, width int, col color.Color) {
	for v := x; v <= x+width; v++ {
		img.Set(v, y, col)
	}
}

func drawLineY(img draw.Image, x, y, height int, col color.Color) {
	for v := y; v <= y+height; v++ {
		img.Set(x, v, col)
	}
}

func drawRect(img draw.Image, x, y, width, height int, col color.Color) {
	for v := x; v <= x+width; v++ {
		drawLineY(img, v, y, height, col)
	}
}
