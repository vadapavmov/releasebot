package structs

import "io"

type SearchEngine interface {
	GetMovie(id string) (Collection, error)
	GetTv(id string) (Collection, error)
}

type Collection interface {
	Name() string
	Description() string
	Star() string
	Language() string
	GenreStr() string
	ReleaseTime() string
	Poster() (io.Reader, error)
}
