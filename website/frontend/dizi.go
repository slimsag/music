package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"path"
	"strings"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/slimsag/music/pkg/dizi"
	"github.com/slimsag/music/website"
)

func diziMusicList() []byte {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "<ul>")
	dir, err := website.Assets.Open("songs")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	infos, err := dir.Readdir(-1)
	if err != nil {
		panic(err)
	}
	for _, i := range infos {
		fpath := path.Join("songs", i.Name())
		info := readSongInfo(fpath)
		u := WithPathQuery("dizi/song/" + i.Name())
		fmt.Fprintf(&buf, `<li><a href="%s">%s - %s (%s key)</a></li>`, u, info.Title, info.Author, info.Key)
	}
	fmt.Fprintf(&buf, "</ul>")
	return buf.Bytes()
}

type songInfo struct {
	Title, Author, Key string
}

func readSongInfo(path string) *songInfo {
	data, err := vfsutil.ReadFile(website.Assets, path)
	if err != nil {
		panic(err)
	}
	info := &songInfo{}
	lines := bytes.Split(data, []byte("\n"))
	for _, l := range lines {
		line := strings.TrimSpace(string(l))
		if line == "" || !strings.Contains(line, ":") {
			return info
		}
		pair := strings.Split(line, ":")
		switch strings.TrimSpace(pair[0]) {
		case "title":
			info.Title = strings.TrimSpace(pair[1])
		case "author":
			info.Author = strings.TrimSpace(pair[1])
		case "key":
			info.Key = strings.TrimSpace(pair[1])
		}
	}
	return info
}

func diziFingeringsByKey() []byte {
	var buf bytes.Buffer

	order := []string{
		"D",
		"C",
		"E",
		"F",
		"G",
		"G Bass",
	}
	fmt.Fprintf(&buf, "<div>\n")
	for _, key := range order {
		fmt.Fprintf(&buf, "<details><summary>%s key dizi</summary>\n", key)
		rng := dizi.Ranges[key]
		fmt.Fprintf(&buf, "<p><strong>Range:</strong> %s - %s</p>\n", rng[0], rng[1])
		writeDiziNoteMappingTable(&buf, key)
		fmt.Fprintf(&buf, "</details>\n")
	}
	fmt.Fprintf(&buf, "</div>\n")
	return buf.Bytes()
}

func diziNoteMappingTable(key string) []byte {
	var buf bytes.Buffer
	writeDiziNoteMappingTable(&buf, key)
	return buf.Bytes()
}

func writeDiziNoteMappingTable(w io.Writer, key string) {
	fmt.Fprintf(w, "<table>\n")
	fmt.Fprintf(w, "<thead>\n")
	fmt.Fprintf(w, "<tr><th>Note</th><th>Simple Notation</th><th>Fingering</th><th>Blowing level</th></tr>\n")
	fmt.Fprintf(w, "</thead>\n")

	fmt.Fprintf(w, "<tbody>\n")
	tier := 1
	for note, f := range dizi.NoteMappings[key] {
		class := ""
		if tier != f.Tier {
			tier = f.Tier
			class = ` class="separated"`
		}
		fmt.Fprintf(w, "<tr%s><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>\n", class, note, f.SimpleNotation, f.HTML(), f.TierString())
	}
	fmt.Fprintf(w, "</tbody>\n")
	fmt.Fprintf(w, "</table>\n")
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

type DiziPageHandler struct {
	songs map[string]struct{}
}

func (d *DiziPageHandler) Handle(urlPath string) (vecty.Component, bool) {
	if d.songs == nil {
		d.songs = make(map[string]struct{})
		dir, err := website.Assets.Open("songs")
		if err != nil {
			panic(err)
		}
		defer dir.Close()
		infos, err := dir.Readdir(-1)
		if err != nil {
			panic(err)
		}
		for _, i := range infos {
			d.songs[i.Name()] = struct{}{}
		}
	}
	song := strings.TrimPrefix(urlPath, "dizi/song/")
	_, ok := d.songs[song]
	if ok {
		return &diziMusicPage{song: song}, true
	}
	return nil, false
}

type diziMusicPage struct {
	vecty.Core
	song string
}

func (d *diziMusicPage) Render() vecty.ComponentOrHTML {
	return elem.Body(
		d.renderSong(),
	)
}

func (p *diziMusicPage) renderSong() vecty.ComponentOrHTML {
	f, err := website.Assets.Open(path.Join("songs/", p.song))
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
			textLine(fmt.Sprintf("Title: %s", title)),
			textLine(fmt.Sprintf("Author: %s", author)),
			textLine(fmt.Sprintf("Key: %s", key)),
		),
		p.sections(song),
	)
}

func textLine(text string) vecty.ComponentOrHTML {
	return elem.Paragraph(
		vecty.Markup(
			vecty.Style("margin", "0"),
		),
		vecty.Text(text),
	)
}

func (p *diziMusicPage) sections(song *song) vecty.ComponentOrHTML {
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

func (p *diziMusicPage) section(song *song, section section) vecty.ComponentOrHTML {
	noteMapping := dizi.NoteMappings[song.get("key", "")]

	var notes vecty.List
	for _, note := range section {
		finger := noteMapping[note]
		//label := fmt.Sprintf("%s%s | %s | %s", note, sharpSpacing, finger.SimpleNotation, finger.Fingers)
		//label := fmt.Sprintf("<span>%s%s | %s</span>", note, sharpSpacing, finger.HTML())
		notes = append(notes, elem.Div(
			vecty.Markup(
				vecty.Style("display", "flex"),
				//vecty.Style("justify-content", "space-between"),
			),
			elem.Span(
				vecty.Markup(
					vecty.Style("display", "inline-block"),
					vecty.Style("border-right", "1px solid black"),
					vecty.Style("width", "2rem"),
				),
				vecty.Text(note),
			),
			elem.Span(
				vecty.Markup(
					vecty.Style("display", "inline-block"),
					vecty.Style("margin-left", "0.25rem"),
					vecty.Markup(
						vecty.UnsafeHTML(finger.HTML()),
					),
				),
			),
		))
	}
	return notes
}
