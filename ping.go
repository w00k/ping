package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// ConstantApp ::: estructura que almacena variables de entorno
type ConstantApp struct {
	EndpointPong string
}

// Response ::: estructura de respuesta del servicio pong
type Response struct {
	Message string `json:"message"`
}

// constantApp ::: inicialización de variable
var constantApp ConstantApp

// función que se ejecuta al inicio, setea la variable de entorno en el struct constantApp
func init() {
	constantApp = ConstantApp{
		EndpointPong: os.Getenv("URL_PONG"),
	}
}

// App export
type App struct {
	Router *mux.Router
}

// callPongService ::: llamada al servicio pong
// antes de ejecutar el servicio es necesario setear la variable de entorno URL_PONG
// en windows con cmd se hace de esta forma
// set URL_PONG=http://localhost:8081/pong
func callPongService() string {
	var responseObj Response
	response, err := http.Get(constantApp.EndpointPong)
	// set URL_PONG=http://localhost:8081/pong
	//response, err := http.Get("http://localhost:8081/pong")

	if err != nil {
		log.Print(err.Error())
		return "error al llamar servicio pong"
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return "error al leer el mensaje de servicio pong"
	}

	json.Unmarshal(responseData, &responseObj)
	return responseObj.Message
}

// pingHandler ::: retorna la respuesta del servicio ping + la repsuesta de pong
func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("pingHandler - request")
	var response map[string]interface{}
	str := callPongService()
	json.Unmarshal([]byte("{ \"send\": \"ping\", \"respond\": \""+str+"\" }"), &response)
	respondWithJSON(w, http.StatusOK, response)
}

// respondWithJSON ::: agrega la cabecera de tipo JSON en la respuesta de ping
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// initialiseRoutes ::: contiene el controlador y path de servicio ping
func (app *App) initialiseRoutes() {
	log.Println("start")
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/ping", pingHandler)
}

// run ::: levanta el servidor en puerto indicado
func (app *App) run() {
	log.Println("ping service on 8080")
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

// main ::: inicio del servicio ping
func main() {
	app := App{}
	app.initialiseRoutes()
	app.run()
}
