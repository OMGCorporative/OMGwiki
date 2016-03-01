package main

//call these on rethink
// createTable()
// id := insertRecord()
// if id != "" {
// 	updateRecord(id)
// }

//create API

//template out
import (
	"encoding/json"
	"fmt"
	r "github.com/dancannon/gorethink"
	//"io/ioutil"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	router  *mux.Router
	session *r.Session
)

type Page struct {
	Title string
	Body  []byte
}

//API struct
type API struct {
	Message string "json:message"
}

//User struct
type User struct {
	Id string `gorethink:"id,omitempty", 
			   json:"id,omitempty"`
	Name string
}

func init() {
	var err error

	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "test",
		MaxOpen:  40,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func NewServer(addr string) *http.Server {
	// Setup router
	router = initRouting()

	// Create and start server
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}

func StartServer(server *http.Server) {
	log.Println("Starting server")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Error: %v", err)
	}
}

func initRouting() *mux.Router {

	r := mux.NewRouter()
	// handles these routes

	r.HandleFunc("/api/{name:[a-z]+}", Hello)
	r.HandleFunc("/", viewHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return r

}

//testing Json string with Handler
func viewHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Add("Content Type", "text/html")

	if err != nil {
		fmt.Println("Somethings not right")
	}

	renderTemplate(w, "index", map[string]interface{}{
	// "Items": items,
	// "Route": "all",
	})
}

func Hello(w http.ResponseWriter, req *http.Request) {

	urlParams := mux.Vars(req)
	name := urlParams["name"]
	HelloMessage := "Hello " + name
	message := API{HelloMessage}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Somethings wrong")
	}
	fmt.Fprintf(w, string(output))

}
