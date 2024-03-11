// app/musixmatch/musixmatch.go
package musixmatch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const musixmatchAPIKey = "YOUR_MUSIXMATCH_API_KEY"

// MusixmatchResponse represents the structure of the Musixmatch API response.
type MusixmatchResponse struct {
	Message struct {
		Body struct {
			Lyrics struct {
				Body string `json:"lyrics_body"`
			} `json:"lyrics"`
		} `json:"body"`
	} `json:"message"`
}

// GetLyrics retrieves the lyrics for a given track and artist from Musixmatch.
func GetLyrics(trackName, artistName string) (string, error) {
	url := fmt.Sprintf("https://api.musixmatch.com/ws/1.1/matcher.lyrics.get?q_track=%s&q_artist=%s&apikey=%s", trackName, artistName, musixmatchAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var musixmatchResponse MusixmatchResponse
	if err := json.Unmarshal(body, &musixmatchResponse); err != nil {
		return "", err
	}

	return musixmatchResponse.Message.Body.Lyrics.Body, nil
}
