package joke

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mclyashko/go-service/internal/provider/joke"
)

type Handler interface {
	Joke(w http.ResponseWriter, r *http.Request)
}

type handlerImpl struct {
	logger *log.Logger
	prov   joke.JokeProvider
}

func New(logger *log.Logger, prov joke.JokeProvider) Handler {
	return &handlerImpl{
		logger: logger,
		prov:   prov,
	}
}

func (h *handlerImpl) Joke(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	j, err := h.prov.GetRandom(ctx)
	if err != nil {
		h.logger.Println(fmt.Errorf("error getting joke: %w", err))
		http.Error(w, "failed to get joke", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(j); err != nil {
		h.logger.Println(fmt.Errorf("error writing response: %W", err))
	}
}
