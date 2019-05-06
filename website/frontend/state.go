package main

import (
	"net/url"
	"strings"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
)

var Hash, Query url.Values

func currentURL() *url.URL {
	return &url.URL{
		Path:     CurrentPath(),
		RawQuery: Query.Encode(),
		Fragment: Hash.Encode(),
	}
}

func CurrentPath() string {
	return js.Global.Get("window").Get("location").Get("pathname").String()
}

func PushState(path string) {
	js.Global.Get("window").Get("history").Call("pushState", "", WithPath(path))
}

func ReplaceState(path string) {
	js.Global.Get("window").Get("history").Call("replaceState", "", WithPath(path))
}

func WithPath(path string) string {
	u := currentURL()
	u.Path = path
	return u.String()
}

func WithPathQuery(path string) string {
	q, _ := url.ParseQuery(Query.Encode())
	q.Set("p", path)
	return (&url.URL{
		Path:     CurrentPath(),
		RawQuery: q.Encode(),
		Fragment: Hash.Encode(),
	}).String()
}

var RerenderOnPopState vecty.Component

func init() {
	refreshHashQuery := func() {
		q := js.Global.Get("window").Get("location").Get("hash").String()
		q = strings.TrimPrefix(q, "#")
		Hash, _ = url.ParseQuery(q)

		q = js.Global.Get("window").Get("location").Get("search").String()
		q = strings.TrimPrefix(q, "?")
		Query, _ = url.ParseQuery(q)
	}
	refreshHashQuery()

	js.Global.Get("window").Set("onpopstate", js.MakeFunc(func(this *js.Object, args []*js.Object) interface{} {
		// Browser forward/back
		refreshHashQuery()

		if RerenderOnPopState != nil {
			vecty.Rerender(RerenderOnPopState)
		}
		return js.Undefined
	}))
}
