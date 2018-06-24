// Package mxl exposes data types for parsing Music XML files, e.g. from MuseScore 2.
package mxl

import "encoding/xml"

const (
	// Header is the XML header that should be prepended to the output of
	// xml.Marshal when encoding a Document.
	Header = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE score-partwise PUBLIC "-//Recordare//DTD MusicXML 3.0 Partwise//EN" "http://www.musicxml.org/dtds/partwise.dtd">
`
)

// Document holds all data for a music xml file
type Document struct {
	XMLName        xml.Name        `xml:"score-partwise"`
	Identification *Identification `xml:"identification,omitempty"`
	Defaults       *Defaults       `xml:"defaults,omitempty"`
	PartList       *PartList       `xml:"part-list,omitempty"`
	Parts          []Part          `xml:"part"`
}

// Identification holds all of the identification information for the document.
type Identification struct {
	Composer string    `xml:"creator,omitempty"`
	Encoding *Encoding `xml:"encoding,omitempty"`
	Rights   string    `xml:"rights,omitempty"`
	Source   string    `xml:"source,omitempty"`
	Title    string    `xml:"movement-title,omitempty"`
}

// Encoding holds encoding information for the Identification.
type Encoding struct {
	Software     string    `xml:"software,omitempty"`
	EncodingDate string    `xml:"encoding-date,omitempty"`
	Supports     []Support `xml:"supports,omitempty"`
}

// Support holds support data for the Encoding.
type Support struct {
	Element   string `xml:"element,attr,omitempty"`
	Attribute string `xml:"attribute,attr,omitempty"`
	Type      string `xml:"type,attr,omitempty"`
	Value     string `xml:"value,attr,omitempty"`
}

// Defaults holds default information for the Document.
type Defaults struct {
	Scaling    *Scaling    `xml:"scaling,omitempty"`
	PageLayout *PageLayout `xml:"page-layout,omitempty"`
	WordFont   *Font       `xml:"word-font,omitempty"`
	LyricFont  *Font       `xml:"lyric-font,omitempty"`
}

// Scaling holds scaling information for the Defaults.
type Scaling struct {
	Millimeters float64 `xml:"millimeters,omitempty"`
	Tenths      int     `xml:"tenths,omitempty"`
}

// PageLayout holds page layout information for the Defaults.
type PageLayout struct {
	PageHeight  float64       `xml:"page-height,omitempty"`
	PageWidth   float64       `xml:"page-width,omitempty"`
	PageMargins []PageMargins `xml:"page-margins,omitempty"`
}

// PageMargins holds page margins information for the PageLayout.
type PageMargins struct {
	Type         string  `xml:"type,attr,omitempty"`
	LeftMargin   float64 `xml:"left-margin,omitempty"`
	RightMargin  float64 `xml:"right-margin,omitempty"`
	TopMargin    float64 `xml:"top-margin,omitempty"`
	BottomMargin float64 `xml:"bottom-margin,omitempty"`
}

// Font holds font information for the Defaults.
type Font struct {
	FontFamily string `xml:"font-family,attr,omitempty"`
	FontSize   int    `xml:"font-size,attr,omitempty"`
}

// PartList holds part list information for the Document.
type PartList struct {
	ScorePart *ScorePart `xml:"score-part,omitempty"`
}

// ScorePart holds score part information for the PartList.
type ScorePart struct {
	ID               string           `xml:"id,attr,omitempty"`
	PartName         *String          `xml:"part-name,omitempty"`
	PartAbbreviation *String          `xml:"part-abbreviation,omitempty"`
	ScoreInstrument  *ScoreInstrument `xml:"score-instrument,omitempty"`
	MidiDevice       *MidiDevice      `xml:"midi-device,omitempty"`
	MidiInstrument   *MidiInstrument  `xml:"midi-instrument,omitempty"`
}

// ScoreInstrument holds score instrument information for the ScorePart.
type ScoreInstrument struct {
	ID             string  `xml:"id,attr,omitempty"`
	InstrumentName *String `xml:"instrument-name,omitempty"`
}

// MidiDevice holds midi device information for the ScorePart.
type MidiDevice struct {
	ID   string `xml:"id,attr,omitempty"`
	Port int    `xml:"port,attr,omitempty"`
}

// MidiInstrument holds midi instrument information for the ScorePart.
type MidiInstrument struct {
	ID          string   `xml:"id,attr,omitempty"`
	MidiChannel *Int     `xml:"midi-channel,omitempty"`
	MidiProgram *Int     `xml:"midi-program,omitempty"`
	Volume      *Float64 `xml:"volume,omitempty"`
	Pan         *Float64 `xml:"pan,omitempty"`
}

// String represents any tag whose value is just a string. e.g. <foobar>somestring</foobar>
type String struct {
	Value string `xml:",chardata"`
}

// Int represents any tag whose value is just a int. e.g. <foobar>1337</foobar>
type Int struct {
	Value int `xml:",chardata"`
}

// Float64 represents any tag whose value is just a float64. e.g. <foobar>13.37</foobar>
type Float64 struct {
	Value float64 `xml:",chardata"`
}

// Part represents a part in a piece of music for the Document.
type Part struct {
	ID       string    `xml:"id,attr,omitempty"`
	Measures []Measure `xml:"measure,omitempty"`
}

// Measure represents a measure in a piece of music for a Part.
type Measure struct {
	Number     int           `xml:"number,attr,omitempty"`
	Width      float64       `xml:"width,attr,omitempty"`
	Print      *Print        `xml:"print,omitempty"`
	Direction  *Direction    `xml:"direction,omitempty"`
	Notes      []Note        `xml:"note,omitempty"`
	Attributes []*Attributes `xml:"attributes,omitempty"`
}

// Print represents printing information for a Measure.
type Print struct {
	NewPage      string       `xml:"new-page,attr,omitempty"`
	NewSystem    string       `xml:"new-system,attr,omitempty"`
	SystemLayout SystemLayout `xml:"system-layout,omitempty"`
}

// SystemLayout represents system layout information for a Print.
type SystemLayout struct {
	SystemMargins     *SystemMargins `xml:"system-margins,omitempty"`
	TopSystemDistance *Float64       `xml:"top-system-distance,omitempty"`
	SystemDistance    *Float64       `xml:"system-distance,omitempty"`
}

// SystemMargins represents system margins information for a SystemLayout.
type SystemMargins struct {
	LeftMargin  *Float64 `xml:"left-margin,omitempty"`
	RightMargin *Float64 `xml:"right-margin,omitempty"`
}

// Attributes represents attributes for a Measure.
type Attributes struct {
	Divisions *Int  `xml:"divisions,omitempty"`
	Key       *Key  `xml:"key,omitempty"`
	Time      *Time `xml:"time,omitempty"`
	Clef      *Clef `xml:"clef,omitempty"`
}

// Key represents a key signature change for a Measure's Attributes.
type Key struct {
	Fifths int    `xml:"fifths,omitempty"`
	Mode   string `xml:"mode,omitempty"`
}

// Time represents a time signature change
type Time struct {
	Beats    int `xml:"beats"`
	BeatType int `xml:"beat-type"`
}

// Clef represents a clef change
type Clef struct {
	Sign string `xml:"sign"`
	Line int    `xml:"line"`
}

// Direction holds direction informatino for a Measure.
type Direction struct {
	Placement     string         `xml:"placement,attr,omitempty"`
	DirectionType *DirectionType `xml:"direction-type,omitempty"`
	Sound         *Sound         `xml:"sound,omitempty"`
}

// DirectionType represents the direction type for a Direction.
type DirectionType struct {
	Metronome *Metronome `xml:"metronome,omitempty"`
}

// Metronome represents a metronome for a DirectionType.
type Metronome struct {
	Parentheses string  `xml:"parentheses,attr,omitempty"`
	BeatUnit    *String `xml:"beat-unit,omitempty"`
	PerMinute   *Int    `xml:"per-minute,omitempty"`
}

// Sound represents a Sound for a Direction.
type Sound struct {
	Tempo int `xml:"tempo,attr,omitempty"`
}

// Note represents a note in a Measure.
type Note struct {
	DefaultX         float64           `xml:"default-x,attr,omitempty"`
	DefaultY         float64           `xml:"default-y,attr,omitempty"`
	Dynamics         float64           `xml:"dynamics,attr,omitempty"`
	Pitch            *Pitch            `xml:"pitch,omitempty"`
	Duration         int               `xml:"duration,omitempty"`
	Voice            int               `xml:"voice,omitempty"`
	Type             string            `xml:"type,omitempty"`
	TimeModification *TimeModification `xml:"time-modification,omitempty"`
	Stem             string            `xml:"stem,omitempty"`
	Beam             []*Beam           `xml:"beam,omitempty"`
	Notations        []*Notation       `xml:"notations,omitempty"`
	Rest             *xml.Name         `xml:"rest,omitempty"`
	Dot              *xml.Name         `xml:"dot,omitempty"`
	Chord            *xml.Name         `xml:"chord,omitempty"`
	Tie              []*Tie            `xml:"tie,omitempty"`
}

// Pitch represents the pitch of a Note.
type Pitch struct {
	Step       string `xml:"step,omitempty"`
	Accidental int8   `xml:"alter,omitempty"` // TODO: rename to Alter and document that it is Accidental?
	Octave     int    `xml:"octave,omitempty"`
}

// TimeModification represents time modification for a Note.
type TimeModification struct {
	ActualNotes *Int    `xml:"actual-notes,omitempty"`
	NormalNotes *Int    `xml:"normal-notes,omitempty"`
	NormalType  *String `xml:"normal-type,omitempty"`
}

// Beam represents the Beam of a Note.
type Beam struct {
	Number int    `xml:"number,attr,omitempty"`
	Value  string `xml:",chardata"`
}

// Notation represents notations for a Note.
type Notation struct {
	Tied   []*Tied `xml:"tied,omitempty"`
	Tuplet *Tuplet `xml:"tuplet,omitempty"`
}

// Tuplet represents a tuplet for a Notation.
type Tuplet struct {
	Type    string `xml:"type,attr,omitempty"`
	Bracket string `xml:"bracket,attr,omitempty"`
}

// Tie represents whether or not a Note is tied.
type Tie struct {
	Type string `xml:"type,attr,omitempty"`
}

// Tied represents whether or not a Notations is tied.
type Tied struct {
	Type string `xml:"type,attr,omitempty"`
}
