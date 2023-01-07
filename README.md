# localMovieManager

## About
This program allows the user to track the series and movies they watch, labeling them as "Planned", "Plan to watch", 
"Watching", "Completed", "On hold" and "Dropped". The project is inspired in [taiga.moe](https://github.com/erengy/taiga)

The movies and series data is provided by the The Movie Database API

### Pre-Alpha
This software is in the earliest stage of development it can be. Currently, it does nothing but test
the TMDb API

---
# Wails instructions
## Installing
Follow the instructions on the [Wails documentation](https://wails.io/docs/gettingstarted/installation/)

## Live Development

To run in live development mode, run `wails dev  -ldflags "-X main.APIKey={Your TMDb API Key}` in the 
project directory. This will run a Vite development server that will provide very fast hot reload of 
your frontend changes. If you want to develop in a browser and have access to your Go methods, there
is also a dev server that runs on http://localhost:34115. Connect to this in your browser, and you can
call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build -ldflags "-X main.APIKey={Your TMDb API Key}`.
