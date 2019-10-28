# Geo location API

`geoapi` API provides endpoints for geo location behaviors

## Required environment variables

| Variable                      | Description                                | Default    | Optional |
|-------------------------------|--------------------------------------------|:----------:|:--------:|
| `GEOAPI_HTTP_PORT`            | TCP port on which to start HTTP server     |    8080    |   true   |
| `GEOAPI_LOCATION_ADDRESS`     | Location service address                   |     -      |    -     |
| `GEOAPI_MODE`                 | Environment mode                           |     dev    |   true   |

## Run linting and testing

```
.scripts/check.sh
```

## Build and run locally

`geoapi` uses [dep](https://golang.github.io/dep/) to manage dependencies.

```
dep ensure -vendor-only -v
make
# export env variables described above ...
./geoapi
```

## Required environment variables to build Docker image

| Variable                   | Description                                            |
| ---------------------------|:------------------------------------------------------:|
| `GEOAPI_GITLAB_TOKEN`      | The token for download of repositories from gitlab.com |
