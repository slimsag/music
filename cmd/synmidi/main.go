package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/otium/ytdl"

	"github.com/gomidi/midi/midimessage/channel" // (Channel Messages)
	"github.com/gomidi/midi/midimessage/meta"    // (Meta Messages)
	"github.com/gomidi/midi/smf"
	"github.com/gomidi/midi/smf/smfwriter"
	// you may also want to use these
	// github.com/gomidi/midi/midimessage/cc         (ControlChange Messages)
	// github.com/gomidi/midi/midimessage/sysex      (System Exclusive Messages)
)

var (
	flagKeyMidpoint          = flag.Int("key-midpoint", -1, "force a key midpoint (vertical center point of black and white keys)")
	flagKeyMidpointHeightMax = flag.Int("key-midpoint-height-max", 216, "max number of pixels from the bottom of a frame to search for the key midpoint")
	flagKeyMidpointTolerence = flag.Int("key-midpoint-tolerence", 150, "tolerence level for matching black/white key regions. Lower == stricter, higher == looser")
	flagDebug                = flag.String("debug", "", "path to write a debug image (e.g. './debug.png')")

	debugImage *image.RGBA
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	switch flag.Args()[0] {
	case "download":
		url := os.Args[2]
		log.Printf("Downloading %s...", url)
		vid, err := ytdl.GetVideoInfo(url)
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create("out.mp4")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		vid.Download(vid.Formats.Best(ytdl.FormatVideoEncodingKey)[0], file)
	case "frames":
		if err := os.RemoveAll("frames"); err != nil {
			log.Fatal(err)
		}
		if err := os.MkdirAll("frames", 0777); err != nil {
			log.Fatal(err)
		}
		//skipFrames := os.Args[2]
		// "-vf", fmt.Sprintf(`select=gte(n\,%s)`, skipFrames), "-r", "1/2.5", "-vsync", "vfr",
		cmd := exec.Command("ffmpeg", "-i", "./out.mp4", "frames/%05d.png")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	case "reduce":
		frameRate, err := detectFrameRate("./out.mp4")
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("frame rate:", frameRate.FPS())
		infos, err := ioutil.ReadDir("frames")
		if err != nil {
			log.Fatal(err)
		}
		if err := os.MkdirAll("frames/reduced", 0777); err != nil {
			log.Fatal(err)
		}
		songFrame := 0
		firstFrame := true
		var baseInfo *info
		lastUpdate := time.Now()
		var frameInfos []*frameInfo
		for _, info := range infos {
			if filepath.Ext(info.Name()) != ".png" {
				continue
			}
			path := filepath.Join("frames", info.Name())

			if firstFrame {
				firstFrame = false

				baseInfo, err = detectInfo(path)
				if err != nil {
					log.Fatal(err)
				}
			}

			frameInfo, err := detectFrameInfo(baseInfo, path)
			if err != nil {
				log.Fatal(err)
			}

			if songFrame != -1 || any(frameInfo.Keys) {
				songFrame++
				// Song has started previously, i.e. first piano key has been
				// hit.
				fs := float64(time.Second)
				frameInfo.Time = time.Duration((float64(songFrame) * fs) / frameRate.FPS())
				frameInfo.debugPrint()
				frameInfos = append(frameInfos, frameInfo)
				//if songFrame > 200 {
				//	break
				//}
			}

			if false {
				if err := os.MkdirAll(filepath.Join("debug"), 0777); err != nil {
					log.Fatal(err)
				}
				if err := writePNG(filepath.Join("debug", info.Name()), frameInfo.debug); err != nil {
					log.Fatal(err)
				}
			}
			if time.Since(lastUpdate) > 1*time.Second {
				lastUpdate = time.Now()
				fmt.Println("Processing frame", info.Name())
			}
			/*
				lineAtTop, err := hasLineAtTop(filepath.Join("frames", info.Name()))
				if err != nil {
					log.Fatal(err)
				}
				if lineAtTop {
					fmt.Println(lineAtTop, info.Name())
					cmd := exec.Command("cp", filepath.Join("frames", info.Name()), filepath.Join("frames/reduced", info.Name()))
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						log.Fatal(err)
					}
				}
			*/
		}
		data, err := json.MarshalIndent(frameInfos, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile("frameInfos.json", data, 0777); err != nil {
			log.Fatal(err)
		}
	/*
		case "mxl":
			// Decode the JSON file.
			f, err := os.Open("frameInfos.json")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			var frameInfos []*frameInfo
			if err := json.NewDecoder(f).Decode(&frameInfos); err != nil {
				log.Fatal(err)
			}

			keyBase := flag.Arg(1)
			split := strings.Split(keyBase, "=")
			cn := split[0]                        // C0, C1, C2, etc
			relIndex, _ := strconv.Atoi(split[1]) // Where C[n] is at, relative to the frame, starting at 0
			finalKeyOffset := cToIndex[cn] - relIndex
			fmt.Println(finalKeyOffset)

			doc := &mxl.Document{}

			doc.PartList = &mxl.PartList{
				ScorePart: &mxl.ScorePart{ID: "P1", PartName: &mxl.String{Value: "Piano"}},
			}

			part := mxl.Part{ID: "P1"}
			for i, f := range frameInfos {
				//if i > 200 {
				//	break
				//}
				//if i > 0 && frameInfosEqual(f.Keys, frameInfos[i-1].Keys) {
				//	continue
				//}

				measure := mxl.Measure{Number: i}
				prev := false
				for keyIndex, k := range f.Keys {
					if k {
						//fmt.Println(finalKeyOffset + keyIndex)
						note := mxl.Note{
							Pitch:    indexToPitch[finalKeyOffset+keyIndex],
							Duration: 12,
							Type:     "eighth",
						}
						if prev {
							note.Chord = &xml.Name{}
						}
						prev = true
						measure.Notes = append(measure.Notes, note)
						//break
					}
				}
				//noChange := false
				//if len(part.Measures) > 0 {
				//	if measureEqual(part.Measures[len(part.Measures)-1], measure) {
				//		noChange = true
				//	}
				//}
				//_ = noChange
				if len(measure.Notes) > 0 { //&& !noChange {
					part.Measures = append(part.Measures, measure)
				}
				f.debugPrint()
			}
			doc.Parts = append(doc.Parts, part)

			if err := mxl.MarshalFile("out.xml", doc); err != nil {
				log.Fatal(err)
			}
	*/
	case "midi":
		// Decode the JSON file.
		f, err := os.Open("frameInfos.json")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		var frameInfos []*frameInfo
		if err := json.NewDecoder(f).Decode(&frameInfos); err != nil {
			log.Fatal(err)
		}

		keyBase := flag.Arg(1)
		split := strings.Split(keyBase, "=")
		cn := split[0]                        // C0, C1, C2, etc
		relIndex, _ := strconv.Atoi(split[1]) // Where C[n] is at, relative to the frame, starting at 0
		finalKeyOffset := cToIndex[cn] - relIndex
		finalKeyOffset += 21 // midi index for key A0
		fmt.Println(finalKeyOffset)

		timeCode := smf.SMPTE30(30)
		writeMIDI := func(wr smf.Writer) {
			wr.SetDelta(0)

			for frameIndex, keyStates := range determineKeyStates(frameInfos) {
				for keyIndex, keyState := range keyStates {
					midiNote := uint8(finalKeyOffset + keyIndex)
					velocity := uint8(90)
					switch keyState {
					case keyStatePressed:
						_, err := wr.Write(channel.Ch2.NoteOn(midiNote, velocity))
						if err != nil {
							log.Fatal(err)
						}
					case keyStateContinued, keyStateNone:
						// Must send a message or else SetDelta takes no effect.
						_, err := wr.Write(channel.Ch2.AfterTouch(255))
						if err != nil {
							log.Fatal(err)
						}
					case keyStatePressedAgain:
						_, err = wr.Write(channel.Ch2.NoteOff(midiNote))
						if err != nil {
							log.Fatal(err)
						}
						wr.SetDelta(1)
						_, err := wr.Write(channel.Ch2.NoteOn(midiNote, velocity))
						if err != nil {
							log.Fatal(err)
						}
					case keyStateReleased:
						_, err = wr.Write(channel.Ch2.NoteOff(midiNote))
						if err != nil {
							log.Fatal(err)
						}
					}
				}
				frameInfos[frameIndex].debugPrint()
				keyStates.debugPrint()
				wr.SetDelta(1 * uint32(timeCode.SubFrames))
			}

			// finishes the first track and writes it to the file
			_, err = wr.Write(meta.EndOfTrack)
			if err != nil {
				log.Fatal(err)
			}

			// the next write writes to the second track
			// after writing delta is always 0, so we start here at the beginning of the second track
			_, err = wr.Write(meta.Text("hello second track!"))
			if err != nil {
				log.Fatal(err)
			}

			// finishes the second track and writes it to the file
			_, err = wr.Write(meta.EndOfTrack)
			if err != nil {
				log.Fatal(err)
			}
		}

		// the number passed to the NumTracks option must match the tracks written
		// if NumTracks is not passed, it defaults to single track (SMF0)
		// if numtracks > 1, SMF1 format is chosen.
		// if TimeFormat is not passed, smf.MetricTicks(960) will be chosen
		err = smfwriter.WriteFile("file.mid", writeMIDI, smfwriter.NumTracks(2), smfwriter.TimeFormat(timeCode))
		if err != nil {
			log.Fatal(err)
		}

	}
}

type keyInfo struct {
	center      image.Point // center of the key
	color       color.Color // color.White or color.Black
	actualColor color.Color // the actual pixel color at the center
}

type info struct {
	keyMidpoint int
	//noteRegionStartY, noteRegionEndY int // TODO(slimsag): this is not that useful / functional
	baseline int
	keys     []keyInfo
	mask     image.Image
}

func detectInfo(img string) (*info, error) {
	f, err := os.Open(img)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	// Setup debug image by copying the src image.
	if *flagDebug != "" {
		s := src.(*image.RGBA)
		debugImage = &image.RGBA{
			Pix:    make([]uint8, len(s.Pix)),
			Stride: s.Stride,
			Rect:   s.Rect,
		}
		copy(debugImage.Pix, s.Pix)
	}

	info := &info{mask: src}

	// Detect the key midpoint.
	info.keyMidpoint = *flagKeyMidpoint
	if info.keyMidpoint == -1 {
		info.keyMidpoint, err = detectKeyMidpoint(src)
		if err != nil {
			return nil, err
		}
	}

	// Detect the key regions.
	info.keys, err = detectKeys(src, info.keyMidpoint)
	if err != nil {
		return nil, err
	}

	// Detect the note region.
	//info.noteRegionStartY, info.noteRegionEndY, err = detectNoteRegion(src, info.keyMidpoint)
	//if err != nil {
	//	return nil, err
	//}

	// Write debug image.
	if *flagDebug != "" {
		if err := writePNG(*flagDebug, debugImage); err != nil {
			return nil, err
		}
	}
	return info, nil
}

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

	/*
		oneKeyPressed := func(start, end int) bool {
			last := frames[start]
			oneKeyPressed := false
			for i := start + 1; i <= end; i++ {
				oneKeyPressed = oneKeyPressed || any(last.Keys) || any(frames[i].Keys)
				last = frames[i]
			}
			return oneKeyPressed
		}

		framesEqual := func(start, end int) bool {
			last := frames[start]
			for i := start + 1; i <= end; i++ {
				//fmt.Println(i)
				if !frameInfosEqual(last.Keys, frames[i].Keys) {
					return false
				}
				last = frames[i]
			}
			return true
		}

		look := 2
		offset := 2
		for i := range frames {
			if i < offset+look || i > len(frames)-offset-look-1 {
				continue
			}

			//fmt.Println(i, len(frames))
			prevEqual := framesEqual(i-offset-look, i-offset-1)
			prevOneKey := oneKeyPressed(i-offset-look, i-offset-1)
			nextEqual := framesEqual(i+offset+1, i+offset+look)
			nextOneKey := oneKeyPressed(i+offset+1, i+offset+look)
			immediateEqual := framesEqual(i-offset, i) || framesEqual(i, i+offset)
			_ = prevOneKey
			_ = nextOneKey
			_ = immediateEqual
			_ = prevEqual
			_ = nextEqual
			//if prevEqual && nextEqual && !immediateEqual {
			//	frames[i].start = true
			//}

			//nextBlank := !any(frames[i+1].Keys)
			if prevEqual && nextEqual && !immediateEqual { //} && !nextBlank {
				frames[i].start = true
				copy(frames[i].Keys, frames[i+offset].Keys)
			}
		}
	*/
}

