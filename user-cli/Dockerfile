FROM golang:1.16-alpine as builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk --no-cache add git

WORKDIR /app/user-cli
COPY . .

# 下载依赖
RUN go mod download

# 构建二进制文件，添加来一些额外参数以便可以在 Alpine 中运行它
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-cli


FROM alpine:latest

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add --no-cache bash supervisor
RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/user-cli/user-cli .

# ENTRYPOINT支持额外参数CMD不支持
ENTRYPOINT [ "./user-cli" ]