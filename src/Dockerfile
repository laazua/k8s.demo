# 使用Alpine Linux 作为基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 将二进制文件从编译环境中复制到Alpine镜像
COPY output /app/

# 暴露应用程序运行的端口（如果有需要）
EXPOSE 8889

# 运行可执行文件
CMD ["./k8s-demo"]

# 命令行打包二进制文件: 
# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o k8s-demo .