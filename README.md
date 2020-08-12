# Simple Go Service
Simple REST service wrote on Golang

### Main features:
* Login
* Registration
* Get user(s)
* Edit user profile
* Upload user avatar
* Get user(s) avatars
* Crop avatars to 160x160 if they are bigger

### &#x1F534; Before run service
* You have to change secret auth key in `src/web/auth.go` line `18`

### Run service with docker:
* Clone repository to your local machine
* Make sure you have installed Docker and Docker Compose (if you don't have real mysql database)
* go into `/src` folder and run `docker-compose up` command
* Run service `go run main.go`

### Run service with real DB:
* Clone repository to your local machine
* Edit `config.json` with your database credentials
* Run service `go run main.go`
