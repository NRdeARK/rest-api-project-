package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var Moves []mov

// {
// 	"ImdbID" : "1",
// 	"Title" : "LOL",
// 	"Rating" : 9.9,
// 	"IsSuperHero" : true
// 	}

type mov struct {
	ImdbID string `json:"imdb_id"`
	Title string `json:"title"`
	Rating float64 `json:"rating"`
	IsSuperHero bool `json:"is_super_hero"`
}

func movieHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Fprintf(w, "hello %s type\n", method)// like print but instead of print send it to w -> responce
	if method == "GET" {
		fmt.Fprintf(w, "some thing 1e\n")
	}
	if method == "POST" {
		fmt.Fprintf(w, "I got it\n")
		body,err := ioutil.ReadAll(r.Body)
		if(err != nil) {
			fmt.Printf("%#v %#v\n", string(body), err)
			return
		}
		m := mov{}
		err = json.Unmarshal(body, &m)
		Moves = append(Moves,m)
		fmt.Fprintf(w, "%#v\n", Moves)
		
	}
	return
}

func main() {
	http.HandleFunc("/mov",movieHandler)//func handling req
	err := http.ListenAndServe("localhost:25653", nil)
	log.Fatal(err)
}