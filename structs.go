package main

type Data struct {
	Adult               bool        `json:"adult"`
	BackdropPath        string      `json:"backdrop_path"`
	BelongsToCollection interface{} `json:"belongs_to_collection"`
	Budget              int         `json:"budget"`
	Genres              []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"genres"`
	Homepage            string        `json:"homepage"`
	ID                  int           `json:"id"`
	ImdbID              interface{}   `json:"imdb_id"`
	OriginalLanguage    string        `json:"original_language"`
	OriginalTitle       string        `json:"original_title"`
	Overview            string        `json:"overview"`
	Popularity          float64       `json:"popularity"`
	PosterPath          string        `json:"poster_path"`
	ProductionCompanies []interface{} `json:"production_companies"`
	ProductionCountries []interface{} `json:"production_countries"`
	ReleaseDate         string        `json:"release_date"`
	Revenue             int           `json:"revenue"`
	Runtime             int           `json:"runtime"`
	SpokenLanguages     []struct {
		EnglishName string `json:"english_name"`
		Iso6391     string `json:"iso_639_1"`
		Name        string `json:"name"`
	} `json:"spoken_languages"`
	Status      string  `json:"status"`
	Tagline     string  `json:"tagline"`
	Title       string  `json:"title"`
	Video       bool    `json:"video"`
	VoteAverage float64 `json:"vote_average"`
	VoteCount   int     `json:"vote_count"`
}
