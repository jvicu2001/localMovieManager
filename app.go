package main

import (
	"context"
	"fmt"

	"github.com/ryanbradynd05/go-tmdb"
)

// App struct
type App struct {
	ctx     context.Context
	tmdbAPI *tmdb.TMDb
}

// NewApp creates a new App application struct
func NewApp(tmdbAPI *tmdb.TMDb) *App {
	return &App{
		tmdbAPI: tmdbAPI,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) TestAPI(query string) string {
	options := make(map[string]string)
	options["include_adult"] = "true"
	result, err := a.tmdbAPI.SearchMovie(query, options)
	if err != nil {
		return "Error: " + err.Error()
	}
	resultJson, err := tmdb.ToJSON(result.Results)
	return resultJson
}
