package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Collection struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	OriginalName     string `json:"original_name"`
	OriginalTitle    string `json:"original_title"`
	OriginalLanguage string `json:"original_language"`
	Overview         string `json:"overview"`
	PosterPath       string `json:"poster_path"`
	ReleaseDate      string `json:"release_date"`
	FirstAirDate     string `json:"first_air_date"`
	Genres           []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
}

func (c *Collection) Poster() (io.Reader, error) {
	resp, err := http.Get(posterBase + c.PosterPath)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("tmdb poster api error: expected status code %d, recevied error code %d", http.StatusOK, resp.StatusCode)
	}
	return resp.Body, nil
}

func (c *Collection) GenreStr() string {
	genres := ""
	for _, genre := range c.Genres {
		genres += "#" + genre.Name
	}
	return genres
}

func (c *Collection) Name() string {
	name := c.Title
	if name == "" {
		name = c.OriginalTitle
	}
	if name == "" {
		name = c.OriginalName
	}
	return name
}

func (c *Collection) Description() string {
	return c.Overview
}

func (c *Collection) Language() string {
	return c.OriginalLanguage
}

func (c *Collection) ReleaseTime() string {
	if c.ReleaseDate != "" {
		return c.ReleaseDate
	}
	return c.FirstAirDate
}

func (c *Collection) unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}
