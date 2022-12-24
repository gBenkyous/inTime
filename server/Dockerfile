# 開発用
FROM golang:1.19-alpine AS dev
RUN apk add --no-cache git make curl

# ビルド用
FROM golang:1.19-alpine AS builder
COPY . /go/src
WORKDIR /go/src
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o /bin/app-server ./cmd/intime/main.go 

# 本番配置用
FROM scratch AS prod
COPY --from=builder /bin/app-server /bin/app-server
ENTRYPOINT ["/bin/app-server"]
