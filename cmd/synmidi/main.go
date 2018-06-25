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

	"github.com/gomidi/midi/midimessage/channel"
	"github.com/gomidi/midi/midimessage/meta"
	"github.com/gomidi/midi/smf"
	"github.com/gomidi/midi/smf/smfwriter"
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
				frameInfo.debugPrintLine()
				frameInfos = append(frameInfos, frameInfo)
			}

			if false {
				if err := os.MkdirAll(filepath.Join("debug"), 0777); err != nil {
					log.Fatal(err)
				}
				if err := writePNG(filepath.Join("debug", info.Name()), frameInfo.debug); err != nil {
					log.Fatal(err)
				}
			}
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
				f.debugPrintLine()
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

			downLeftKeys := make([]bool, len(frameInfos[0].Keys))
			downRightKeys := make([]bool, len(frameInfos[0].Keys))

			for frameIndex, keys := range determineKeyStates(frameInfos) {
				for keyIndex, key := range keys {
					midiNote := uint8(finalKeyOffset + keyIndex)
					velocity := uint8(90)
					leftCh := channel.Ch0
					rightCh := channel.Ch1
					ch := leftCh
					if !key.isLeftHand {
						ch = rightCh
					}

					if key.state == keyStatePressed || key.state == keyStatePressedAgain {
						if key.isLeftHand && downRightKeys[keyIndex] {
							// Must release the right key.
							_, err = wr.Write(rightCh.NoteOff(midiNote))
							if err != nil {
								log.Fatal(err)
							}
							wr.SetDelta(1)
						}
						if !key.isLeftHand && downLeftKeys[keyIndex] {
							// Must release the left key.
							_, err = wr.Write(leftCh.NoteOff(midiNote))
							if err != nil {
								log.Fatal(err)
							}
							wr.SetDelta(1)
						}
					}

					switch key.state {
					case keyStatePressed:
						if key.isLeftHand {
							downLeftKeys[keyIndex] = true
						} else {
							downRightKeys[keyIndex] = true
						}

						_, err := wr.Write(ch.NoteOn(midiNote, velocity))
						if err != nil {
							log.Fatal(err)
						}
					case keyStateContinued, keyStateNone:
						// Must send a message or else SetDelta takes no effect.
						_, err := wr.Write(ch.AfterTouch(255))
						if err != nil {
							log.Fatal(err)
						}
					case keyStatePressedAgain:
						_, err = wr.Write(ch.NoteOff(midiNote))
						if err != nil {
							log.Fatal(err)
						}
						wr.SetDelta(1)
						_, err := wr.Write(ch.NoteOn(midiNote, velocity))
						if err != nil {
							log.Fatal(err)
						}
					case keyStateReleased:
						if key.isLeftHand {
							downLeftKeys[keyIndex] = false
						} else {
							downRightKeys[keyIndex] = false
						}

						_, err = wr.Write(ch.NoteOff(midiNote))
						if err != nil {
							log.Fatal(err)
						}
					}
				}
				frameInfos[frameIndex].debugPrint()
				keys.debugPrint()
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

type frameInfo struct {
	debug         image.Image
	ImgPath       string
	Keys          []bool
	IsLeftHandKey []bool
	Time          time.Duration
}

func (f frameInfo) debugPrint() {
	fmt.Print(f.ImgPath)
	fmt.Print("  ")
	for i, k := range f.Keys {
		if k {
			if f.IsLeftHandKey[i] {
				fmt.Printf("L")
			} else {
				fmt.Printf("R")
			}
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
	frameInfo.IsLeftHandKey = make([]bool, len(info.keys))
	for i, k := range info.keys {
		c1 := color.RGBAModel.Convert(src.At(k.center.X, k.center.Y)).(color.RGBA)
		c2 := color.RGBAModel.Convert(k.actualColor).(color.RGBA)
		dist := uint32(distu8(c1.R, c2.R)) + uint32(distu8(c1.G, c2.G)) + uint32(distu8(c1.B, c2.B))
		if dist > 100 {
			frameInfo.Keys[i] = true

			// Left hand keys are usually this blue color.
			blueKey := color.RGBA{R: 116, G: 203, B: 197, A: c1.A}

			// Right hand keys are usually this greenish color (I am colorblind, it might be yellow-ish).
			greenishKey := color.RGBA{R: 233, G: 188, B: 78, A: c1.A}

			blueDist := uint32(distu8(c1.R, blueKey.R)) + uint32(distu8(c1.G, blueKey.G)) + uint32(distu8(c1.B, blueKey.B))
			greenishDist := uint32(distu8(c1.R, greenishKey.R)) + uint32(distu8(c1.G, greenishKey.G)) + uint32(distu8(c1.B, greenishKey.B))

			frameInfo.IsLeftHandKey[i] = blueDist < greenishDist
		}
	}
	return frameInfo, nil
}
