# Gator blog aggregator

## Requirements
- To build this app, it's necessary to have 
    - Go version 1.26.2 
    - postgres 
    - docker
    - linux
- The app utilizes a dockerized postgres DB intsance to store all data
- The docker command
```bash
docker run --name gator -p 5432:5432 -e POSTGRES_USER=gator -e POSTGRES_PASSWORD=gator -d postgres
```
- The project's mod.go
```Go
module github.com/SamuelAboelkhir/blogAggregator

go 1.26.2

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/lib/pq v1.12.3 // indirect
)
```
## Installation
- To install the app, simply use the following command
```bash
go install github.com/SamuelAboelkhir/blogAggregator

# If you're asked to specify a version
go install github.com/SamuelAboelkhir/blogAggregator@latest
```
- Once installed, you'll also need to setup a `.gatorconfig.json` config file in your home directory
- Here's my config file as an example
```json
{
    "db_url":"postgres://gator:gator@localhost:5432/gator?sslmode=disable",
}
```
- It's important to specify the connection string of your postgres DB
- The app will add a username field to specify the currently logged in user
```json
{
    "db_url":"postgres://gator:gator@localhost:5432/gator?sslmode=disable",
    "current_user_name":"Samuel"
}
```
## Commands
- As a blog aggregator, this app can
    - `addfeed <url>`: Given a url for a blog feed to follow
    - `agg <time-between-requests>`: Aggregate posts from the feed
        - This command expects a string time value that in parsed into a duration internally to specify the intervals between aggregations
    - `following`: Display your followed feeds
    - `browse`: Browse the last few posts (2 by default but you can change the limit)
    - `login <name>`: Login to specific users
    - `register <name>`: Register new users
    - `follow <existing-feed-url>`: Follow any of the existing feeds
        - Note that the feed must have already been added to the database first
    - `unfollow <url>`: Unfollow feeds
    - `users`: Display all users (highlights the currently logged in one)
    - `resset`: Reset the database
