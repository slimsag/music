package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type frameRate struct {
	numerator, denominator int
}

func (f frameRate) FPS() float64 {
	return float64(f.numerator) / float64(f.denominator)
}

// FrameSeconds returns the amount of time a single frame takes, in seconds.
func (f frameRate) FrameSeconds() float64 {
	return 1.0 / f.FPS()
}

func detectFrameRate(videoFile string) (frameRate, error) {
	cmd := exec.Command("ffprobe", "-v", "0", "-of", "csv=p=0", "-select_streams", "0", "-show_entries", "stream=r_frame_rate", videoFile)
	cmd.Stderr = os.Stderr
	var buf bytes.Buffer
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		return frameRate{}, err
	}
	split := strings.Split(strings.TrimSpace(buf.String()), "/")
	numerator, _ := strconv.Atoi(split[0])
	denominator, _ := strconv.Atoi(split[1])
	return frameRate{numerator: numerator, denominator: denominator}, nil
}
