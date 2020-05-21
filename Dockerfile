ARG GO_VERSION=1.14
ARG VERSION=dev

FROM golang:$GO_VERSION

RUN go get -u github.com/gobuffalo/packr/packr

WORKDIR /scores/backend

COPY ./go.* ./

RUN go mod download

COPY . .

WORKDIR /scores/backend/cmd/api

RUN packr

RUN go install -ldflags="-X main.version=$VERSION"

ENTRYPOINT [ "/go/bin/api" ]

EXPOSE 8080
