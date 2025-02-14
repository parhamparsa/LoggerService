FROM golang:1.22-alpine as builder

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o talon-backend-assignment ./cmd

FROM scratch

ENV HTTP_PORT=8080
EXPOSE $HTTP_PORT

COPY --from=builder /app/talon-backend-assignment .
ENTRYPOINT [ "./talon-backend-assignment" ]
