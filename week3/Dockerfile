FROM golang:1.16 AS builder
COPY . /src
WORKDIR /src
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server *.go

FROM alpine
COPY --from=builder /src/server /app/server
WORKDIR /app
EXPOSE 80
CMD ["./server"]
