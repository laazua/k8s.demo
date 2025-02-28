### dockerfile

- **ADD**
1. 支持从远程 HTTPS 和 Git URL 获取文件保存到容器中
2. 构建上下文添加文件时自动提取 tar 文件

- **COPY**
1. 复制 文件保存到容器中
2. 支持多阶段构建

- **CMD**
1. 为容器提供默认的命令和参数,但可以被 docker run 命令传递的参数覆盖

- **ENTRYPOINT**
1. 定义了容器启动时执行的命令,并且不能被覆盖
2. 如果ENTRYPOINT和CMD同时存在,则CMD中的参数时传递给ENTRYPOINT中的参数

- **示例**
1. 以当前目录下的项目为例
2. 打包镜像: podman build -t dfs:v1 .
3. 运行命令: podman run -it --name dfs -p8022:8022 dfs:v1 -help  
   命令行参数-help 会传递给ENTRYPOINT中作为dfs的参数
4. 运行命令: podman run -it --name dfs -p8022:8022 dfs:v1 -version  
   命令行参数-version 会传递给ENTRYPOINT中作为dfs的参数
5. 运行命令: podman run -it --name dfs -p8022:8022 dfs:v1  
   此时命令行没有传入参数,则容器运行的最终命令为ENTRYPOINT+CMD,其中CMD  
   中的参数将作为ENTRYPOINT的命令参数
