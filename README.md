# go-projects

golang

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
