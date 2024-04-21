/*
Contains the utility functions for the music handlers
*/
package music_handler

// getAuthorizedMusicFields filters a list of requested fields to only return the fields that the user is authorized to see.
// Returns all fields if the initial list is empty.
func getAuthorizedMusicFields(fields []string) []string {
	authorized_fields_map := map[string]bool{
		"Id": true, "Title": true, "Artist": true, "Genres": true,
		"Album": true, "Path": true, "Nblistened": true, "Ratings": true, "Nbratings": true,
	}

	if len(fields) == 0 {
		// Convert map keys to slice
		authorized_fields := make([]string, 0, len(authorized_fields_map))
		for field := range authorized_fields_map {
			authorized_fields = append(authorized_fields, field)
		}
		return authorized_fields
	}

	authorized_fields := make([]string, 0, len(fields))
	for _, field := range fields {
		if _, ok := authorized_fields_map[field]; ok {
			authorized_fields = append(authorized_fields, field)
		}
	}
	return authorized_fields
}

// getAuthorizedMusicFilters filters a map of requested filters to only return the filters that the user is authorized to see.
// Returns all filters if the initial map is empty.
func getAuthorizedMusicFilters(filters map[string]interface{}) map[string]interface{} {
	authorized_filters_map := map[string]bool{
		"Id": true, "Title": true, "Artist": true, "Genres": true, "Album": true,
		"Path": true, "Nblistened": true, "Ratings": true, "Nbratings": true,
	}

	authorized_filters := make(map[string]interface{})
	for filter, value := range filters {
		if _, ok := authorized_filters_map[filter]; ok {
			authorized_filters[filter] = value
		}
	}
	return authorized_filters
}
