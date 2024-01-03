FROM golang:1.21 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o /server main.go

FROM gcr.io/distroless/base-debian11 as final

COPY --from=builder /server /server

ENV PORT 8080
EXPOSE $PORT

ENTRYPOINT ["/server"]
