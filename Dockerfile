FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o math .

FROM scratch

COPY --from=builder /app/math .

CMD ["./math"]
