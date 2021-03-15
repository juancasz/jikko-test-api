package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"strconv"
	"sort"
)

func GetUserInfo(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	//Converts the id parameter from a string to an int
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err == nil {
		log.Println("Get info about user id #", id)

		//Retrieves the infos about the user
		user := FindUserByID(id)
		if(user.ID == 0){
			errorMessage := ErrorMessage{Message: "Usuario Inexistente"}
			json.NewEncoder(writer).Encode(errorMessage)
			return
		}
		json.NewEncoder(writer).Encode(user)
	} else {
		http.Error(writer, err.Error(), http.StatusBadRequest)
        return
	}
}

func SortArray(writer http.ResponseWriter, request *http.Request) { 
	log.Println("Sort Array")
	initHeaders(writer)
	var input UnsortedArray
	
	err := json.NewDecoder(request.Body).Decode(&input)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusBadRequest)
        return
    }

	array := input.Unsorted
	if(len(array) == 0){
		errorMessage := ErrorMessage{Message: "Sin Información"}
		json.NewEncoder(writer).Encode(errorMessage)
		return
	}
	log.Println("unsorted array",array)

	//sort array
	sort.Ints(array)

	//move repeated elements to the end of the array
	read := 0
	write := 0
	log.Println("array repeated",array)
	for read<len(array){
		// Swap the values pointed at by read and write.
		temp := array[write]	
		array[write] = array[read]
		array[read] = temp
		/* Advance the read pointer forward to the next unique value.  Since we
      	* moved the unique value to the write location, we compare values
      	* against array[write] instead of array[read].
      	*/
		for read<len(array) && (array[write] == array[read] || array[read] == temp){
			read = read+1
		}
		write = write + 1
	}

	output := SortedArray{Sorted: array}

	log.Println("sorted array",array)
	json.NewEncoder(writer).Encode(output)
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}
