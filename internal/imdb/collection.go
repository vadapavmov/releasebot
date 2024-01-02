package imdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Collection struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Plot   string   `json:"plot"`
	Image  string   `json:"image"`
	Genre  []string `json:"genre"`
	Rating struct {
		Count int     `json:"count"`
		Star  float64 `json:"star"`
	} `json:"rating"`
	SpokenLanguages []struct {
		Language string `json:"language"`
		ID       string `json:"ID"`
	} `json:"spokenLanguages"`
	ReleaseDetailed struct {
		Date  string `json:"date"`
		Day   int    `json:"day"`
		Month int    `json:"month"`
		Year  int    `json:"year"`
	} `json:"releaseDetailed"`
}

func (c *Collection) Name() string {
	return fmt.Sprintf("%s (%d)", c.Title, c.ReleaseDetailed.Year)
}

func (c *Collection) Description() string {
	return c.Plot
}

func (c *Collection) Star() string {
	return fmt.Sprintf("%.1f/10 (IMDB)", c.Rating.Star)
}

func (c *Collection) Language() string {
	if len(c.SpokenLanguages) < 0 {
		return ""
	}
	return c.SpokenLanguages[0].Language
}

func (c *Collection) ReleaseTime() string {
	releaseDetails := c.ReleaseDetailed
	return fmt.Sprintf("%d-%d-%d", releaseDetails.Year, releaseDetails.Month, releaseDetails.Day)
}

func (c *Collection) GenreStr() string {
	genres := ""
	for _, genre := range c.Genre {
		genres += "#" + genre + " "
	}
	return genres
}

func (c *Collection) Poster() (io.Reader, error) {
	resp, err := http.Get(c.Image)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("imdb poster api error: expected status code %d, recevied error code %d", http.StatusOK, resp.StatusCode)
	}
	return resp.Body, nil
}

func (c *Collection) unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}
