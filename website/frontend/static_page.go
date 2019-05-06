package main

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/slimsag/music/website"
	"gopkg.in/russross/blackfriday.v2"
)

type StaticPageHandler struct {
	Dir   string
	pages map[string]struct{}
}

func (h *StaticPageHandler) Handle(urlPath string) (vecty.Component, bool) {
	if h.pageExists(urlPath + ".md") {
		return &StaticPage{File: path.Join(h.Dir, urlPath+".md")}, true
	}
	return nil, false
}

func (p *StaticPageHandler) pageExists(path string) bool {
	if p.pages != nil {
		_, ok := p.pages[path]
		return ok
	}
	p.pages = map[string]struct{}{}
	_ = vfsutil.Walk(website.Assets, p.Dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		path = strings.TrimPrefix(path, p.Dir+"/")
		p.pages[path] = struct{}{}
		return nil
	})
	_, ok := p.pages[path]
	return ok
}

type StaticPage struct {
	vecty.Core
	File string
}

func (p *StaticPage) Render() vecty.ComponentOrHTML {
	f, err := website.Assets.Open(p.File)
	if err != nil {
		return elem.Body(vecty.Text(err.Error()))
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return elem.Body(vecty.Text(err.Error()))
	}

	email, _ := base64.StdEncoding.DecodeString("c3RlcGhlbi5ndXRla2Fuc3RAZ21haWwuY29t")
	data = bytes.Replace(data, []byte("{{.MyEmail}}"), email, -1)
	if bytes.Contains(data, []byte("{{.DiziFingeringsByKey}}")) {
		data = bytes.Replace(data, []byte("{{.DiziFingeringsByKey}}"), diziFingeringsByKey(), -1)
	}
	if bytes.Contains(data, []byte("{{.DiziFingeringsForKey \"D\"}}")) {
		data = bytes.Replace(data, []byte("{{.DiziFingeringsForKey \"D\"}}"), diziNoteMappingTable("D"), -1)
	}
	if bytes.Contains(data, []byte("{{.DiziMusicList}}")) {
		data = bytes.Replace(data, []byte("{{.DiziMusicList}}"), diziMusicList(), -1)
	}

	return elem.Body(
		vecty.Markup(
			vecty.Class("static-page"),
		),
		elem.Anchor(
			vecty.Markup(
				prop.Href("/"),
				vecty.Class("logo"),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("static-page__center"),
				vecty.UnsafeHTML(string(blackfriday.Run(data))),
			),
		),
	)
}
