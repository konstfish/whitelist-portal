package templates

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed static
var StaticFiles embed.FS

func GetStaticFiles() fs.FS {
	subFS, err := fs.Sub(StaticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	return subFS
}
