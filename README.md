# BeatBoxBox

BeatBoxBox is a comprehensive music serving webserver that enables users to upload, download, and listen to music online. It offers a rich set of features for interacting with music through a web interface, supported by a robust backend.

## Features

### Playlists
Manage custom playlists to organize your favorite tracks.

### User Authentication
Secure login and registration for users.

### User Preferences Analysis & Suggestions
Leverage user data to recommend and customize music suggestions.

### Music Upload
Allow users to upload their own music files to the platform.

### Music Download
Enable users to download music for offline listening.

### Music Streaming
Stream music directly without the need to download.

### Music Search
Search functionality to find music by various criteria like artist, title, or album.

### Music Recommendation
Deliver music recommendations based on user behavior and preferences.

### Music Rating
Facilitate users to rate music, which influences recommendations.

## Technologies

### Backend
Golang based backend server that serves the frontend and handles the business logic.

### Frontend
Vue.js based frontend that interacts with the backend server.

## Getting Started

To get the project running locally:

1. Clone the repository:
`git clone https://github.com/BoxBoxJason/BeatBoxBox.git`

2. Install the dependencies:
   - npm
   - go
   - vue-cli

3. Build the frontend using vue-cli:
   1. `cd frontend`
   2. `npm install`
   3. `npm run build`

4. Run the backend server:
   1. `cd backend`
   2. `go run cmd/beatboxbox/main.go`

5. Access the web interface with your web browser at `https://localhost:8080`
