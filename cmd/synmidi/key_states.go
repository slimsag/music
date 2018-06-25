package main

import "fmt"

type key struct {
	state      keyState
	isLeftHand bool
}

type keyState int

const (
	keyStateNone         keyState = iota
	keyStatePressed      keyState = iota
	keyStateContinued    keyState = iota
	keyStatePressedAgain keyState = iota
	keyStateReleased     keyState = iota
)

type keyList []key

func (k keyList) debugPrint() {
	for _, k := range k {
		switch k.state {
		case keyStateNone:
			fmt.Printf(" ")
		case keyStatePressed:
			if k.isLeftHand {
				fmt.Printf("↙")
			} else {
				fmt.Printf("↘")
			}
		case keyStateContinued:
			if k.isLeftHand {
				fmt.Printf("L")
			} else {
				fmt.Printf("R")
			}
		case keyStatePressedAgain:
			if k.isLeftHand {
				fmt.Printf("⇐")
			} else {
				fmt.Printf("⇒")
			}
		case keyStateReleased:
			if k.isLeftHand {
				fmt.Printf("↖")
			} else {
				fmt.Printf("↗")
			}
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
func determineKeyStates(frames []*frameInfo) []keyList {
	results := make([]keyList, len(frames))
	for frameIndex, frame := range frames {
		frameResult := make(keyList, len(frame.Keys))

		if frameIndex == 0 || frameIndex+1+1 >= len(frames) {
			// Too early/late to do any lookbehind / lookahead. These keys do
			// not need correction anyways.
			for keyIndex, k := range frame.Keys {
				if k {
					frameResult[keyIndex] = key{state: keyStatePressed, isLeftHand: frame.IsLeftHandKey[keyIndex]}
				} else {
					frameResult[keyIndex] = key{state: keyStateReleased, isLeftHand: frame.IsLeftHandKey[keyIndex]}
				}
			}
			results[frameIndex] = frameResult
			continue
		}

		for keyIndex, k := range frame.Keys {
			downCurrentFrame := k
			downPreviousFrame := frames[frameIndex-1].Keys[keyIndex]
			downNextNextFrame := frames[frameIndex+1+1].Keys[keyIndex]

			calculatedDownPreviousFrame := false
			if len(results) > 0 {
				lastFrame := results[frameIndex-1][keyIndex]
				calculatedDownPreviousFrame = lastFrame.state == keyStatePressed || lastFrame.state == keyStatePressedAgain || lastFrame.state == keyStateContinued
			}

			// TODO(slimsag): If a key starts on the left hand and gets PressedAgain with the
			// right hand, we would not send the proper states. Fixing this would probably require
			// separating keyStatePressedAgain events entirely.
			isLeftHand := frame.IsLeftHandKey[keyIndex]
			if !k {
				// Key was not pressed down, so we had no way to know if it
				// was really a left key or not.
				if downNextNextFrame {
					isLeftHand = frames[frameIndex+1+1].IsLeftHandKey[keyIndex]
				} else if downPreviousFrame {
					isLeftHand = frames[frameIndex-1].IsLeftHandKey[keyIndex]
				}
			}

			switch {
			case !downPreviousFrame && downCurrentFrame && !calculatedDownPreviousFrame:
				frameResult[keyIndex] = key{state: keyStatePressed, isLeftHand: isLeftHand}
			case downPreviousFrame && !downCurrentFrame && downNextNextFrame:
				frameResult[keyIndex] = key{state: keyStatePressedAgain, isLeftHand: isLeftHand}
			case downPreviousFrame && !downCurrentFrame && !downNextNextFrame:
				frameResult[keyIndex] = key{state: keyStateReleased, isLeftHand: isLeftHand}
			case (downPreviousFrame || calculatedDownPreviousFrame) && (downCurrentFrame || downNextNextFrame):
				frameResult[keyIndex] = key{state: keyStateContinued, isLeftHand: isLeftHand}
			}
		}
		results[frameIndex] = frameResult
	}
	return results
}
