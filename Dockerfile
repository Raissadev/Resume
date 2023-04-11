FROM golang:1.19-alpine AS build

LABEL project="42"

ENV PATH="$PATH:$(go env GOPATH)/bin"
ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build
ENV GOOS linux

RUN apk update && \
	apk add ca-certificates && \
	apk add tzdata

WORKDIR /go/src

COPY ./backend ./
COPY ./frontend/build /public

RUN go mod tidy

RUN go build -o ./bin/application

FROM alpine:latest as finally

WORKDIR /app

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /go/src/bin/application ./
COPY ./frontend/build /public

ENTRYPOINT ["/app/application"]