# 第一阶段：使用 Golang 1.23.4 编译tdl主程序

FROM golang:1.24.0-alpine3.21 AS builderTDL

LABEL authors="zen"

COPY origin_tdl /root/tdl

WORKDIR /root/tdl

RUN go build -o /usr/local/bin/tdl main.go

# 第二阶段：使用 Golang 1.23.4 编译fastTDL主程序

FROM golang:1.24.0-alpine3.21 AS builderFastTDL

LABEL authors="zen"

COPY FastTdl /root/FastTdl

WORKDIR /root/FastTdl

RUN go build -o /usr/local/bin/fastTDL main.go

# 第三阶段：使用alpine3.21作为基础镜像

FROM alpine:3.21

LABEL authors="zen"

COPY --from=builderTDL /usr/local/bin/tdl /usr/local/bin/tdl
COPY --from=builderFastTDL /usr/local/bin/fastTDL /usr/local/bin/fastTDL

# 中文支持

RUN apk add  ffmpeg nano 

# 设置环境变量

ENV LANG=zh_CN.UTF-8 \
    LANGUAGE=zh_CN:zh \
    LC_ALL=zh_CN.UTF-8

# 启动程序
WORKDIR /data
ENTRYPOINT ["fastTDL"]