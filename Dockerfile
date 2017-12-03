FROM arm32v7/golang:1.9.2

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/src/scores-backend

RUN go get "github.com/gin-contrib/sessions"
RUN go get "github.com/gin-gonic/gin"
RUN go get "github.com/jinzhu/gorm/dialects/sqlite"
RUN go get "github.com/jinzhu/gorm"
RUN go get "golang.org/x/oauth2"
RUN go get "golang.org/x/oauth2/google"

COPY . .

# RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]

EXPOSE 8080
