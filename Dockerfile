FROM golang:1.23.3-bookworm
#docker run -it --rm --name tdl golang:1.23.2-bookworm bash
#docker exec -it tdl bash
LABEL authors="zen"

# 设置非交互模式
ENV DEBIAN_FRONTEND=noninteractive

# 更换完整源
COPY debian.sources /etc/apt/sources.list.d/debian.sources

# 更新软件包并安装依赖
RUN apt update && \
    apt install -y --no-install-recommends locales \
    wget curl build-essential  openssh-server nano git && \
    apt clean && \
    rm -rf /var/lib/apt/lists/*

# 配置 Go 环境

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go env -w GOBIN=/go/bin

# 复制文件
RUN git clone https://github.com/iyear/tdl.git /root/tdl
WORKDIR /root/tdl
RUN go build -o /usr/local/bin/tdl

# 中文支持
RUN apt update && \
    apt install -y --no-install-recommends locales && \
    echo "zh_CN.UTF-8 UTF-8" >> /etc/locale.gen && \
    locale-gen zh_CN.UTF-8 && \
    update-locale LANG=zh_CN.UTF-8

# 设置环境变量
ENV LANG=zh_CN.UTF-8 \
    LANGUAGE=zh_CN:zh \
    LC_ALL=zh_CN.UTF-8

# 设置 root 密码
RUN echo "root:123456" | chpasswd

# 允许 root 登录 SSH
RUN echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config && \
    echo 'PasswordAuthentication yes' >> /etc/ssh/sshd_config

# 设置其他环境变量
ENV PATH="$PATH:/usr/local/go/bin"

# 天朝特色：更换源
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources

# # 启动 SSH 服务
# WORKDIR /
# ENTRYPOINT ["service", "ssh", "start", "-D"]

# 复制文件
RUN git clone https://github.com/zhangyiming748/FastTdl.git /root/fastTDL
WORKDIR /root/fastTDL
RUN go build -o /usr/local/bin/fastTDL /root/fastTDL/main.go


# 启动程序
WORKDIR /data
ENTRYPOINT ["fastTDL"]