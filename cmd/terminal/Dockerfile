FROM golang

WORKDIR /go/src/github.com/raphi011/scores

RUN go get -u github.com/golang/dep/...
COPY Gopkg* ./
RUN dep ensure

COPY . .

WORKDIR /go/src/github.com/raphi011/scores/cmd/terminal

RUN go install -ldflags="-X main.version=$(git describe --always --long --dirty)"

ENTRYPOINT ["/go/bin/terminal"]
CMD ["-url", "http://backend:8080", "scrape"]
