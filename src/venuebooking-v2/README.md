To update the Database config, goto `app/venuebooking_v1/server/internal/config/config.go` 
and change the Database connection string in function called `DBConnectionString`.

Run all commands from projects home folder.

1) `go mod tidy`  <= to download all project's dependencies.
2) `go mod vendor`  <= for vendoring the dependencies.
3) `go run cmd/venuebooking-v1/main.go`  <= to run the application.


Project Structure Explaind :

/api

    OpenAPI/Swagger specs, JSON schema files, protocol definition files.

/lib

    Library code that's ok to use by external applications. Other projects will import these libraries expecting them to work.


/vendor

    Application dependencies, currently managed by go mod.

/cmd

    Main applications for this project.The directory name for each application should match the name of the executable you want to have.

/app

    Contains the main business logic of application.
    
    /db :
        All Database queries required by your application.

    /internal :
        Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself.  
        