type frameInfo struct {
	debug   image.Image
	ImgPath string
	Keys    []bool
	Time    time.Duration
}

func (f frameInfo) debugPrint() {
	fmt.Print(f.ImgPath)
	fmt.Print("  ")
	for _, k := range f.Keys {
		if k {
			fmt.Printf("X")
			continue
		}
		fmt.Printf("-")
	}
	fmt.Printf(" ")
}

func (f frameInfo) debugPrintLine() {
	f.debugPrint()
	fmt.Println(f.Time)
}

func detectFrameInfo(info *info, img string) (*frameInfo, error) {
	f, err := os.Open(img)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	//unmask(src.(*image.RGBA), info.mask, info.noteRegionStartY, info.noteRegionEndY)

	frameInfo := &frameInfo{debug: src, ImgPath: img}
	frameInfo.Keys = make([]bool, len(info.keys))
	for i, k := range info.keys {
		c1 := color.RGBAModel.Convert(src.At(k.center.X, k.center.Y)).(color.RGBA)
		c2 := color.RGBAModel.Convert(k.actualColor).(color.RGBA)
		dist := uint32(distu8(c1.R, c2.R)) + uint32(distu8(c1.G, c2.G)) + uint32(distu8(c1.B, c2.B))
		//fmt.Println(dist)
		if dist > 100 {
			frameInfo.Keys[i] = true
		}

		//dist := distu32(r1, r2) + distu32(g1, g2) + distu32(b1, b2) + distu32(a1, a2)
		//fmt.Println(distu32(r1, r2), distu32(g1, g2), distu32(b1, b2), distu32(a1, a2))
		//fmt.Println(src.At(k.center.X, k.center.Y), k.actualColor)
		//fmt.Println(dist)
		//if dist > 2000*4 {
		//	frameInfo.Keys[i] = true
		//} else {
		//}
	}
	return frameInfo, nil
}

/*
func hasLineAtTop(img string) (bool, error) {
	fmt.Println(img)
	f, err := os.Open(img)
	if err != nil {
		return false, err
	}
	defer f.Close()

	src, _, err := image.Decode(f)
	if err != nil {
		return false, err
	}
	//b := src.Bounds()
	out := make([]uint32, 15)
	for i := 0; i < len(out); i++ {
		r, g, b, _ := src.At(250, i).RGBA()
		out[i] = (r / 3) + (g / 3) + (b / 3)
	}

	a := (out[0] / 5) + (out[1] / 5) + (out[2] / 5) + (out[3] / 5) + (out[4] / 5)
	b := (out[5] / 5) + (out[6] / 5) + (out[7] / 5) + (out[8] / 5) + (out[9] / 5)
	c := (out[10] / 5) + (out[11] / 5) + (out[12] / 5) + (out[13] / 5) + (out[14] / 5)
	if a > b && c > b || a < b && c < b {
		return true, nil
	}
	return false, nil
}
*/
