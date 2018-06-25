# synmidi - Convert Synthesia YouTube videos to Midi files

This directory contains a tool capable of converting Synthesia YouTube videos into actual Midi files.

The process (and code) is quite complex, but ultimately it produces quite good results.

![image](https://user-images.githubusercontent.com/3173176/41830562-a2fb374a-77f6-11e8-8888-92c60a8b4a92.png)

## Dependencies

- Go
- ffmpeg

## Installation

```bash
go get -u github.com/slimsag/music/cmd/synmidi
```

## Example

```bash
# Create a directory to work in.
mkdir myvideo && cd myvideo

# Download the YouTube video (out.mp4):
synmidi download https://www.youtube.com/watch?v=lJJTBAqJoKM

# Convert the video to a series of PNG frames (frames/ dir):
synmidi frames

# Reduce the PNG frames to their vital components (frameInfos.json):
synmidi reduce

# IMPORTANT: You need to tell synmidi where one of the C keys is in the video!
# To do so, view the video and start counting the keys (white and black) from
# the very left of the screen. Start counting at 0. When you encounter a C key,
# write down that C key's octave (number next to it) and the key count.

# Finally, produce a midi file (file.mid), entering the C key you found and how
# many keys from the left it is (what you counted in the above step):
synmidi midi C2=7
```

## Note

Converting MP4 videos to midi files is quite a complex and error-prone process. This tool may not
work if your video contains e.g. start or end screens that are not Synthesia themselves. Open an issue if you need help converting a video. :)
