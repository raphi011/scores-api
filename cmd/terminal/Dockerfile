FROM golang

WORKDIR /go/src/github.com/raphi011/scores

COPY . .

WORKDIR /go/src/github.com/raphi011/scores/cmd/terminal

RUN go install -ldflags="-X main.version=$(git describe --always --long)"

ENTRYPOINT ["/go/bin/terminal"]
CMD ["-url", "http://backend:8080", "scrape"]
