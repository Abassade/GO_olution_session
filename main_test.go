package main

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
    return router
}

func TestGetPeople(t *testing.T) {
    request, _ := http.NewRequest("GET", "/people", nil)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
    assert.Equal(t, 200, response.Code, "OK response is expected")
}

// func TestGetPerson(t *testing.T) {
//     request, _ := http.NewRequest("GET", "/person/sfg", nil)
//     response := httptest.NewRecorder()
//     Router().ServeHTTP(response, request)
//     assert.Equal(t, 200, response.Code, "OK response is expected")
// }

// func TestGetPersonWithbalance(t *testing.T) {
//     request, _ := http.NewRequest("GET", "/person/$5655", nil)
//     response := httptest.NewRecorder()
//     Router().ServeHTTP(response, request)
//     assert.Equal(t, 200, response.Code, "OK response is expected")
// }