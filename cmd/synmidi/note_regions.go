package main

import (
	"errors"
	"image"
	"image/color"
)

func detectNoteRegion(src image.Image, keyMidpoint int) (noteRegionStartY, noteRegionEndY int, err error) {
	logSection("Detecting note region..")
	b := src.Bounds()

	row := func(y int) []color.Gray {
		r := make([]color.Gray, b.Dx())
		for x := 0; x < b.Dx(); x++ {
			c := color.GrayModel.Convert(src.At(x, y)).(color.Gray)
			r[x] = c
		}
		return r
	}

	rowDist := func(r1, r2 []color.Gray) int {
		var sum int
		for x := 0; x < len(r1); x++ {
			c1 := r1[x]
			c2 := r2[x]
			if c1.Y != c2.Y {
				sum++
			}
		}
		return sum
	}

	var last []color.Gray
	for y := 0; y < b.Max.Y; y++ {
		r := row(y)
		if last == nil {
			last = r
			continue
		}
		dist := rowDist(last, r)
		last = r

		if y > 300 && dist > 1000 {
			if debugImage != nil {
				drawLineX(debugImage, 0, 0, b.Dx(), color.Black)
				drawLabel(debugImage, 200, 0, "^ note region start")

				drawLineX(debugImage, 0, y, b.Dx(), color.Black)
				drawLabel(debugImage, 200, y, "^ note region end")
			}
			return 0, y, nil
		}
	}
	return 0, 0, errors.New("failed to detect note region")
}
