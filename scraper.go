package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Anime struct {
	showName string
	episodeNumber int
}

func Scrape(anime Anime, res chan string) {
	url := "https://www3.gogoanime.cm/" + anime.FormatAnimeInfo()
	c := colly.NewCollector()
	c.OnHTML("a[data-video]", func (e *colly.HTMLElement) {
		res <- e.Attr("data-video")
	})
	c.OnScraped(func(r *colly.Response) {
		res <- ""  // send an empty string to signify that no results were found
		log.Println("SCRAPED: " + url)
	}) 
	c.Visit(url)
}

func (anime Anime) FormatAnimeInfo() string {
	formattedShowName := strings.ReplaceAll(strings.ToLower(anime.showName), " ", "-")
	formattedEpisodeName := "episode-" + strconv.Itoa(anime.episodeNumber)
	return formattedShowName + "-" + formattedEpisodeName
}
