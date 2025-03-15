# go-projects

golang

## Init go project

`go mod init github.com/mmxcrono/some-project`

## Update dependencies

`go mod tidy`

## Build

`go build main.go`

OR

`go build main.go myapp`

## Run

`go run main.go`

If there are multiple go files

`go run .`

## Docker

Build

`docker build -t my-go-app .`

Run

`docker run --rm -p 8080:8080 --name my-go-container my-go-app`

Shell

`docker run -it --rm my-go-app sh`

## Folder Structure

```
/project
  ├── cmd/ # Main applications
  ├── pkg/ # Reusable packages
  ├── internal/ # Internal packages (not exported)
  ├── api/ # API definitions
  ├── config/ # Configuration files
  ├── scripts/ # Build and automation scripts
  ├── test/ # Additional test data
  ├── go.mod # Go module file
  ├── go.sum # Module dependencies checksum
```

- Use short, lowercase, singular package names (e.g., auth, not authentication).
- Avoid generic names like util or helpers.

## MySQL / MariaDB

Install at the go project level, not where main.go

`go get -u github.com/go-sql-driver/mysql`

Run mysql via docker

`docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=db_name -p 3306:3306 -d mysql:latest`
