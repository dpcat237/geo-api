# Create builder image
FROM golang:alpine as builder
ARG GEOAPI_GITLAB_TOKEN

WORKDIR /go/src/gitlab.com/dpcat237/geoapi

# Download dependencies
RUN apk update && apk upgrade && apk add git
RUN git config --global url."http://dpcat237:${GEOAPI_GITLAB_TOKEN}@gitlab.com/".insteadOf "https://gitlab.com/"
RUN go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only -v

# Build the binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/geoapi

# Create final image
FROM alpine
EXPOSE 3000 5000
COPY --from=builder /go/bin/geoapi /go/bin/geoapi
RUN addgroup usgeoapi && adduser -S -G usgeoapi usgeoapi
USER usgeoapi
ENTRYPOINT ["/go/bin/geoapi"]
