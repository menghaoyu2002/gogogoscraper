package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func HandleWatch(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	showName := p.ByName("name")
	episodeNumber, err := strconv.Atoi(p.ByName("episode"))

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		tmpl, _ := template.ParseFiles("./static/homepage.html")
		tmpl.Execute(rw, "Error: Invalid Episode Number")
		return 
	}

	res := make(chan string, 10)
	Scrape(Anime{showName, episodeNumber }, res)
	animeURL := <- res

	if animeURL == "" {
		rw.WriteHeader(http.StatusBadRequest)
		tmpl, _ := template.ParseFiles("./static/homepage.html")
		tmpl.Execute(rw, "Error: Anime not found. Check that the name of the show and episode number are both valid")
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
		tmpl, _ := template.ParseFiles("./static/homepage.html")
		tmpl.Execute(rw, "")
	})

	router.GET("/watch/:name/:episode", HandleWatch)

	router.GET("/watch/", func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rw.WriteHeader(http.StatusBadRequest)
		tmpl, _ := template.ParseFiles("./static/homepage.html")
		tmpl.Execute(rw, "Error: Please Enter BOTH a show name and episode number")
	})

	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}
