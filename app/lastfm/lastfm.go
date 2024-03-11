// app/lastfm/lastfm.go
package lastfm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const lastFMAPIKey = "YOUR_LASTFM_API_KEY"

// LastFMResponse represents the structure of the Last.fm API response.
type LastFMResponse struct {
	TopTrack struct {
		Name   string `json:"name"`
		Artist struct {
			Name string `json:"name"`
		} `json:"artist"`
	} `json:"track"`
}

// GetTopTrack retrieves the top track in the specified region from Last.fm.
func GetTopTrack(region string) (string, error) {
	url := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&country=%s&api_key=%s&format=json", region, lastFMAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var lastFMResponse LastFMResponse
	if err := json.Unmarshal(body, &lastFMResponse); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s by %s", lastFMResponse.TopTrack.Name, lastFMResponse.TopTrack.Artist.Name), nil
}
