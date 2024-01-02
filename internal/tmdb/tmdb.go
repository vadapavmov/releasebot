package tmdb

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/vadapavmov/releasebot/internal/structs"
)

const base = "https://api.themoviedb.org/3"
const posterBase = "https://image.tmdb.org/t/p/original/"
const timeout = 30 * time.Second

type Engine struct {
	token  string
	client http.Client
}

func New(token string) *Engine {
	return &Engine{token, http.Client{Timeout: timeout}}
}

func (e *Engine) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+e.token)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (e *Engine) GetMovie(id string) (structs.Collection, error) {
	path := base + "/movie/" + id
	req, err := e.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("imdb api error: expected status code %d, recevied error code %d", http.StatusOK, resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	c := new(Collection)
	if err = c.unmarshal(body); err != nil {
		return nil, err
	}
	return c, nil
}

func (e *Engine) GetTv(id string) (structs.Collection, error) {
	path := base + "/tv/" + id
	req, err := e.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("imdb api error: expected status code %d, recevied error code %d", http.StatusOK, resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	c := new(Collection)
	if err = c.unmarshal(body); err != nil {
		return nil, err
	}
	return c, nil
}
