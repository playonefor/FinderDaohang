FROM golang:1.20 as builder

ENV GOPROXY=https://goproxy.io,direct

WORKDIR /app/FinderDaohang

COPY . .

RUN go mod init FinderDaohang

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o finderdaohang .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app/FinderDaohang

COPY --from=builder /app/FinderDaohang .

CMD ["./finderdaohang"]
