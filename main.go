package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Story ...
type Story struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

// Stories ...
type Stories []Story

func getAllStories(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	myStories := Stories{Story{Title: "some title",Desc: "desc",Content: "cont"}}
	json.NewEncoder(w).Encode(myStories)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "home?")
}

func handleRequest(){
	PORT := "8080"
	http.HandleFunc("/api", homePage)
	http.HandleFunc("/api/stories", getAllStories)
	fmt.Println("Server starting on port...", PORT)
	log.Fatal(http.ListenAndServe(":" + PORT, nil))
}

func main(){
	handleRequest()
}