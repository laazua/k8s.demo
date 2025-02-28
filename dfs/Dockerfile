# 第一阶段：构建 Go 项目
FROM docker.io/golang:1.23 AS builder

# 设置工作目录
WORKDIR /app

# 复制 Go 模块文件并下载依赖
COPY . .
RUN go mod tidy

# 构建 Go 项目
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dfs .


#FROM scratch
FROM docker.io/alpine:latest

# 从第一阶段复制构建好的二进制文件到当前镜像
COPY --from=builder /app/dfs /usr/local/bin/dfs

EXPOSE 8022

ENTRYPOINT ["dfs"]
CMD ["-run"]
