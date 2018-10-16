FROM golang

ARG app_env
ARG backend_url
ENV APP_ENV $app_env
ENV BACKEND_URL $backend_url

WORKDIR /code/scores

COPY . .

WORKDIR /code/scores/cmd/api

RUN go install -ldflags="-X main.version=$(git describe --always --long)"

ENTRYPOINT [ "/go/bin/api" ]
CMD ["-provider", "sqlite3", "-connection", "/srv/scores/scores.db", "-goauth", "/srv/scores/client_secret.json"]

EXPOSE 8080
