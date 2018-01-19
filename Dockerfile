FROM arm32v7/golang:1.9.2

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/src/scores-backend/

RUN go get "github.com/gin-contrib/sessions"
RUN go get "github.com/gin-gonic/gin"
RUN go get "github.com/mattn/go-sqlite3"
RUN go get "golang.org/x/oauth2"
RUN go get "golang.org/x/oauth2/google"

COPY . .

# RUN go-wrapper download

WORKDIR /go/src/scores-backend/cmd/web
RUN go-wrapper install

CMD ["go-wrapper", "run", "-db", "/srv/scores/scoresdb", "-goauth", "/srv/scores/client_secret.json"]

EXPOSE 8080
