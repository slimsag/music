// Package dizi defines generic informtion about Dizi flutes.
package dizi

// Fingering describes the finger positions that produce a note.
type Fingering struct {
	// The note this fingering produces in simple notation (often used in
	// Chinese sheet music.)
	SimpleNotation string

	// Fingers are the finger positions needed to produce the note. Leftmost
	// being closest to the mouthpiece.
	Fingers string

	// Alternative is an alternative finger position to produce the same note.
	// It is used based on preference and playing style. Sometimes switching to
	// an alternative fingering can be easier depending on the previous
	// fingering.
	Alternative string

	// Tier is the blowing tier needed to produce the note. One being regular
	// blowing, two being faster blowing to produce a higher pitch note, and
	// three being hard-to-reach altissimo octaves.
	Tier int
}

// Alt returns the alternative fingering, or if there is no Alternative
// the regular Fingers are returned.
func (f Fingering) Alt() string {
	if f.Alternative != "" {
		return f.Alternative
	}
	return f.Fingers
}

// Altis tells if this fingering produces an altissimo octave (a VERY hard to
// reach, extremely high pitch sound.)
func (f Fingering) Altis() bool {
	return f.Tier == 3
}

// NoteMappings is a map of Dizi key names (e.g. "C", "D", "E", etc.) to their
// respective note -> fingering position mapping.
var NoteMappings = map[string]map[string]Fingering{
	"G Bass": map[string]Fingering{
		"D4":  {"5̣", "●●●‧●●●", "", 1},
		"E4":  {"6̣", "●●●‧●●○", "", 1},
		"E#4": {"7̣", "●●●‧●○○", "", 1},
		"G4":  {"1", "●●●‧○○○", "", 1},
		"A4":  {"2", "●●○‧○○○", "", 1},
		"B4":  {"3", "●○○‧○○○", "", 1},
		"C5":  {"4", "○●●‧○○○", "○●●‧○●●", 1},
		"C#5": {"_", "○○○‧○○○", "", 1},

		"D5":  {"5", "○●●⁚●●●", "", 2},
		"E5":  {"6", "●●●⁚●●○", "", 2},
		"F#5": {"7", "●●●⁚●○○", "", 2},
		"G5":  {"1̇", "●●●⁚○○○", "", 2},
		"A5":  {"2̇", "●●○⁚○○○", "", 2},
		"B5":  {"3̇", "●○○⁚○○○", "", 2},

		"C6":  {"4̇", "○●●⁝●●○", "◐○○⁝○○○", 3},
		"C#6": {"_", "○○○⁝●●●", "", 3},
		"D6":  {"5̇", "○●●⁝●●●", "●●○⁝●●○", 3},
	},
	"C": map[string]Fingering{
		"G4":  {"5̣", "●●●‧●●●", "", 1},
		"A4":  {"6̣", "●●●‧●●○", "", 1},
		"B4":  {"7̣", "●●●‧●○○", "", 1},
		"C5":  {"1", "●●●‧○○○", "", 1},
		"D5":  {"2", "●●○‧○○○", "", 1},
		"E5":  {"3", "●○○‧○○○", "", 1},
		"F5":  {"4", "○●●‧○○○", "○●●‧○●●", 1},
		"F#5": {"_", "○○○‧○○○", "", 1},

		"G5": {"5", "○●●⁚●●●", "", 2},
		"A5": {"6", "●●●⁚●●○", "", 2},
		"B5": {"7", "●●●⁚●○○", "", 2},
		"C6": {"1̇", "●●●⁚○○○", "", 2},
		"D6": {"2̇", "●●○⁚○○○", "", 2},
		"E6": {"3̇", "●○○⁚○○○", "", 2},

		"F6":  {"4̇", "○●●⁝●●○", "◐○○⁝○○○", 3},
		"F#6": {"_", "○○○⁝●●●", "", 3},
		"G6":  {"5̇", "○●●⁝●●●", "●●○⁝●●○", 3},
	},
	"D": map[string]Fingering{
		"A4":  {"5̣", "●●●‧●●●", "", 1},
		"B4":  {"6̣", "●●●‧●●○", "", 1},
		"C#5": {"7̣", "●●●‧●○○", "", 1},
		"D5":  {"1", "●●●‧○○○", "", 1},
		"E5":  {"2", "●●○‧○○○", "", 1},
		"F#5": {"3", "●○○‧○○○", "", 1},
		"G5":  {"4", "○●●‧○○○", "○●●‧○●●", 1},
		"G#5": {"_", "○○○‧○○○", "", 1},

		"A5":  {"5", "○●●⁚●●●", "", 2},
		"B5":  {"6", "●●●⁚●●○", "", 2},
		"C#6": {"7", "●●●⁚●○○", "", 2},
		"D6":  {"1̇", "●●●⁚○○○", "", 2},
		"E6":  {"2̇", "●●○⁚○○○", "", 2},
		"F#6": {"3̇", "●○○⁚○○○", "", 2},

		"G6":  {"4̇", "○●●⁝●●○", "◐○○⁝○○○", 3},
		"G#6": {"_", "○○○⁝●●●", "", 3},
		"A6":  {"5̇", "○●●⁝●●●", "●●○⁝●●○", 3},
	},
	"E": map[string]Fingering{
		"B4":  {"5̣", "●●●‧●●●", "", 1},
		"C#5": {"6̣", "●●●‧●●○", "", 1},
		"D#5": {"7̣", "●●●‧●○○", "", 1},
		"E5":  {"1", "●●●‧○○○", "", 1},
		"F#5": {"2", "●●○‧○○○", "", 1},
		"G#5": {"3", "●○○‧○○○", "", 1},
		"A5":  {"4", "○●●‧○○○", "○●●‧○●●", 1},
		"A#5": {"_", "○○○‧○○○", "", 1},

		"B5":  {"5", "○●●⁚●●●", "", 2},
		"C#6": {"6", "●●●⁚●●○", "", 2},
		"D#6": {"7", "●●●⁚●○○", "", 2},
		"E6":  {"1̇", "●●●⁚○○○", "", 2},
		"F#6": {"2̇", "●●○⁚○○○", "", 2},
		"G#6": {"3̇", "●○○⁚○○○", "", 2},

		"A6":  {"4̇", "○●●⁝●●○", "◐○○⁝○○○", 3},
		"A#6": {"_", "○○○⁝●●●", "", 3},
		"B6":  {"5̇", "○●●⁝●●●", "●●○⁝●●○", 3},
	},
	"F": map[string]Fingering{
		"C5":  {"5̣", "●●●‧●●●", "", 1},
		"D5":  {"6̣", "●●●‧●●○", "", 1},
		"E5":  {"7̣", "●●●‧●○○", "", 1},
		"F5":  {"1", "●●●‧○○○", "", 1},
		"G5":  {"2", "●●○‧○○○", "", 1},
		"A5":  {"3", "●○○‧○○○", "", 1},
		"A#5": {"4", "○●●‧○○○", "○●●‧○●●", 1},
		"B5":  {"_", "○○○‧○○○", "", 1},

		"C6": {"5", "○●●⁚●●●", "", 2},
		"D6": {"6", "●●●⁚●●○", "", 2},
		"E6": {"7", "●●●⁚●○○", "", 2},
		"F6": {"1̇", "●●●⁚○○○", "", 2},
		"G6": {"2̇", "●●○⁚○○○", "", 2},
		"A6": {"3̇", "●○○⁚○○○", "", 2},

		"A#6": {"4̇", "○●●⁝●●○", "◐○○⁝○○○", 3},
		"B6":  {"_", "○○○⁝●●●", "", 3},
		"C7":  {"5̇", "○●●⁝●●●", "●●○⁝●●○", 3},
	},
	"G": map[string]Fingering{
		"D5":  {"5̣", "●●●‧●●●", "", 1},
		"E5":  {"6̣", "●●●‧●●○", "", 1},
		"E#5": {"7̣", "●●●‧●○○", "", 1},
		"G5":  {"1", "●●●‧○○○", "", 1},
		"A5":  {"2", "●●○‧○○○", "", 1},
		"B5":  {"3", "●○○‧○○○", "", 1},
		"C6":  {"4", "○●●‧○○○", "○●●‧○●●", 1},
		"C#6": {"_", "○○○‧○○○", "", 1},

		"D6":  {"5", "○●●⁚●●●", "", 2},
		"E6":  {"6", "●●●⁚●●○", "", 2},
		"F#6": {"7", "●●●⁚●○○", "", 2},
		"G6":  {"1̇", "●●●⁚○○○", "", 2},
		"A6":  {"2̇", "●●○⁚○○○", "", 2},
		"B6":  {"3̇", "●○○⁚○○○", "", 2},

		"C7":  {"4̇", "○●●⁝●●○", "◐○○⁝○○○", 3},
		"C#7": {"_", "○○○⁝●●●", "", 3},
		"D7":  {"5̇", "○●●⁝●●●", "●●○⁝●●○", 3},
	},
}
