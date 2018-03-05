FROM arm32v7/golang:latest

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/src/github.com/raphi011/scores

RUN go get "github.com/gin-contrib/sessions"
RUN go get "github.com/gin-gonic/gin"
RUN go get "github.com/mattn/go-sqlite3"
RUN go get "golang.org/x/oauth2"
RUN go get "golang.org/x/oauth2/google"
RUN go get "golang.org/x/crypto/pbkdf2"

COPY . .

WORKDIR /go/src/github.com/raphi011/scores/cmd/web
RUN go-wrapper install

CMD ["go-wrapper", "run", "-db", "/srv/scores/scores.db", "-goauth", "/srv/scores/client_secret.json"]

EXPOSE 8080
