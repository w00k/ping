package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

// App export
type App struct {
	Router *mux.Router
}

func callPongService() string {
	var responseObj Response
	response, err := http.Get("http://localhost:8081/pong")

	if err != nil {
		log.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &responseObj)
	return responseObj.Message
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("pingHandler - request")
	var response map[string]interface{}
	str := callPongService()
	json.Unmarshal([]byte("{ \"send\": \"ping\", \"respond\": \""+str+"\" }"), &response)
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
