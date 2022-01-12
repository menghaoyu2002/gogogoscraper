package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Anime struct {
	showName string
	episodeNumber int
}

func Scrape(anime Anime) {
	url := "https://www3.gogoanime.cm/" + anime.FormatAnimeInfo()
	fmt.Println(url)
	c := colly.NewCollector()
	c.OnHTML("a[data-video]", func (e *colly.HTMLElement) {
		fmt.Println(e.Attr("data-video"))
	})
	c.Visit(url)
}

func (anime Anime) FormatAnimeInfo() string {
	formattedShowName := strings.ReplaceAll(strings.ToLower(anime.showName), " ", "-")
	formattedEpisodeName := "episode-" + strconv.Itoa(anime.episodeNumber)
	return formattedShowName + "-" + formattedEpisodeName
}
