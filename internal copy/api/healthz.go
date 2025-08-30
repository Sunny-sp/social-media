package main

import "net/http"

func (app *application) healthzCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Working Fine!\n"))
}
