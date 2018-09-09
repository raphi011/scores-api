FROM golang

ARG app_env
ARG backend_url
ENV APP_ENV $app_env
ENV BACKEND_URL $backend_url

WORKDIR /go/src/github.com/raphi011/scores

RUN go get -u github.com/golang/dep/...

COPY Gopkg.lock .
COPY Gopkg.toml .
RUN dep ensure -vendor-only

COPY . .

WORKDIR /go/src/github.com/raphi011/scores/cmd/api

RUN go install -ldflags="-X main.version=$(git describe --always --long --dirty)"

ENTRYPOINT [ "/go/bin/api" ]
CMD ["-provider", "sqlite3", "-connection", "/srv/scores/scores.db", "-goauth", "/srv/scores/client_secret.json"]

EXPOSE 8080
