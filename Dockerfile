FROM golang:1.14.4 as mod
LABEL stage=mod
ARG GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct
WORKDIR /root/myapp/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

FROM mod as builder
LABEL stage=intermediate0
ARG LDFLAGS
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rabbitmq_exporter -ldflags "${LDFLAGS}" main.go


FROM alpine:3.11.6
WORKDIR /root
COPY --from=builder /root/myapp/rabbitmq_exporter /rabbitmq_exporter
ENTRYPOINT ["/rabbitmq_exporter"]
