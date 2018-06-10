FROM golang:latest

WORKDIR /go/src/github.com/raphi011/scores

RUN go get "github.com/mattn/go-sqlite3"
RUN go get "golang.org/x/crypto/pbkdf2"

COPY . .

WORKDIR /go/src/github.com/raphi011/scores/cmd/terminal
RUN go install

ENTRYPOINT ["/go/bin/terminal"]
CMD ["-url", "http://backend:8080", "scrape"]
