package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chrom9k/GOCRUDAPI/users"
	"github.com/gorilla/mux"
)

var jsonEncode = json.Marshal
var jsonDecode = json.Unmarshal

func main(w http.ResponseWriter, r *http.Request) {
	response, _ := jsonEncode("You are on the site")
	fmt.Fprint(w, string(response))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var response []byte
	users, error := users.GetUsers()
	if error == nil {
		response, _ = jsonEncode(users)
	} else {
		response, _ = jsonEncode("No data")
	}

	fmt.Fprintf(w, string(response))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var response []byte
	variables := mux.Vars(r)
	id, _ := strconv.Atoi(variables["id"])
	user, error := users.GetUser(id)
	if error == nil {
		response, _ = jsonEncode(user)
	} else {
		response, _ = jsonEncode("No data")
	}
	fmt.Fprintf(w, string(response))
}

func createUser(w http.ResponseWriter, r *http.Request) {

}

func updateUser(w http.ResponseWriter, r *http.Request) {

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

}

func SetRouting() {
	router := mux.NewRouter()

	router.HandleFunc("/", main).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id:[0-9]+}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id:[0-9]+}", deleteUser).Methods("DELETE")

	http.Handle("/", router)
}
