package main

import (
  "encoding/json"
  "flag"
  "os"
  "strings"

  "github.com/gocolly/colly/v2"
)

type Review struct {
  Album string `json:"album"`
  Artist string `json:"artist"`
  Score string `json:"score"`
}

func main() {
  var artistName string
  flag.StringVar(&artistName, "artist", "outkast", "Name of artist to search")
  flag.Parse()

  reviews := make([]*Review, 0)

  albumCollector := colly.NewCollector(
    colly.AllowedDomains("pitchfork.com"),
  )

  reviewCollector := colly.NewCollector(
    colly.AllowedDomains("pitchfork.com"),
  )


  albumCollector.OnHTML(`a[href].review__link`, func(e *colly.HTMLElement) {
    reviewURL := e.Request.AbsoluteURL(e.Attr("href"))
    if strings.Index(reviewURL, "/reviews/") != -1 {
      reviewCollector.Visit(reviewURL)
    }
  })

  reviewCollector.OnHTML(".review-detail", func (e *colly.HTMLElement) {
    album := e.ChildText(".single-album-tombstone__review-title")
    artist := e.ChildText(".single-album-tombstone__artist-links li a")
    score := e.ChildText(".score")

    review := &Review{
      Album: album,
      Artist: artist,
      Score: score,
    }

    reviews = append(reviews, review)
  })

  albumCollector.Visit("https://pitchfork.com/search/?query=" + artistName)

  enc := json.NewEncoder(os.Stdout)
  enc.SetIndent("", " ")

  enc.Encode(reviews)
}
