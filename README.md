# FastTdl

一个基于tdl的高效Telegram文件下载工具，支持批量下载、自动重试、多线程下载等功能。

## 功能特点

- 支持Telegram频道和群组文件的批量下载
- 支持多线程并发下载，显著提升下载速度
- 自动重试机制，确保下载稳定性
- Docker容器化部署，使用方便
- 支持代理设置
- 中文友好

## 环境要求

- Docker环境
- 或Go 1.24+（本地编译运行）

## 快速开始

### Docker方式（推荐）

1. 克隆项目
```bash
git clone https://github.com/yourusername/FastTdl.git
cd FastTdl
```

2. 配置docker-compose.yml
```yaml
version: '3.3'
services:
  fastTDL:
    image: yourimage/fasttdl:latest
    container_name: fastTDL
    volumes:
      - /path/to/your/.tdl:/root/.tdl  # 替换为你的.tdl目录
      - /path/to/download:/data        # 替换为你的下载目录
    restart: always
```

3. 启动服务
```bash
docker-compose up -d
```

### 本地编译方式

1. 克隆并编译
```bash
git clone https://github.com/yourusername/FastTdl.git
cd FastTdl
go build -o fastTDL main.go
```

2. 运行程序
```bash
./fastTDL
```

## 使用方法

### 1. 导出消息

```bash
tdl chat export -c <频道ID或用户名> --reply <消息ID> [--proxy <代理地址>]
```

示例：
```bash
tdl chat export -c channelname --reply 12345 --proxy http://127.0.0.1:7890
```

### 2. 下载文件

```bash
tdl download --threads <线程数> --file tdl-export.json [--proxy <代理地址>]
```

示例：
```bash
tdl download --threads 8 --file tdl-export.json --proxy http://127.0.0.1:7890
```

## 消息ID获取方法

1. 频道ID获取：
   - 在Telegram设置中启用开发者模式（设置 -> 高级 -> 开发者选项）
   - 点击频道信息查看频道ID

2. 消息ID获取：
   - 消息链接格式：https://t.me/c/channelid/messageid
   - 或 https://t.me/channelname/messageid

## 常见问题

1. **下载速度慢？**
   - 尝试增加线程数
   - 检查网络连接和代理设置
   - 确保服务器带宽充足

2. **下载失败？**
   - 检查TDL登录状态
   - 确认消息ID正确
   - 验证代理配置

## 许可证

本项目采用 MIT 许可证

## 致谢

- [tdl](https://github.com/iyear/tdl) - 核心下载引擎

以上是Claude自动生成的内容
----

# [项目名称] README

## 一、项目概述
本项目旨在帮助你通过便捷的方式实现特定文件的下载功能，以下是具体的操作步骤及相关配置说明。

## 二、前期准备
### 获取 Telegram 账号登陆状态文件
你需要先通过 [tdl](https://github.com/iyear/tdl) 项目登陆一次 Telegram 账号，这样会生成对应的登陆状态文件。该文件保存的位置因操作系统不同而有所差异：
- **Windows 系统**：通常会保存在你的个人文件夹下的 `.tdl` 文件夹中，具体路径为 `%USERPROFILE%\.tdl`。例如，如果你的用户名是 `John`，那么这个路径可能就是 `C:\Users\John\.tdl`。
- **Linux 和 macOS 系统**：会保存在 `$HOME/.tdl` 目录下，例如在 Linux 系统中，如果你的用户名是 `user`，一般路径就是 `/home/user/.tdl`。

在后续操作中，我们将上述不同系统下对应的登陆状态文件所在位置统一简称为“登陆状态”。

### 准备下载链接文件
你需要将想要下载的链接按行逐一写入到项目所在文件夹下新建的 `post.link` 文件中。具体操作如下：
1. 找到本项目所在的文件夹（你可以通过查看项目的存放路径来确定）。
2. 在该文件夹内，使用文本编辑器（如 Windows 下的记事本、Linux 下的 Vim 或 macOS 下的 TextEdit 等）新建一个名为 `post.link` 的文件。
3. 将你收集好的需要下载的链接，每个链接单独占一行，依次写入到这个 `post.link` 文件当中。

## 三、配置 `docker-compose.yml` 文件
在项目目录下找到 `docker-compose.yml` 文件，需要对其进行如下关键配置修改：
1. **指定登陆状态文件路径**：找到文件中对应的配置位置，将 `:/root/.tdl` 冒号前半部分填写为你之前获取到的登陆状态所在的实际路径。例如，在 Windows 系统下按照上述路径获取到的登陆状态路径为 `C:\Users\John\.tdl`，那么此处就应填写 `C:\Users\John\.tdl`；在 Linux 系统下如果是 `/home/user/.tdl`，则填写 `/home/user/.tdl`。这样做是为了让容器能够正确读取到登陆状态信息，确保下载过程的顺利进行。
2. **指定项目所在位置**：对于 `:/data` 冒号前半部分，需要填写本项目所在的完整位置。同样以实际所在的文件夹路径为准，比如项目存放在 `D:\myproject`（Windows 系统下示例），那就填写 `D:\myproject`；在 Linux 系统下如果存放在 `/opt/project`，则填写 `/opt/project`。这一步的目的是让容器知晓在哪里去查找相关的配置文件以及后续存放下载文件的位置基础设定。
3. **配置代理信息（如果需要）**：找到 `PROXY` 相关配置项，将其修改为宿主机上实际的 IP 地址和端口号。这一步主要是针对需要通过代理服务器进行网络访问的情况，确保容器内的下载操作能够通过正确的代理连接到目标资源。

## 四、执行下载操作
完成上述 `docker-compose.yml` 文件的配置修改后，在命令行或终端中进入到该 `docker-compose.yml` 文件所在的目录（即项目目录），然后执行 `docker-compose` 命令（具体的命令语法可能因 `docker-compose` 版本等因素略有不同，常见的如 `docker-compose up -d` 用于在后台启动相关服务等），开始下载操作。下载完成后，文件会自动保存在项目所在文件夹的同级目录下，你可以方便地到该位置去查看和使用下载好的文件。

# 支持的格式

```
channle = 频道id
fid = 文件id
tag = 归档的文件夹名
subtag = 归档的子文件夹名
fname = 下载后重命名的正式文件名(不包含扩展名)
capacity = 向下下载n个文件
offset = 下载之后的第n个文件
```

```
https://t.me/${channel}/${fid}
https://t.me/${channel}/${fid}#${tag}
https://t.me/${channel}/${fid}#${tag}&${subtag}
https://t.me/${channel}/${fid}#${tag}&${subtag}@${fname}

https://t.me/${channel}/${fid}%${capacity}
https://t.me/${channel}/${fid}#${tag}%${capacity}
https://t.me/${channel}/${fid}#${tag}&${subtag}%${capacity}
https://t.me/${channel}/${fid}#${tag}&${subtag}@${fname}%${capacity}

https://t.me/${channel}/${fid}+${offset}
https://t.me/${channel}/${fid}#${tag}+${offset}
https://t.me/${channel}/${fid}#${tag}&${subtag}+${offset}
https://t.me/${channel}/${fid}#${tag}&${subtag}@${fname}+${offset}
```


