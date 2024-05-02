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
  "music_file": "Music File",
  "illustration_file": "Music Illustration File"
}
```

### GET
- **/api/musics/download/:id** - Download a music by its id.
- **/api/musics/download?music_ids=1,2,3** - Download musics by their ids.
- **/api/musics/:id** - Get a music by its id (json response).
- **/api/musics** - Get all musics (json response).
- **/api/musics?music_ids=1,2,3** - Get musics by their ids (json response).

### PUT
- **/api/musics/:id** - Update a music by its id.

### DELETE
- **/api/musics/:id** - Delete a music by its id.
- **/api/musics?music_ids=1,2,3** - Delete musics by their ids.

## Playlists Endpoints
@ /api/playlists

### POST
- **/api/playlists** - Create a playlist.
```json
{
  "title": "Playlist Title",
  "description": "Playlist Description",
  "creator_id": "Playlist Creator Id",
  "musics": "Music Ids"
}
```

### GET
- **/api/playlists/:id** - Get a playlist by its id (json response).
- **/api/playlists** - Get all playlists (json response).
- **/api/playlists?playlist_ids=1,2,3** - Get playlists by their ids (json response).
- **/api/playlists/download/:id** - Download a playlist by its id.
- **/api/playlists/download?playlist_ids=1,2,3** - Download playlists by their ids.

### PUT
- **/api/playlists/:id** - Update a playlist by its id.

### DELETE
- **/api/playlists/:id** - Delete a playlist by its id.
- **/api/playlists?playlist_ids=1,2,3** - Delete playlists by their ids.

## Users Endpoints
@ /api/users

## Artists Endpoints
@ /api/artists

### POST
- **/api/artists** - Create an artist.
```json
{
  "name": "Artist Name",
  "description": "Artist Description",
}
```
