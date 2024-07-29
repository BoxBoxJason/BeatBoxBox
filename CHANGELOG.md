# BeatBoxBox CHANGELOG
All notable changes to this project will be documented in this file.

## [Frontend Release] - ????-??-??

### Added
- User authentication / registration page
- Music upload page
- Music player page (with download & search options)

### Changed

## [API Review & Fixes] - 2024-08-??

### Added
- OpenAPI documentation for all API endpoints
- API versioning with `/api/v1` prefix
- static API documentation

### Changed
- Moved /api routes to /api/v1 routes to implement versionning
- Fixed all detected errors & unexpected behaviors on handlers
- Added & Edited & Reviewed some controllers functionnalities

## [Controller Review & Fixes] - 2024-07-21

### Added
- Added CI/CD static & dynamic testing workflow to ensure quality

### Changed
- Fixed all detected errors & unexpected behaviors on controllers
- Added & Edited & Reviewed some model functionnalities

## [Model Review & Fixes] - 2024-07-05

### Added
- Added Roles to manage users permissions

### Changed
- Fixed all detected errors & unexpected behaviors on models


## [Playlist, Artist, Album REST API Release] - 2024-05-14

### Added
- Playlist CRUD operations
- Album CRUD operations
- Artist CRUD operations
- Quickstart script for easy setup (on Linux, Mac & Windows)

### Changed
- Updated ORM models for Playlist, Album, Artist, User & Music
- Updated database schema to include Playlist, Album, Artist tables

## [User REST API Release] - 2024-04-27

### Added
- ORM models for all database tables
- User registration API endpoint
- User authentication API endpoint
- All user CRUD operations
- Launch webserver & database via docker compose
- All tables (Playlist, Album, Artist) CRUD operations

### Changed

## [Musics REST API Release] - 2024-04-20

### Added

- Music upload API endpoint
- Music download API endpoint
- Music streaming API endpoint
- Music search API endpoint
- ORM model for music
- All music CRUD operations
