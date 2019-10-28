package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

type Container struct{
	Container []Person `json:"person_nodes"`
}

type Person struct{
	Id string `json:"_id"`
	Name string `json: "name"`
	Guid string `json: index`
	IsActive string `json: index`
	Balance string `json: index`
	Picture string `json: index`
	Age string `json: index`
	EyeColor string `json: index`
	Gender string `json: index`
	Company string `json: index`
	Email string `json: index`
	Phone string `json: index`
	Address string `json: index`
	About string `json: index`
	Registered string `json: index`
	Latitude string `json: index`
	Longitude string `json: index`
	Tags []string `json: friends`
	Friends Friends `json: friends`
	Greeting string `json: friends`
	FavoriteFruit string `json: friends`
	

}

type Friends struct {
	id int `json: index`
	name string `json: index`
}

func Home(response http.ResponseWriter, request *http.Request){
	json.NewEncoder(response).Encode("home")
}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	file, _ := ioutil.ReadFile("data.json")
	var data interface{}
	err := json.Unmarshal(file, &data)

	if err != nil {
        log.Fatal("Cannot unmarshal the json ", err)
    }
	  json.NewEncoder(response).Encode(data)
	
}
func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	file, _ := ioutil.ReadFile("data.json")
	data := Container{}
	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Container); i++ {
		if data.Container[i].Id == id {
			fmt.Println("val: ", data.Container[i])
			json.NewEncoder(response).Encode(data.Container[i])
		}
		fmt.Println("val: ", data.Container[i])	
	}
}

func GetPersonByEarning(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	balance := params["balance"]
	file, _ := ioutil.ReadFile("data.json")
	data := Container{}
	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.Container); i++ {
		if data.Container[i].Balance == balance {
			fmt.Println("val: ", data.Container[i])
			json.NewEncoder(response).Encode(data.Container[i])
		}
		fmt.Println("val: ", data.Container[i])	
	}
}

func main() {
	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/person/{earn}", GetPersonByEarning).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}