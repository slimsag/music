package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/slimsag/music/pkg/dizi"
	"github.com/slimsag/music/website"
)

func main() {
	vecty.SetTitle("Dizi")
	vecty.RenderBody(&PageView{})
}

type PageView struct {
	vecty.Core
	chosenSong string
}

func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Markup(
			vecty.Style("height", "100%"),
		),
		vecty.If(p.chosenSong == "", p.songList()),
		vecty.If(p.chosenSong != "", p.song(p.chosenSong)),
	)
}

func (p *PageView) songList() vecty.ComponentOrHTML {
	dir, err := website.Assets.Open("songs")
	if err != nil {
		return elem.Div(vecty.Text(err.Error()))
	}
	defer dir.Close()
	infos, err := dir.Readdir(-1)
	if err != nil {
		return elem.Div(vecty.Text(err.Error()))
	}
	var items vecty.List
	for _, i := range infos {
		i := i
		items = append(items, elem.Div(
			elem.Anchor(
				vecty.Markup(
					vecty.Property("href", ""),
					event.Click(func(e *vecty.Event) {
						p.chosenSong = "songs/" + i.Name()
						vecty.Rerender(p)
						e.Call("preventDefault")
					}),
				),
				vecty.Text(i.Name()),
			),
		))
	}
	return items
}

func (p *PageView) song(path string) vecty.ComponentOrHTML {
	f, err := website.Assets.Open(path)
	if err != nil {
		return elem.Div(vecty.Text(err.Error()))
	}
	defer f.Close()
	song, err := parseSong(f)
	if err != nil {
		return elem.Div(vecty.Text(err.Error()))
	}

	title := song.get("title", "no title")
	author := song.get("author", "no author")
	key := song.get("key", "no key")
	if _, ok := dizi.NoteMappings[key]; !ok {
		return elem.Div(vecty.Text(fmt.Sprintf("no known dizi note mapping for song key %q", key)))
	}
	return elem.Div(
		vecty.Markup(
			vecty.Style("display", "flex"),
			vecty.Style("height", "100%"),
			vecty.Style("flex-flow", "wrap"),
			vecty.Style("flex-direction", "column"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Style("border", "1px solid black"),
				vecty.Style("padding", "5px"),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Property("href", ""),
					event.Click(func(e *vecty.Event) {
						p.chosenSong = ""
						vecty.Rerender(p)
						e.Call("preventDefault")
					}),
				),
				vecty.Text("back"),
			),
			elem.Preformatted(
				vecty.Markup(
					vecty.Style("margin", "0"),
				),
				vecty.Text(fmt.Sprintf("title: %s\nauthor: %s\nkey: %s", title, author, key)),
			),
		),
		p.sections(song),
	)
}

func (p *PageView) sections(song *song) vecty.ComponentOrHTML {
	var sections vecty.List
	for _, section := range song.sections {
		sections = append(sections, elem.Div(
			vecty.Markup(
				vecty.Style("border", "1px solid black"),
				vecty.Style("padding", "5px"),
			),
			p.section(song, section),
		))
	}
	return sections
}

func (p *PageView) section(song *song, section section) vecty.ComponentOrHTML {
	noteMapping := dizi.NoteMappings[song.get("key", "")]

	var notes vecty.List
	for _, note := range section {
		finger := noteMapping[note]
		sharpSpacing := " "
		if strings.Contains(note, "#") {
			sharpSpacing = ""
		}
		//label := fmt.Sprintf("%s%s | %s | %s", note, sharpSpacing, finger.SimpleNotation, finger.Fingers)
		label := fmt.Sprintf("%s%s | %s", note, sharpSpacing, finger.Accented())
		notes = append(notes, elem.Preformatted(
			vecty.Markup(
				vecty.Style("margin", "0"),
			),
			vecty.Text(label),
		))
	}
	return notes
}

type section []string // notes

type song struct {
	metadata map[string]string
	sections []section
}

func (s *song) get(key, defaultValue string) string {
	if v, ok := s.metadata[key]; ok {
		return v
	}
	return defaultValue
}

func parseSong(r io.Reader) (*song, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	s := &song{
		metadata: map[string]string{},
	}
	var pendingSection section
	readingMetadata := true
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			readingMetadata = false
			if len(pendingSection) > 0 {
				s.sections = append(s.sections, pendingSection)
				pendingSection = nil
			}
			continue
		}
		if readingMetadata {
			split := strings.Split(line, ":")
			if len(split) == 2 {
				s.metadata[strings.TrimSpace(split[0])] = strings.TrimSpace(split[1])
			} else {
				fmt.Printf("ignoring unsupported metadata key:value pair with multiple colons %q\n", line)
			}
			continue
		}
		line = strings.TrimPrefix(line, ".") // TODO(slimsag): handle taps
		pendingSection = append(pendingSection, line)
	}
	return s, nil
}
