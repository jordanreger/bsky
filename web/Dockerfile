ARG GO_VERSION=1
FROM golang:${GO_VERSION}-alpine as builder

# fix x509 cert error
RUN apk update && apk add ca-certificates

# clear cache
RUN go clean -modcache; cd bsky && go get -u all; cd ..; cd web && go get -u all; cd ..;

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .

FROM alpine:latest

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
