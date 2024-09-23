package app

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var static embed.FS

var AppDistFS fs.FS

func init() {
	AppDistFS, _ = fs.Sub(static, "dist")
}
