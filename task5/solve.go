package main

import (
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
)

type UrlAdd struct {
	Url string `json:"url"`
}

type PostResponse struct {
	Key string `json:"key"`
}

type LongUrl struct {
	Location string `json:"Location"`
}

var (
	urls   = make(map[int]LongUrl)
	nextId = 0
)

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	var newUrl UrlAdd
	err := json.NewDecoder(r.Body).Decode(&newUrl)
	if err != nil {
		panic(err)
	}

	urls[nextId] = LongUrl{Location: newUrl.Url}

	response, err := json.Marshal(PostResponse{Key: strconv.Itoa(nextId)})
	if err != nil {
		panic(err)
	}
	nextId++

	w.Write(response)
}

func getURl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["key"])
	if err != nil {
		panic(err)
	}

	if longUrl, ok := urls[key]; ok {
		response, err := json.Marshal(longUrl)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusMovedPermanently)
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	// handles
	router.HandleFunc("/", shortenUrl).Methods("POST")
	router.HandleFunc("/{key}", getURl)

	http.ListenAndServe(":8082", router)
}
