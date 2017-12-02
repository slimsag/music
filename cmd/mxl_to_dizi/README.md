# mxl_to_dizi - Converts Music XML to Dizi tablature

Converts Music XML to Dizi tablature. Note that the mxl file must already be transposed into the right key, e.g. by opening in MuseScore 2 and clicking `Notes -> Transpose` and selecting the Dizi key you're trying to play in.

Also note it works best on songs where only a single note is played at a time, i.e. no musical chords (because Dizi cannot play multiple notes at once).

# Installation

```Bash
go get -u github.com/slimsag/music/cmd/mxl_to_dizi
```

# Usage

```Bash
mxl_to_dizi <dizi_key> <file.mxl>
```
