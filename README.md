# <GOREST-API>

## Description

Example of simple API rest with GOLANG 1.18. 

Packages in /src: 
- /api -> package with routes and swagger config, using gin (https://github.com/gin-gonic/gin) for handler and context.
- /config -> package to open config files with envoirments*.env.
- /adapter ->  Integration with other services . 
- /adapter/model -> Structs of request for the integrations API.
- /docs -> swagger documentation, using https://github.com/swaggo/files and https://github.com/swaggo/swag. For integration with GIN https://github.com/swaggo/gin-swagger.
- /model -> package with any struct or models for the application.
- /service -> package with validation and logic.  
- main -> config and initializing routes.
 
## Contents

- [Installation](#installation)
- [Usage](#usage)

## Installation

- Install Go 1.18 or later from https://go.dev/doc/install.

## Usage

- Clone the project.
- Open cmd in the base path and run the command "go mod tidy" (doc reference https://go.dev/ref/mod#go-mod-tidy).
- Move into the folder "src" and run command "go run main.go" (doc reference https://go.dev/ref/mod#go-run).

For build application you need to run command "go build [-o output] [build flags] [packages]" (doc reference of build flags https://pkg.go.dev/cmd/go).

File src/notes contains the notes of the  