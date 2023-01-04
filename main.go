package main

import (
	"context"
	"embed"
	"github.com/joho/godotenv"
	"github.com/ryanbradynd05/go-tmdb"
	"localMovieManager/pkg/modules"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Global TMDb object
	var tmbdAPI *tmdb.TMDb

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Create an instance of the app structure
	// Pass the TMDB object to everyone
	app := NewApp(tmbdAPI)
	search := modules.NewSearch(tmbdAPI)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "localMovieManager",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			search.SetContext(ctx)
		},
		Bind: []interface{}{
			app,
			search,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
