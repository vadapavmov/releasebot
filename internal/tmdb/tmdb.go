package tmdb

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const base = "https://api.themoviedb.org/3"
const posterBase = "https://image.tmdb.org/t/p/original/"
const timeout = 30 * time.Second

type Client struct {
	token  string
	client http.Client
}

func New(token string) *Client {
	return &Client{token, http.Client{Timeout: timeout}}
}

func (t *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+t.token)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (t *Client) GetMovie(id string) (*Collection, error) {
	path := base + "/movie/" + id
	req, err := t.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("tmdb api error: expected status code %d, recevied error code %d", http.StatusOK, resp.StatusCode)
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

func (t *Client) GetTv(id string) (*Collection, error) {
	path := base + "/tv/" + id
	req, err := t.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("tmdb api error: expected status code %d, recevied error code %d", http.StatusOK, resp.StatusCode)
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
