package movie

import (
	"context"
	movieDomain "github.com/null-like/movie-backend/domain/movie"
)

type Usecase interface {
	GetMovieInfo(c context.Context, movieId int) (movieDomain.Movie, error)
}
