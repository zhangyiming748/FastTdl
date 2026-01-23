# FastTdl

FastTdl 是一个高效的Telegram文件下载工具，支持批量下载、自动归档、SQLite防重复下载等功能。

## 功能特点

- 支持Telegram频道文件批量下载
- 支持灵活的URL参数配置（标签、子标签、文件名、偏移量、容量等）
- SQLite数据库记录已下载文件，避免重复下载
- 支持代理设置
- 支持评论区下载功能
- 支持定时提醒功能，监控长时间下载任务
- 自动创建归档目录结构

## 环境要求

- Go 1.19+
- [tdl](https://github.com/iyear/tdl) 工具
- SQLite3

## 快速开始

### 1. 安装 tdl 工具

首先需要安装 [tdl](https://github.com/iyear/tdl) 工具，这是FastTdl的基础依赖：

```bash
go install github.com/iyear/tdl@latest
```

### 2. 配置 tdl 登录

使用 tdl 登录你的Telegram账号：

```bash
tdl login
```

### 3. 准备下载链接文件

创建 `post.link` 文件，每行放置一个下载链接，支持以下格式：

```
https://t.me/${channel}/${fid}
https://t.me/${channel}/${fid}#${tag}
https://t.me/${channel}/${fid}#${tag}&${subtag}
https://t.me/${channel}/${fid}#${tag}&${subtag}@${fname}
https://t.me/${channel}/${fid}%${capacity}
https://t.me/${channel}/${fid}#${tag}+${offset}
```

### 4. 编译和运行

```bash
go build -o fasttdl .
./fasttdl <main_folder> <post_link_file>
```

例如：
```bash
./fasttdl ./downloads post.link
```

## 参数说明

- `main_folder`: 主下载目录，所有文件将下载到此目录
- `post_link_file`: 包含下载链接的文件路径

## URL格式详解

### 基础格式
- `https://t.me/${channel}/${fid}` - 基础下载链接

### 标签格式
- `#` - 主文件夹标签，指定文件下载到哪个主文件夹
- `&` - 子文件夹标签，指定文件下载到哪个子文件夹  
- `@` - 自定义文件名，指定下载后重命名的文件名

### 批量下载格式
- `%${capacity}` - 下载当前文件及之后的n个文件
- `+${offset}` - 从当前文件偏移n位开始下载

### 示例
```
https://t.me/channel_name/12345                    # 基础下载
https://t.me/channel_name/12345#videos            # 下载到 videos 主文件夹
https://t.me/channel_name/12345#videos&anime      # 下载到 videos/anime 子文件夹
https://t.me/channel_name/12345#videos&anime@my_video # 下载并重命名为 my_video
https://t.me/channel_name/12345%10                # 下载当前文件及后续9个文件
https://t.me/channel_name/12345+5                 # 从第5个文件开始下载
```

## 特殊功能

### 评论区下载
当检测到链接中包含 "comment" 关键字时，程序会自动调用讨论区下载功能。

### 自动归档
程序会根据URL中的标签自动创建归档目录结构，并对特定关键词进行转换（如蒂法→最终幻想/蒂法）。

### 重复下载防护
使用SQLite数据库记录已下载的文件，防止重复下载相同的文件。

### 超时提醒
对于长时间下载的任务，程序会每60秒发出一次提醒。

## 配置文件

项目根目录下的配置文件：
- `zh_cn2en_us.md` - 中英文对照表，用于标签转换

## 数据库

SQLite数据库用于存储：
- 已下载文件记录
- 频道信息
- 防止重复下载

## 常见问题

1. **无法下载？**
   - 确认 tdl 已正确安装并登录
   - 检查网络连接和代理设置
   - 确认下载链接格式正确

2. **重复下载问题？**
   - 程序会自动使用SQLite数据库避免重复下载
   - 检查数据库是否正常工作

3. **如何设置代理？**
   - 默认使用 `DEFAULT_PROXY` 常量
   - 可在 `constant` 包中修改代理设置

## 许可证

本项目采用 MIT 许可证

## 致谢

- [tdl](https://github.com/iyear/tdl) - 提供Telegram下载核心功能
