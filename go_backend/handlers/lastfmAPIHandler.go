package handlers

import (
	"encoding/json"
	"fmt"
)

// take in spotify response obj, return list of recommendations (artists + albums + track)
func GetLastFMRecsFromSpotify(response SpotifyResponseObj) ([]byte, error) {
	// track -- artist tags, album tags, track tags, track similar
	// album -- artist tags, album tags
	// artist -- artist tags, artists similar

	// map with arist, track, album keys
	resp := make(map[string][]string)
	// resp["Artist"] = []string{}
	// resp["ArtistSimilar"] = []string{}
	// resp["Track"] = []string{}
	// resp["TrackSimilar"] = []string{}
	// resp["Album"] = []string{}

	if response.TrackName != nil {
		// track similar
		// get similar tracks using track name, artist name
		resp["TrackSimilar"] = GetLastFMSimilarTracksBuiltin(*response.TrackName, response.ArtistName)
	}

	// always executes anyways
	if response.AlbumName != nil {
		// artists similar
		resp["ArtistSimilar"] = GetLastFMSimilarArtistsBuiltin(response.ArtistName)
	}

	tags := []string{}

	// use tags
	resp = GetLastFMRecsByTags(resp, tags)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("error marshalling LastFM response recs: %w", err)
	}

	return jsonResp, nil
}

// given track name and artist name, return track recommendations from LastFM
func GetLastFMSimilarTracksBuiltin(trackName string, artistName string) []string {
	return []string{}
}

// given artist name, return artist recs from LastFM
func GetLastFMSimilarArtistsBuiltin(artistName string) []string {
	return []string{}
}

// given resp map and tags list, return updated resp map
func GetLastFMRecsByTags(resp map[string][]string, tags []string) map[string][]string {
	resp["Track"] = GetLastFMTracksByTag(tags)
	resp["Album"] = GetLastFMAlbumsByTag(tags)
	resp["Artist"] = GetLastFMArtistsByTag(tags)
	return resp
}

// given list of tags, return track recommendations from LastFM
func GetLastFMTracksByTag(tags []string) []string {
	return []string{}
}

// given list of tags, return album recommendations from LastFM
func GetLastFMAlbumsByTag(tags []string) []string {
	return []string{}
}

// given list of tags, return artist recommendations from LastFM
func GetLastFMArtistsByTag(tags []string) []string {
	return []string{}
}
