FROM golang:alpine AS build

COPY . /go/src/github.com/co0p/quotes
WORKDIR /go/src/github.com/cso0p/quotes
RUN GOOS=linux GOARCH=amd64  CGO_ENABLED=0 go build -o quotes

FROM alpine
COPY --from=build /go/src/github.com/co0p/quotes/quotes /bin/quotes
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/quotes