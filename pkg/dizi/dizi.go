// Package dizi defines generic informtion about Dizi flutes.
package dizi

// Fingering describes the finger positions that produce a note.
type Fingering struct {
	SimpleNotation string
	Fingers        string
	Altis          bool
}

// A map of Dizi key names (e.g. "C", "D", "E", etc.) to their respective note -> fingering position mapping.
var NoteMappings = map[string]map[string]Fingering{
	"D": map[string]Fingering{
		"A4":  {"5̣", "●●● ●●●", false},
		"B4":  {"6̣", "●●● ●●○", false},
		"C#5": {"7̣", "●●● ●○○", false},
		"D5":  {"1", "●●● ○○○", false},
		"E5":  {"2", "●●○ ○○○", false},
		"F#5": {"3", "●○○ ○○○", false},
		"G5":  {"4", "○●● ○○○ : ○●● ○●●", false},
		"G#5": {"_", "○○○ ○○○", false},
		"A5":  {"5", "○●● ●●●", false},
		"B5":  {"6", "●●● ●●○", false},
		"C#6": {"7", "●●● ●○○", false},
		"D6":  {"1̇", "●●● ○○○", false},
		"E6":  {"2̇", "●●○ ○○○", false},
		"F#6": {"3̇", "●○○ ○○○", false},
		"G6":  {"4̇", "○●● ●●○ : ◐○○ ○○○", true},
		"G#6": {"_", "○○○ ●●●", true},
		"A6":  {"5̇", "○●● ●●● : ●●○ ●●○", true},
	},
}
