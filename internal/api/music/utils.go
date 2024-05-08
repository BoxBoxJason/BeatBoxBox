/*
Contains the utility functions for the music handlers
*/
package music_handler

import (
	"net/url"
	"strconv"
	"strings"
)

// parseURLParams parses the URL parameters and returns a map of the authorized fields to update
func parseURLParams(query_params url.Values) (map[string]interface{}, error) {
	update_dict := make(map[string]interface{})
	title := query_params.Get("title")
	if title != "" {
		update_dict["Title"] = title
	}
	artist_id_str := query_params.Get("artist_id")
	if artist_id_str != "" {
		artist_id, err := strconv.Atoi(artist_id_str)
		if err != nil {
			return update_dict, err
		}
		update_dict["ArtistId"] = artist_id
	}
	album_id_str := query_params.Get("album_id")
	if album_id_str != "" {
		album_id, err := strconv.Atoi(album_id_str)
		if err != nil {
			return update_dict, err
		}
		update_dict["AlbumId"] = album_id
	}
	raw_genres := query_params.Get("genres")
	if raw_genres != "" {
		genres := strings.Split(raw_genres, ",")
		update_dict["Genres"] = genres
	}

	return update_dict, nil
}
