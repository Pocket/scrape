package server

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed assets/*
var assetsDir embed.FS

func assetsHandler() http.Handler {
	d, err := fs.Sub(assetsDir, "assets")
	if err != nil {
		panic(err)
	}
	// TODO: Disable directory listing
	return http.StripPrefix("/assets", http.FileServer(http.FS(d)))
}