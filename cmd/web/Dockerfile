FROM golang

ARG app_env
ARG backend_url
ENV APP_ENV $app_env
ENV BACKEND_URL $backend_url

WORKDIR /go/src/github.com/raphi011/scores

RUN go get "github.com/gin-contrib/sessions"
RUN go get "github.com/gin-gonic/gin"
RUN go get "github.com/mattn/go-sqlite3"
RUN go get "golang.org/x/oauth2"
RUN go get "golang.org/x/oauth2/google"
RUN go get "golang.org/x/crypto/pbkdf2"
RUN go get "github.com/PuerkitoBio/goquery"

COPY . .

WORKDIR /go/src/github.com/raphi011/scores/cmd/web
RUN go install

ENTRYPOINT [ "/go/bin/web" ]
CMD ["-db", "/srv/scores/scores.db", "-goauth", "/srv/scores/client_secret.json"]

EXPOSE 8080
