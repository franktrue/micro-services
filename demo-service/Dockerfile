# 使用 golang 官方镜像提供 Go 运行环境，并且命名为 buidler 以便后续引用
FROM golang:1.16-alpine as builder
# 启用 Go Modules 并设置 GOPROXY
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装 git
RUN apk --no-cache add git

# 设置工作目录
WORKDIR /app/demo-service
# 将当前项目所在目录代码拷贝到镜像中
COPY . .

# 下载依赖
RUN go mod download

# 构建二进制文件，添加来一些额外参数以便可以在 Alpine 中运行它
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o demo-service

# 下面是第二阶段的镜像构建，和之前保持一致
FROM alpine:latest

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update && apk add --no-cache bash supervisor
RUN mkdir /app
WORKDIR /app
# 服务二进制可执行文件
# 这一步不再从宿主机拷贝二进制文件，而是从上一个阶段构建的 builder 容器中拉取
COPY --from=builder /app/demo-service/demo-service .

ADD supervisord.conf /etc/supervisord.conf

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
