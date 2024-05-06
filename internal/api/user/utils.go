package user_handler

import (
	"net/url"
)

// parseURLParams parses the URL parameters and returns a map of the authorized fields to update
func parseURLParams(query_params url.Values) (map[string]interface{}, error) {
	update_dict := make(map[string]interface{})
	username := query_params.Get("username")
	if username != "" {
		update_dict["pseudo"] = username
	}
	email := query_params.Get("email")
	if email != "" {
		update_dict["email"] = email
	}

	return update_dict, nil
}
