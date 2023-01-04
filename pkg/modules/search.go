package modules

import (
	"context"
	"github.com/ryanbradynd05/go-tmdb"
)

type Search struct {
	ctx     context.Context
	tmdbAPI *tmdb.TMDb
}

func NewSearch(tmdbAPI *tmdb.TMDb) *Search {
	return &Search{
		tmdbAPI: tmdbAPI,
	}
}

func (a *Search) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *Search) shutdown(ctx context.Context) {}

func (a *Search) SetContext(ctx context.Context) {
	a.ctx = ctx
}
