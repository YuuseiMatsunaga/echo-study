FROM golang:latest

WORKDIR /app
COPY ./ /app

RUN go install github.com/cosmtrek/air@latest \
    && go get github.com/rubenv/sql-migrate/... \
    && go install github.com/rubenv/sql-migrate/...

ENV CGO_ENABLED=0 \
    GOOS=linux

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]