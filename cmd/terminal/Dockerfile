FROM golang

WORKDIR /code/scores

COPY . .

WORKDIR /code/scores/cmd/terminal

RUN go install -ldflags="-X main.version=$(git describe --always --long)"

ENTRYPOINT ["/go/bin/terminal"]
CMD ["-url", "http://backend:8080", "scrape"]
