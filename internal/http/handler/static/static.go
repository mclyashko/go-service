package static

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed files
var content embed.FS

type Handler interface {
	Files(w http.ResponseWriter, r *http.Request)
}

type handlerImpls struct {
	fs http.Handler
}

func New() Handler {
	fsys, _ := fs.Sub(content, "files")
	return &handlerImpls{
		fs: http.FileServer(http.FS(fsys)),
	}
}

func (h *handlerImpls) Files(w http.ResponseWriter, r *http.Request) {
	h.fs.ServeHTTP(w, r)
}
