package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func main(){
    fmt.Println("REST API started")
    router := mux.NewRouter()

	InitDatas()
	buildRoutes(router)

    log.Fatal(http.ListenAndServe(":8000", router))
}

//Bar's routes
func buildRoutes(router *mux.Router) {
    prefix := "/api"
    router.HandleFunc(prefix + "/user/{id}", GetUserInfo).Methods("GET")
    router.HandleFunc(prefix+"/sort", SortArray).Methods("POST")
}