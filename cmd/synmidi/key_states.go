package main

import "fmt"

type keyState int

const (
	keyStateNone         keyState = iota
	keyStatePressed      keyState = iota
	keyStateContinued    keyState = iota
	keyStatePressedAgain keyState = iota
	keyStateReleased     keyState = iota
)

type keyStateList []keyState

func (k keyStateList) debugPrint() {
	for _, k := range k {
		switch k {
		case keyStateNone:
			fmt.Printf(" ")
		case keyStatePressed:
			fmt.Printf("↓")
		case keyStateContinued:
			fmt.Printf("|")
		case keyStatePressedAgain:
			fmt.Printf("↕")
		case keyStateReleased:
			fmt.Printf("↑")
		default:
			fmt.Printf("�")
		}
	}
	fmt.Println("")
}

// determineKeyStates determines the note start/continue/stop events for the given frames.
//
// Synthesia indicates to us on the piano keys that a note has changed by
// making the key flash. For example, if A0 is pressed, and then pressed again,
// Synthesia will add a "blank" frame inbetween where that key is not pressed:
//
// 	0: -X------X--X----X---X--X-
// 	1: -X------X--X----X---X--X-
// 	2: -X------X--X----X---X--X-
// 	3: ---X--------------------- <- Blank frame added for notes that were also held down on the previous frame.
// 	4: ---X----X--X----X---X--X-
// 	5: ---X----X--X----X---X--X-
// 	6: ---X----X--X----X---X--X-
//
// Unfortunetly, Synthesia isn't super precise about this (or some video
// recordings of it aren't, at least). As such, there can sometimes be a number
// of these blank frames:
//
// 	0: -X------X--X----X---X--X-
// 	1: -X------X--X----X---X--X-
// 	2: -X------X--X----X---X--X-
// 	3: ---X--------------------- <- Blank frame added for notes that were also held down on the previous frame.
// 	4: ---X--------------------- <- Blank frame added for notes that were also held down on the previous frame.
// 	5: ---X----X--X----X---X--X-
// 	6: ---X----X--X----X---X--X-
//
// This functions takes into account the above in order to produce a list of
// key states.
func determineKeyStates(frames []*frameInfo) []keyStateList {
	results := make([]keyStateList, len(frames))
	for frameIndex, frame := range frames {
		frameResult := make(keyStateList, len(frame.Keys))

		if frameIndex == 0 || frameIndex+1+1 >= len(frames) {
			// Too early/late to do any lookbehind / lookahead. These keys do
			// not need correction anyways.
			for keyIndex, key := range frame.Keys {
				if key {
					frameResult[keyIndex] = keyStatePressed
				} else {
					frameResult[keyIndex] = keyStateReleased
				}
			}
			results[frameIndex] = frameResult
			continue
		}

		for keyIndex, key := range frame.Keys {
			downCurrentFrame := key
			downPreviousFrame := frames[frameIndex-1].Keys[keyIndex]
			downNextNextFrame := frames[frameIndex+1+1].Keys[keyIndex]

			calculatedDownPreviousFrame := false
			if len(results) > 0 {
				lastFrame := results[frameIndex-1][keyIndex]
				calculatedDownPreviousFrame = lastFrame == keyStatePressed || lastFrame == keyStatePressedAgain || lastFrame == keyStateContinued
			}

			switch {
			case !downPreviousFrame && downCurrentFrame && !calculatedDownPreviousFrame:
				frameResult[keyIndex] = keyStatePressed
			case downPreviousFrame && !downCurrentFrame && downNextNextFrame:
				frameResult[keyIndex] = keyStatePressedAgain
			case downPreviousFrame && !downCurrentFrame && !downNextNextFrame:
				frameResult[keyIndex] = keyStateReleased
			case (downPreviousFrame || calculatedDownPreviousFrame) && (downCurrentFrame || downNextNextFrame):
				frameResult[keyIndex] = keyStateContinued
			}
		}
		results[frameIndex] = frameResult
	}
	return results
}
