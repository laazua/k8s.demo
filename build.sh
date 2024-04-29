#!/usr/bin/bash

# 创建output文件夹,并生成打包镜像所需要的文件
if [ ! -d "output" ];then
    mkdir output
fi

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o k8s-demo . && \
    cp k8s-demo config.yaml output && \
    rm k8s-demo

