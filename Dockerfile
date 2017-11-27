FROM arm32v7/golang:1.9.2

WORKDIR /go/src/scores-backend
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]

EXPOSE 8080
