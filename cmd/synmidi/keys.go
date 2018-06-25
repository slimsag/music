package main

import (
	"image"
	"image/color"
)

func detectKeys(src image.Image, keyMidpoint int) ([]keyInfo, error) {
	logSection("Detecting key centers..")
	b := src.Bounds()

	const debugOffset = 15
	var (
		keys      []keyInfo
		lastPoint image.Point
		lastColor color.Color
	)
	for x := 0; x < b.Max.X; x++ {
		c := color.GrayModel.Convert(src.At(x, keyMidpoint))
		if c.(color.Gray).Y > 127 {
			c = color.White
		} else {
			c = color.Black
		}
		if x == 0 {
			// First iteration. Key has started.
			lastColor = c
			lastPoint = image.Point{X: x, Y: keyMidpoint}
			continue
		}
		if lastColor != c {
			keyStart := lastPoint
			keyEnd := image.Point{X: x, Y: keyMidpoint}
			if debugImage != nil {
				drawLineY(debugImage, keyStart.X, keyStart.Y+debugOffset, 50, color.RGBA{0, 128, 0, 255})
			}

			width := keyEnd.X - keyStart.X

			// Filter out dividers between two white keys.
			if width > 5 {
				keyCenter := image.Point{X: keyStart.X + width/2, Y: keyStart.Y}
				actualColor := src.At(keyCenter.X, keyCenter.Y)
				keys = append(keys, keyInfo{
					center:      keyCenter,
					color:       lastColor,
					actualColor: actualColor,
				})
				if debugImage != nil {
					size := 6
					drawRect(debugImage, keyCenter.X-size, keyCenter.Y+size+debugOffset, size*2, size*2, color.RGBA{128, 0, 0, 255})
					size = 4
					drawRect(debugImage, keyCenter.X-size, keyCenter.Y+4+size+debugOffset, size*2, size*2, actualColor)
				}
			}
			lastColor = c
			lastPoint = keyEnd
		}
		//debugImage.Set(x, keyMidpoint, c)
	}
	return keys, nil
}
