package joke

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Joke struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

type JokeProvider interface {
	GetRandom(ctx context.Context) (*Joke, error)
}

type httpJokeProvider struct {
	client *http.Client
}

func NewHTTPJokeProvider() JokeProvider {
	return &httpJokeProvider{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (p *httpJokeProvider) GetRandom(ctx context.Context) (*Joke, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://official-joke-api.appspot.com/random_joke", nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request joke failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var j Joke
	if err := json.NewDecoder(resp.Body).Decode(&j); err != nil {
		return nil, fmt.Errorf("decode joke failed: %w", err)
	}

	return &j, nil
}
