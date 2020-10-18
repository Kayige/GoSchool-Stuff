To update the Database config, goto `app/venuebooking_v1/server/internal/config/config.go` 
and change the Database connection string in function called `DBConnectionString`.

Run all commands from projects home folder.

1) `go mod tidy`  <= to download all project's dependencies.
2) `go mod vendor`  <= for vendoring the dependencies.
3) `go run cmd/venuebooking-v1/main.go`  <= to run the application.
