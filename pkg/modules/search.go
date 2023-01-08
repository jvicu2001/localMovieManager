package modules

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ryanbradynd05/go-tmdb"
)

type SearchStruct struct {
	ctx     context.Context
	tmdbAPI *tmdb.TMDb
}

func NewSearchStruct(tmdbAPI *tmdb.TMDb) *SearchStruct {
	return &SearchStruct{
		tmdbAPI: tmdbAPI,
	}
}

func (a *SearchStruct) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *SearchStruct) shutdown(ctx context.Context) {}

func (a *SearchStruct) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *SearchStruct) SearchMovies(query string, options SearchMoviesOptions) string {
	optionsMap := make(map[string]string)
	searchMoviesOptionsToMap(options, &optionsMap)
	queryResult, _ := a.tmdbAPI.SearchMovie(query, optionsMap)
	result, _ := tmdb.ToJSON(queryResult)
	return result
}

type SearchMoviesOptions struct {
	Language           *string      `json:"language,omitempty"`
	Page               *json.Number `json:"page,omitempty"`
	IncludeAdult       *bool        `json:"include_adult,omitempty"`
	Region             *string      `json:"region,omitempty"`
	Year               *json.Number `json:"year,omitempty"`
	PrimaryReleaseYear *json.Number `json:"primary_release_year,omitempty"`
}

func searchMoviesOptionsToMap(options SearchMoviesOptions, result *map[string]string) {
	temp := make(map[string]interface{})

	// Convert the struct to JSON
	jsonBytes, _ := json.Marshal(options)

	// Convert the JSON to a map
	err := json.Unmarshal(jsonBytes, &temp)
	if err != nil {
		panic(err)
	}

	for s, i := range temp {
		if i != nil {
			(*result)[s] = fmt.Sprintf("%v", i)
		}

	}
}

func (a *SearchStruct) SearchTV(query string, options SearchTVOptions) string {
	optionsMap := make(map[string]string)
	searchTVOptionsToMap(options, &optionsMap)
	queryResult, _ := a.tmdbAPI.SearchTv(query, optionsMap)
	result, _ := tmdb.ToJSON(queryResult)
	return result
}

type SearchTVOptions struct {
	Language         *string      `json:"language,omitempty"`
	Page             *json.Number `json:"page,omitempty"`
	IncludeAdult     *bool        `json:"include_adult,omitempty"`
	FirstAirDateYear *json.Number `json:"first_air_date_year,omitempty"`
}

func searchTVOptionsToMap(options SearchTVOptions, result *map[string]string) {
	temp := make(map[string]interface{})

	// Convert the struct to JSON
	jsonBytes, _ := json.Marshal(options)

	// Convert the JSON to a map
	err := json.Unmarshal(jsonBytes, &temp)
	if err != nil {
		panic(err)
	}

	for s, i := range temp {
		if i != nil {
			//fmt.Printf("%s %v\n", s, i)
			(*result)[s] = fmt.Sprintf("%v", i)
		}

	}
}

func (a *SearchStruct) SearchPeople(query string, options SearchPeopleOptions) string {
	optionsMap := make(map[string]string)
	searchPeopleOptionsToMap(options, &optionsMap)
	queryResult, _ := a.tmdbAPI.SearchPerson(query, optionsMap)
	result, _ := tmdb.ToJSON(queryResult)
	return result
}

type SearchPeopleOptions struct {
	Language     *string      `json:"language,omitempty"`
	Page         *json.Number `json:"page,omitempty"`
	IncludeAdult *bool        `json:"include_adult,omitempty"`
	Region       *string      `json:"region,omitempty"`
}

func searchPeopleOptionsToMap(options SearchPeopleOptions, result *map[string]string) {
	temp := make(map[string]interface{})

	// Convert the struct to JSON
	jsonBytes, _ := json.Marshal(options)

	// Convert the JSON to a map
	err := json.Unmarshal(jsonBytes, &temp)
	if err != nil {
		panic(err)
	}

	for s, i := range temp {
		if i != nil {
			(*result)[s] = fmt.Sprintf("%v", i)
		}

	}
}
