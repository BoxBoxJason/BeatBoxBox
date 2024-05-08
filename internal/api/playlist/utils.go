package playlist_handler

import (
	"net/url"
	"strconv"
	"strings"
)

func parseURLParams(query_params url.Values) (map[string]interface{}, error) {
	update_dict := make(map[string]interface{})
	title := query_params.Get("title")
	if title != "" {
		update_dict["title"] = title
	}
	description := query_params.Get("description")
	if description != "" {
		update_dict["description"] = description
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
