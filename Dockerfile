FROM golang:1.19.2 as GOPREP

WORKDIR /app

ADD cmd /app/cmd
ADD internal /app/internal

COPY go.mod /app
COPY go.sum /app




RUN go env -w GOPRIVATE="github.com/byt3-m3"

RUN go mod download


RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 && go build -o /go/bin/zip_fetcher -ldflags="-X 'main.Version=v1.0.0'" /app/cmd/main.go


FROM ubuntu as build

COPY uszips.csv /uszips.csv

COPY --from=GOPREP /go/bin/zip_fetcher /usr/bin/zip_fetcher