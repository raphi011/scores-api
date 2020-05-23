ARG GO_VERSION=1.14
ARG VERSION=dev

FROM golang:$GO_VERSION as builder
RUN go get -u github.com/gobuffalo/packr/packr
WORKDIR /scores/backend
COPY ./go.* ./
RUN go mod download
COPY . .
WORKDIR /scores/backend/cmd/api
RUN packr
# the portable build tag removes the SQLITE driver that relies on cgo
RUN CGO_ENABLED=0 GOOS=linux go install -ldflags="-X main.version=$VERSION" -tags=portable

FROM alpine:latest
EXPOSE 8080
WORKDIR /scores
COPY --from=builder /go/bin/api .
ENTRYPOINT ["/scores/api"]
