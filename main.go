package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func HandleWatch(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	showName := p.ByName("name")
	episodeNumber, err := strconv.Atoi(p.ByName("episode"))

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "Invalid Episode Number")
		return 
	}

	res := make(chan string, 10)
	Scrape(Anime{showName, episodeNumber }, res)
	animeURL := <- res

	if animeURL == "" {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "Error: Anime not found. Check that the name of the show and episode number are both valid")
		return 
	}

	animeURL = "https:" + animeURL
	tmpl, _ := template.ParseFiles("./static/player.html")
	tmpl.Execute(rw, animeURL)
}

func main () { 
	router := httprouter.New()
	router.ServeFiles("/styles/*filepath", http.Dir("./static/styles"))
	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.ServeFile(rw, r, "./static/homepage.html")
	})
	router.GET("/watch/:name/:episode", HandleWatch)

	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
