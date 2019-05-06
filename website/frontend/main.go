package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

func main() {
	root := &Router{
		Handlers: []Handler{
			&StaticPageHandler{Dir: "pages"},
			&DiziPageHandler{},
		},
	}
	RerenderOnPopState = root
	vecty.RenderBody(root)
}

type Handler interface {
	Handle(path string) (c vecty.Component, ok bool)
}

type Router struct {
	Handlers []Handler
	vecty.Core
}

func (p *Router) Render() vecty.ComponentOrHTML {
	path := Query.Get("p")
	if path == "" {
		path = "index"
	}

	for _, h := range p.Handlers {
		c, ok := h.Handle(path)
		if !ok {
			continue
		}
		return c
	}
	return elem.Body(vecty.Text("page not found"))
}

/*
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
*/
