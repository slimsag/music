// +build dev

package website

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("./assets")
