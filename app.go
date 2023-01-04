package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ryanbradynd05/go-tmdb"
	"io"
	"os"
)

// Config file struct
type configFileStruct struct {
	ApiKey string `json:"api_key"`
}

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

	// Initialize TMDb object
	if _, err := os.Stat("./data/config.json"); err == nil {
		configFile, err := os.Open("./data/config.json")
		if err != nil {
			panic(err)
		}

		configFileBytes, err := io.ReadAll(configFile)
		if err != nil {
			panic(err)
		}
		defer configFile.Close()

		var configfileJson configFileStruct
		err = json.Unmarshal(configFileBytes, &configfileJson)
		if err != nil {
			panic(err)
		}

		a.tmdbAPI = tmdb.Init(tmdb.Config{
			APIKey:   configfileJson.ApiKey,
			UseProxy: false,
			Proxies:  nil,
		})
	} else {
		configFile, err := os.Create("data/config.json")
		if err != nil {
			panic(err)
		}

		configFileJson := configFileStruct{ApiKey: ""}

		configFileBytes, err := json.Marshal(configFileJson)
		if err != nil {
			panic(err)
		}

		configFile.Write(configFileBytes)
		defer configFile.Close()
	}
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

func (a *App) UpdateAPIkey(key string) {
	configFile, err := os.Open("./data/config.json")
	if err != nil {
		panic(err)
	}

	configFileBytes, err := io.ReadAll(configFile)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	var configfileJson configFileStruct
	err = json.Unmarshal(configFileBytes, &configfileJson)
	if err != nil {
		panic(err)
	}

	configfileJson.ApiKey = key

	configFileBytes, err = json.Marshal(&configfileJson)

	configFile.Write(configFileBytes)

	a.tmdbAPI = tmdb.Init(tmdb.Config{
		APIKey:   key,
		UseProxy: false,
		Proxies:  nil,
	})
}
