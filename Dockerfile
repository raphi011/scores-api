FROM arm32v7/golang:1.9.2

WORKDIR /backend

ADD . /go/src/scores-backend

RUN go install scores-backend

ENTRYPOINT /go/bin/scores-backend

EXPOSE 8080
