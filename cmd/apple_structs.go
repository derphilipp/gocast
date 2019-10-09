package cmd

import (
	"fmt"
	"strings"
	"time"
)

type podcastJSONInfo struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		ArtistName             string    `json:"artistName"`
		ArtworkURL100          string    `json:"artworkUrl100"`
		ArtworkURL30           string    `json:"artworkUrl30"`
		ArtworkURL60           string    `json:"artworkUrl60"`
		ArtworkURL600          string    `json:"artworkUrl600"`
		CollectionCensoredName string    `json:"collectionCensoredName"`
		CollectionExplicitness string    `json:"collectionExplicitness"`
		CollectionHdPrice      int       `json:"collectionHdPrice"`
		CollectionID           int       `json:"collectionId"`
		CollectionName         string    `json:"collectionName"`
		CollectionPrice        float64   `json:"collectionPrice"`
		CollectionViewURL      string    `json:"collectionViewUrl"`
		ContentAdvisoryRating  string    `json:"contentAdvisoryRating"`
		Country                string    `json:"country"`
		Currency               string    `json:"currency"`
		FeedURL                string    `json:"feedUrl"`
		GenreIds               []string  `json:"genreIds"`
		Genres                 []string  `json:"genres"`
		Kind                   string    `json:"kind"`
		PrimaryGenreName       string    `json:"primaryGenreName"`
		ReleaseDate            time.Time `json:"releaseDate"`
		TrackCensoredName      string    `json:"trackCensoredName"`
		TrackCount             int       `json:"trackCount"`
		TrackExplicitness      string    `json:"trackExplicitness"`
		TrackHdPrice           int       `json:"trackHdPrice"`
		TrackHdRentalPrice     int       `json:"trackHdRentalPrice"`
		TrackID                int       `json:"trackId"`
		TrackName              string    `json:"trackName"`
		TrackPrice             float64   `json:"trackPrice"`
		TrackRentalPrice       int       `json:"trackRentalPrice"`
		TrackViewURL           string    `json:"trackViewUrl"`
		WrapperType            string    `json:"wrapperType"`
	} `json:"results"`
}

func (p podcastJSONInfo) String() string {
	var resultString strings.Builder
	for i, podcast := range p.Results {
		if i != 0 {
			resultString.WriteString("\n")
		}
		resultString.WriteString(fmt.Sprintf("Name: %s\n", podcast.CollectionName))
		resultString.WriteString(fmt.Sprintf("Artist: %s\n", podcast.ArtistName))
		resultString.WriteString(fmt.Sprintf("ID: %d\n", podcast.CollectionID))
		resultString.WriteString("Genres: \n")
		for _, genre := range podcast.Genres {
			resultString.WriteString(fmt.Sprintf("    %s\n", genre))
		}
		resultString.WriteString("GenresIDs: \n")
		for _, genreID := range podcast.GenreIds {
			resultString.WriteString(fmt.Sprintf("    %s\n", genreID))
		}
	}
	return resultString.String()
}

type podcastCharts struct {
	Feed struct {
		Title  string `json:"title"`
		ID     string `json:"id"`
		Author struct {
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"author"`
		Links []struct {
			Self      string `json:"self,omitempty"`
			Alternate string `json:"alternate,omitempty"`
		} `json:"links"`
		Copyright string `json:"copyright"`
		Country   string `json:"country"`
		Icon      string `json:"icon"`
		Updated   string `json:"updated"`
		Results   []struct {
			ArtistName    string `json:"artistName"`
			ID            string `json:"id"`
			ReleaseDate   string `json:"releaseDate"`
			Name          string `json:"name"`
			Kind          string `json:"kind"`
			Copyright     string `json:"copyright,omitempty"`
			ArtworkURL100 string `json:"artworkUrl100"`
			Genres        []struct {
				GenreID string `json:"genreId"`
				Name    string `json:"name"`
				URL     string `json:"url"`
			} `json:"genres"`
			URL                   string `json:"url"`
			ArtistID              string `json:"artistId,omitempty"`
			ArtistURL             string `json:"artistUrl,omitempty"`
			ContentAdvisoryRating string `json:"contentAdvisoryRating,omitempty"`
		} `json:"results"`
	} `json:"feed"`
}

type podcastGenreCharts struct {
	Feed struct {
		Author struct {
			Name struct {
				Label string `json:"label"`
			} `json:"name"`
			URI struct {
				Label string `json:"label"`
			} `json:"uri"`
		} `json:"author"`
		Entry []struct {
			ImName struct {
				Label string `json:"label"`
			} `json:"im:name,omitempty"`
			ImImage []struct {
				Label      string `json:"label"`
				Attributes struct {
					Height string `json:"height"`
				} `json:"attributes"`
			} `json:"im:image"`
			Summary struct {
				Label string `json:"label"`
			} `json:"summary"`
			ImPrice struct {
				Label      string `json:"label"`
				Attributes struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"attributes"`
			} `json:"im:price"`
			ImContentType struct {
				Attributes struct {
					Term  string `json:"term"`
					Label string `json:"label"`
				} `json:"attributes"`
			} `json:"im:contentType"`
			Rights struct {
				Label string `json:"label"`
			} `json:"rights,omitempty"`
			Title struct {
				Label string `json:"label"`
			} `json:"title,omitempty"`
			Link struct {
				Attributes struct {
					Rel  string `json:"rel"`
					Type string `json:"type"`
					Href string `json:"href"`
				} `json:"attributes"`
			} `json:"link"`
			ID struct {
				Label      string `json:"label"`
				Attributes struct {
					ImID string `json:"im:id"`
				} `json:"attributes"`
			} `json:"id"`
			ImArtist struct {
				Label string `json:"label"`
			} `json:"im:artist,omitempty"`
			Category struct {
				Attributes struct {
					ImID   string `json:"im:id"`
					Term   string `json:"term"`
					Scheme string `json:"scheme"`
					Label  string `json:"label"`
				} `json:"attributes"`
			} `json:"category"`
			ImReleaseDate struct {
				Label      string `json:"label"`
				Attributes struct {
					Label string `json:"label"`
				} `json:"attributes"`
			} `json:"im:releaseDate"`
		} `json:"entry"`
		Updated struct {
			Label string `json:"label"`
		} `json:"updated"`
		Rights struct {
			Label string `json:"label"`
		} `json:"rights"`
		Title struct {
			Label string `json:"label"`
		} `json:"title"`
		Icon struct {
			Label string `json:"label"`
		} `json:"icon"`
		Link []struct {
			Attributes struct {
				Rel  string `json:"rel"`
				Type string `json:"type"`
				Href string `json:"href"`
			} `json:"attributes,omitempty"`
		} `json:"link"`
		ID struct {
			Label string `json:"label"`
		} `json:"id"`
	} `json:"feed"`
}
