#第一构建阶段:基于Golang编译环境的基础镜像，将用于编译下面的微服务模块
FROM golang:1.20 AS builder
#设置环境变量
ARG APP_RELATIVE_PATH
#拷贝微服务文件到指定目录
COPY . /src
#设置当前工作目录
WORKDIR /src/app/${APP_RELATIVE_PATH}
#生成当前微服务的可执行二进制文件
RUN make build
#第二构建阶段:基于debian:stable-slim环境的基础镜像，在这往后的任何内容都将保留在该容器映像中
FROM debian:stable-slim

ARG APP_RELATIVE_PATH
#运行命令 该命令用于更新和安装必须的工具
RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y
#复制当前微服务二进制文件到指定目录，但该文件只存在于前一个 Docker 阶段，这就是使用 --from=builder 复制它的原因
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app
COPY --from=builder /src/config.yaml /app/conf/config.yaml
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/configs/authz/casbin_model.conf /app/conf/authz/casbin_model.conf
#设置新的工作目录
WORKDIR /app


CMD ["./server", "-conf", "/app/conf/config.yaml"]