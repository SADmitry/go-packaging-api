# Go Packaging API

This is a simple Go API that calculates the optimal package distribution for a given number of items.

## Installation

### 1. Install Go (if not installed)

On macOS, install Go using Homebrew:

```sh
brew install go
```

Verify the installation:

```sh
go version
```

### 2. Clone the repository

```sh
git clone git@github.com:SADmitry/go-packaging-api.git
cd go-packaging-api
```

## Running the API Locally

### 1. Build the project

```sh
go build -o go-packaging-api
```

### 2. Run the API

```sh
./go-packaging-api
```

The API will start on `http://localhost:8080`.

### 3. Test the API

Send a test request using `curl`:

```sh
curl -X POST http://localhost:8080/calculate -d '{"items": 12001}' -H "Content-Type: application/json"
```

Expected response:

```json
{"packs":{"5000":2,"2000":1,"250":1}}
```

## Running with Docker

### 1. Build the Docker image

```sh
docker build -t go-packaging-api .
```

### 2. Run the container

```sh
docker run -p 8080:8080 go-packaging-api
```

### 3. Test the API (same as above)

## Configuration

The API reads package sizes from `config.json`. Modify this file to adjust package sizes. Example `config.json`:

```json
{
  "pack_sizes": [250, 500, 1000, 2000, 5000]
}
```

## OpenAPI Documentation

The API is documented using OpenAPI. You can view the API spec using Swagger UI:

```sh
docker run -p 8081:8080 -e SWAGGER_JSON=/api/openapi.yaml -v $(pwd)/openapi.yaml:/api/openapi.yaml swaggerapi/swagger-ui
```

Then open `http://localhost:8081` in your browser.
