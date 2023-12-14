FROM golang:1.20.5-bullseye as builder

ENV GOPROXY=https://goproxy.io,direct

WORKDIR /app/FinderDaohang

COPY . .

RUN CGO_CFLAGS="-g -O2 -Wno-return-local-addr" CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o finderdaohang .

CMD ./finderdaohang
