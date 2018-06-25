package main

import (
	"image"
	"image/color"
)

func detectKeyMidpoint(src image.Image) (int, error) {
	logSection("Detecting key midpoint..")
	b := src.Bounds()
	maxHeight := *flagKeyMidpointHeightMax

	yStart := -1
	yEnd := -1
	for y := b.Max.Y - maxHeight; y < b.Max.Y; y++ {
		white := 0
		black := 0
		for x := 0; x < b.Max.X; x++ {
			g := color.GrayModel.Convert(src.At(x, y)).(color.Gray)
			if g.Y > 127 {
				white++
			} else {
				black++
			}
		}
		if dist(white, black) < *flagKeyMidpointTolerence {
			if yStart == -1 {
				yStart = y
			}
			yEnd = y
		}
	}

	middle := yStart + (dist(yStart, yEnd) / 2)
	if debugImage != nil {
		drawLineX(debugImage, 0, yStart, b.Dx(), color.Black)
		drawLabel(debugImage, 0, yStart, "^ top of black keys")

		drawLineX(debugImage, 0, yEnd, b.Dx(), color.Black)
		drawLabel(debugImage, 0, yEnd, "^ bottom of black keys")

		drawLineX(debugImage, 0, middle, b.Dx(), color.Black)
		drawLabel(debugImage, 0, middle, "^ middle of black keys")
	}
	return middle, nil
}
