package cast

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
)

func buildPodcastGenreChartURL(countryCode, genreID string) string {
	baseURL, err := url.Parse(fmt.Sprintf("https://itunes.apple.com/%s/rss/toppodcasts/limit=200/genre=%s/explicit=true/json", countryCode, genreID))
	if err != nil {
		log.Fatalf("Malformed URL: %v", err.Error())
	}
	return baseURL.String()

}

func findPodcastInGenreCharts(podcastID string, pc *podcastGenreCharts) (bool, int) {
	for i, podcast := range pc.Feed.Entry {
		if podcast.ID.Attributes.ImID == podcastID {
			return true, i + 1
		}
	}
	return false, 0
}

func PrintPodcastGenreRank(podcastID, genreID, countryCode string) {
	countryCode = validateCountryCode(countryCode)
	url := buildPodcastGenreChartURL(countryCode, genreID)
	body := loadBodyfromURL(url)
	result := podcastGenreCharts{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	isFound, rank := findPodcastInGenreCharts(podcastID, &result)
	if isFound {
		fmt.Println(rank)
	} else {
		os.Exit(1)
	}
}
