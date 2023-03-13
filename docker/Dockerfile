FROM golang:1.19 AS build

COPY . /app/
WORKDIR /app/

RUN go get ./... \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/app/

    
CMD ["/app/app"]
