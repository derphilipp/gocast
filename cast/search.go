package cast

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func buildPodcastInfoURL(searchterm, country string) string {
	baseURL, err := url.Parse("https://itunes.apple.com/search")
	if err != nil {
		log.Fatalf("Malformed URL: %v", err.Error())
	}
	params := url.Values{}
	params.Add("media", "podcast")
	params.Add("country", country)
	params.Add("term", searchterm)
	baseURL.RawQuery = params.Encode()
	return baseURL.String()
}

func validateCountryCode(countrycode string) string {
	if len(countrycode) != 2 {
		log.Fatalf("%s is an invalid countrycode. Countrycodes must be 2 characters long. \nExample: DE EN FR", countrycode)
	}
	return strings.ToLower(countrycode)
}

// SearchPodcast searches for Podcast using searchterm and countrycode
func SearchPodcast(searchterm, countrycode string) {
	countrycode = validateCountryCode(countrycode)
	url := buildPodcastInfoURL(searchterm, countrycode)
	result := podcastJSONInfo{}
	body := loadBodyfromURL(url)
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(result)
}
