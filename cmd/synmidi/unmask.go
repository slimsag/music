package main

import (
	"image"
	"image/color"
	"image/draw"
)

func unmask(dst draw.Image, src image.Image, yStart, yStop int) {
	b := src.Bounds()
	for y := yStart; y <= yStop; y++ {
		for x := 0; x < b.Dx(); x++ {
			srcColor := src.At(x, y)
			dstColor := dst.At(x, y)
			if srcColor == dstColor {
				dst.Set(x, y, color.Black)
			} else {
				dst.Set(x, y, color.White)
			}
		}
	}
}
