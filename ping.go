package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App export
type App struct {
	Router *mux.Router
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("pingHandler - request")
	var response map[string]interface{}
	json.Unmarshal([]byte(`{ "message": "ping" }`), &response)
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func (app *App) initialiseRoutes() {
	log.Println("start")
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/ping", pingHandler)
}

func (app *App) run() {
	log.Println("ping service on 8080")
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

func main() {
	app := App{}
	app.initialiseRoutes()
	app.run()
}
