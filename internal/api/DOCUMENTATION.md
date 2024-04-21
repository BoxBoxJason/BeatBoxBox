# API documentation

This is the BeatBoxBox API documentation. It is a work in progress.

## Musics Endpoints
@ /api/musics

### POST

- **/api/musics/upload** - Upload a music file.
```json
{
  "title": "Music Title",
  "artist": "Music Artist",
  "album": "Music Album",
  "genres": "Music Genres",
  "year": "Music Year",
  "file": "Music File"
}
```

### GET

- **/api/musics/download/:id** - Download a music by its id.
- **/api/musics/download?music_ids=1,2,3** - Download musics by their ids.
- **/api/musics/:id** - Get a music by its id (json response).
- **/api/musics** - Get all musics (json response).
- **/api/musics?music_ids=1,2,3** - Get musics by their ids (json response).

### PUT

- **/api/musics/update/:id** - Update a music by its id.

### DELETE

- **/api/musics/delete/:id** - Delete a music by its id.
- **/api/musics/delete?music_ids=1,2,3** - Delete musics by their ids.

## Playlists Endpoints
@ /api/playlists

## Users Endpoints
@ /api/users

## Artists Endpoints
@ /api/artists
