FROM golang:1.15 as builder
MAINTAINER milkomeda.org "admin@milkomeda.org"

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .

# CGO_ENABLED禁用cgo 然后指定OS等，并go build
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o pi content/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/pi /app

# 指定运行时环境变量
ENV PORT=80
EXPOSE 80

ENTRYPOINT ["/app/pi"]