package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/slimsag/music/pkg/dizi"
	"github.com/slimsag/music/pkg/mxl"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("usage: mxl_to_dizi <dizi_key> <file.xml>")
	}
	key := os.Args[1]
	file := os.Args[2]

	// Validate chosen Dizi key.
	noteMapping, ok := dizi.NoteMappings[key]
	if !ok {
		var keys []string
		for k := range dizi.NoteMappings {
			keys = append(keys, k)
		}
		log.Fatalf("unsupported Dizi key %q, valid keys are: %q", key, keys)
	}

	// Read and parse file.
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var d mxl.Doc
	if err := xml.Unmarshal(data, &d); err != nil {
		log.Fatal(err)
	}

	// Print general information about the piece.
	if d.Title != "" {
		fmt.Println("- Title:", d.Title)
	}
	if d.Composer != "" {
		fmt.Println("- Composer:", d.Composer)
	}
	if d.Date != "" {
		fmt.Println("- Date:", d.Date)
	}
	if d.Software != "" {
		fmt.Println("- Software:", d.Software)
	}
	if d.Source != "" {
		fmt.Println("- Source:", d.Source)
	}
	parts := 0
	measures := 0
	for _, p := range d.Parts {
		parts++
		for range p.Measures {
			measures++
		}
	}
	fmt.Println("- Parts:", parts)
	fmt.Println("- Measures:", measures)
	fmt.Println("")

	// Convert to tablature.
	fmt.Println("```")
	for _, p := range d.Parts {
		for _, m := range p.Measures {
			for _, n := range m.Notes {
				if n.Tie.Type == "stop" {
					continue
				}
				if n.Pitch.Step == "" {
					continue // TODO(slimsag): What is this for?
				}

				// Determine the final note (e.g. "C#7" or "C2")
				sharp := ""
				if n.Pitch.Accidental == 1 {
					sharp = "#"
				}
				// TODO(slimsag): We bump the octave up by one here because FL Studio 6 exports C5 as C4 in midi... wow.
				note := fmt.Sprintf("%v%v%v", n.Pitch.Step, sharp, n.Pitch.Octave+1)

				// Print the tab.
				tab, ok := noteMapping[note]
				if !ok {
					fmt.Printf("!!! note %s cannot be played on %s Dizi !!!\n", note, key)
					continue
				}
				if sharp == "" {
					note = note + " "
				}
				fmt.Printf("| %s | %s | %s\n", note, tab.SimpleNotation, tab.Fingers)
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n----------------------------------------\n")
	}
	fmt.Println("```")
}
