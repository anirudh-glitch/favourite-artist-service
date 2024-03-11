package handlers

import (
	"favourite-artist-service/app/lastfm"
	"favourite-artist-service/app/musixmatch"
	"log"

	"github.com/gin-gonic/gin"
)

// GetInformation fetches information based on the provided parameters.
func GetInformation(region, trackName, artistName string) (gin.H, error) {
	lastfmResult, err := lastfm.GetTopTrack(region)
	if err != nil {
		log.Printf("Error getting top track from LastFM: %v", err)
		return nil, err
	}

	musixmatchResult, err := musixmatch.GetLyrics(trackName, artistName)
	if err != nil {
		log.Printf("Error getting lyrics from Musixmatch: %v", err)
		return nil, err
	}

	// Add more logic as needed

	response := gin.H{
		"lastfm_track":      lastfmResult,
		"musixmatch_lyrics": musixmatchResult,
	}

	return response, nil
}
