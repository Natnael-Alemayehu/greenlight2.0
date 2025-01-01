package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/natnael-alemayehu/greenlight2.0/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create movie handler")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Persumed Inocent",
		Runtime:   102,
		Genres:    []string{"Thriller", "Drama"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusAccepted, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
