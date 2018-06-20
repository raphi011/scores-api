FROM golang

WORKDIR /go/src/github.com/raphi011/scores

RUN go get "github.com/mattn/go-sqlite3"
RUN go get "golang.org/x/crypto/pbkdf2"
RUN go get "github.com/PuerkitoBio/goquery"

COPY . .

WORKDIR /go/src/github.com/raphi011/scores/cmd/terminal
RUN go install

ENTRYPOINT ["/go/bin/terminal"]
CMD ["-url", "http://backend:8080", "scrape"]
