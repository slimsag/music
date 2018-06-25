package main

import (
	"fmt"
)

func logSection(name string) {
	fmt.Println("-", name)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(a, b int) int {
	return abs(a - b)
}

func distu8(a, b uint8) uint8 {
	if a < b {
		return b - a
	}
	return a - b
}

func distu32(a, b uint32) uint32 {
	if a < b {
		return b - a
	}
	return a - b
}

/*
var indexToPitch []*mxl.Pitch

func init() {
	indexToPitch = append(indexToPitch, []*mxl.Pitch{
		&mxl.Pitch{Step: "A", Octave: 0},
		&mxl.Pitch{Step: "A", Octave: 0, Accidental: 1},
		&mxl.Pitch{Step: "B", Octave: 0},
	}...)
	for octave := 1; octave <= 7; octave++ {
		indexToPitch = append(indexToPitch, []*mxl.Pitch{
			&mxl.Pitch{Step: "C", Octave: octave},
			&mxl.Pitch{Step: "C", Octave: octave, Accidental: 1},
			&mxl.Pitch{Step: "D", Octave: octave},
			&mxl.Pitch{Step: "D", Octave: octave, Accidental: 1},
			&mxl.Pitch{Step: "E", Octave: octave},
			&mxl.Pitch{Step: "F", Octave: octave},
			&mxl.Pitch{Step: "F", Octave: octave, Accidental: 1},
			&mxl.Pitch{Step: "G", Octave: octave},
			&mxl.Pitch{Step: "G", Octave: octave, Accidental: 1},
			&mxl.Pitch{Step: "A", Octave: octave},
			&mxl.Pitch{Step: "A", Octave: octave, Accidental: 1},
			&mxl.Pitch{Step: "B", Octave: octave},
		}...)
	}
	indexToPitch = append(indexToPitch, &mxl.Pitch{Step: "C", Octave: 8})
}
*/

var cToIndex = map[string]int{
	"C1": 3,
	"C2": 3 + 12,
	"C3": 3 + (12 * 2),
	"C4": 3 + (12 * 3),
	"C5": 3 + (12 * 4),
	"C6": 3 + (12 * 5),
	"C7": 3 + (12 * 6),
	"C*": 3 + (12 * 7),
}

func frameInfosEqual(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

/*
func measureEqual(a, b mxl.Measure) bool {
	if len(a.Notes) != len(b.Notes) {
		return false
	}
	for i, a := range a.Notes {
		b := b.Notes[i]
		if *a.Pitch != *b.Pitch {
			return false
		}
		if a.Duration != b.Duration {
			return false
		}
		if a.Type != b.Type {
			return false
		}
		if (a.Chord != nil) != (b.Chord != nil) {
			return false
		}
	}
	return true
}
*/

func any(keys []bool) bool {
	for _, k := range keys {
		if k {
			return true
		}
	}
	return false
}
