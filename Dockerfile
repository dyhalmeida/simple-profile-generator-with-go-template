FROM golang:1.22.2-alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM scratch
COPY --from=builder /app/main /app/main
COPY --from=builder /app/templates /templates
EXPOSE 3333
ENTRYPOINT ["/app/main"]