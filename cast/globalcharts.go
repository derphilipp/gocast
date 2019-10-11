package cast

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
)

func buildPodcastChartURL(countryCode string) string {
	baseURL, err := url.Parse(fmt.Sprintf("https://rss.itunes.apple.com/api/v1/%s/podcasts/top-podcasts/all/200/explicit.json", countryCode))
	if err != nil {
		log.Fatalf("Malformed URL: %v", err.Error())
	}
	return baseURL.String()
}

func findPodcastInCharts(podcastID string, pc *podcastCharts) (bool, int) {

	for i, podcast := range pc.Feed.Results {
		if podcast.ID == podcastID {
			return true, i + 1
		}
	}
	return false, 0
}

// PrintPodcastGlobalRank prints a given Podcast (via ID) inside the global charts of the given country
func PrintPodcastGlobalRank(podcastID, countryCode string) {
	countryCode = validateCountryCode(countryCode)
	url := buildPodcastChartURL(countryCode)
	body := loadBodyfromURL(url)
	result := podcastCharts{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	isFound, rank := findPodcastInCharts(podcastID, &result)
	if isFound {
		fmt.Println(rank)
	} else {
		os.Exit(1)
	}
}
