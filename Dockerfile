FROM golang:1.20
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /go/src/app/server
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["go", "run", "cmd/url_shortener/main.go"]