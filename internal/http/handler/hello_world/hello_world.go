package hello_world

import (
	"fmt"
	"net/http"
)

type Handler interface {
    Hello(w http.ResponseWriter, r *http.Request)
}

type handlerImpl struct{}

func New() Handler {
    return &handlerImpl{}
}

func (h *handlerImpl) Hello(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "World"
    }
    fmt.Fprintf(w, "Hello, %s!", name)
}
