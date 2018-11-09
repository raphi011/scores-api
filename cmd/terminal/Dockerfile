FROM golang:1.11.1

WORKDIR /code/scores

RUN go get -u github.com/gobuffalo/packr/packr

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /code/scores/cmd/terminal

RUN packr
RUN go install -ldflags="-X main.version=$(git describe --always --long)"

ENTRYPOINT ["/go/bin/terminal"]
CMD ["-url", "http://backend:8080", "scrape"]
