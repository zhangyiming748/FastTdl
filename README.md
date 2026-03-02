# FastTdl

FastTdl 是一个高效的 Telegram 文件下载和归档工具，基于 Go 语言开发，集成了批量下载、智能归档、重复下载防护、评论区下载等核心功能。

## 🌟 核心功能

- **批量下载**: 支持 Telegram 频道文件批量下载
- **智能归档**: 自动创建目录结构，支持标签分类
- **重复防护**: SQLite 数据库记录已下载文件，避免重复下载
- **评论区支持**: 自动检测并下载评论区内容
- **代理支持**: 内置代理配置，适应不同网络环境
- **定时提醒**: 长时间下载任务自动提醒
- **格式转换**: 视频转 H.265，图片转 AVIF 优化存储空间
- **视频旋转**: 支持批量旋转视频文件（90°、180°、270°）
- **CLI 命令行**: 基于 Cobra 的现代化命令行界面

## 🚀 环境要求

- Go 1.26+
- [tdl](https://github.com/iyear/tdl) 工具
- SQLite3

## ⚡ 快速开始

### 1. 安装依赖

```bash
# 安装 tdl 工具（FastTdl 的核心依赖）
go install github.com/iyear/tdl@latest

# 克隆项目
git clone https://github.com/zen/Github/tdl-cli.git
cd tdl-cli

# 安装 Go 依赖
go mod download
```

### 2. 配置 Telegram 登录

```bash
# 登录 Telegram 账号
tdl login
```

### 3. 准备下载链接文件

创建 `post.link` 文件，每行一个下载链接：

```txt
https://t.me/channel_name/12345
https://t.me/channel_name/12345#videos
https://t.me/channel_name/12345#videos&anime
https://t.me/channel_name/12345#videos&anime@custom_name
https://t.me/channel_name/12345%10
https://t.me/channel_name/12345+5
```

### 4. 编译和使用

```bash
# 编译项目
go build -o fasttdl .

# 查看帮助
./fasttdl --help

# 下载文件
./fasttdl tdl --root "./downloads" --postlink "./post.link" --proxy "http://127.0.0.1:8889"

# 归档文件
./fasttdl archive --root "./downloads"
```

## 📖 命令行参数

### 下载命令 (`tdl`)

```bash
./fasttdl tdl --root <目录> --postlink <链接文件> [--proxy <代理地址>]
```

参数说明：

- `--root`: 主下载目录（必需）
- `--postlink`: 包含下载链接的文件路径（必需）
- `--proxy`: 代理地址（可选，默认 `http://127.0.0.1:8889`）

### 归档命令 (`archive`)

```bash
./fasttdl archive --root <目录>
```

参数说明：

- `--root`: 要归档的目录路径（必需）

### 视频旋转命令 (`rotate`)

```bash
./fasttdl rotate --root <目录> --direction <角度>
```

参数说明：

- `--root`: 要旋转视频的目录路径（必需）
- `--direction`: 旋转方向，支持 90、180、270 度（必需）

使用示例：

```bash
# 旋转目录下所有视频文件90度
./fasttdl rotate --root "./videos" --direction "90"

# 旋转目录下所有视频文件180度
./fasttdl rotate --root "./videos" --direction "180"

# 旋转目录下所有视频文件270度
./fasttdl rotate --root "./videos" --direction "270"
```

## 🔗 URL 格式详解

### 基础语法

| 符号 | 功能 | 示例 |
| :---: | :---: | :---: |
| `#` | 主文件夹标签 | `#videos` → 下载到 videos 目录 |
| `&` | 子文件夹标签 | `#videos&anime` → 下载到 videos/anime 目录 |
| `@` | 自定义文件名 | `@my_video.mp4` → 下载并重命名为指定名称 |
| `%` | 批量下载容量 | `%10` → 下载当前文件及后续9个文件 |
| `+` | 偏移量 | `+5` → 从第5个文件开始下载 |

### 完整示例

```bash
# 基础下载
https://t.me/channel_name/12345

# 指定主目录
https://t.me/channel_name/12345#videos

# 指定子目录
https://t.me/channel_name/12345#videos&anime

# 自定义文件名
https://t.me/channel_name/12345#videos&anime@episode1.mp4

# 批量下载10个文件
https://t.me/channel_name/12345%10

# 从第5个文件开始下载
https://t.me/channel_name/12345+5

# 组合使用
https://t.me/channel_name/12345#videos&anime@episode1.mp4%5
```

## 🛠️ 高级功能

### 评论区下载

当链接中包含 "comment" 关键字时，自动启用评论区下载模式：

```bash
https://t.me/channel_name/12345?comment=67890
```

### 标签自动转换

支持中文标签自动转换为标准目录结构（配置文件 `zh_cn2en_us.md`）：

```txt
蒂法 → 最终幻想/蒂法
鸣人 → 火影忍者/鸣人
```

### 视频和图片优化

归档功能支持自动格式转换：

```bash
# 视频转 H.265 编码优化存储
./fasttdl archive --root "./videos"

# 图片转 AVIF 格式优化存储
./fasttdl archive --root "./images"
```

### 视频旋转功能

新增的视频旋转功能支持批量处理目录中的所有视频文件：

```bash
# 旋转视频文件
./fasttdl rotate --root "./videos" --direction "90"

# 支持的角度：
# 90  - 顺时针旋转90度
# 180 - 旋转180度
# 270 - 逆时针旋转90度（或顺时针270度）
```

该功能特点：

- 支持多种视频格式
- 保持原有视频质量和编码
- 批量处理整个目录
- 处理过程中显示进度信息

## 📁 项目结构

```shell
tdl-cli/
├── archive/          # 归档处理模块
├── constant/         # 常量和配置
├── core/            # 核心业务逻辑
├── discussions/     # 讨论区下载
├── model/           # 数据模型
├── sqlite/          # 数据库操作
├── tdl/             # tdl 工具封装
├── util/            # 工具函数
├── main.go          # 程序入口
├── post.link        # 下载链接文件示例
└── zh_cn2en_us.md   # 标签转换配置
```

## ⚙️ 配置说明

### 代理设置

默认代理地址：`http://127.0.0.1:8889`

可在 `constant/param.go` 中修改：

```go
const DEFAULT_PROXY = "http://127.0.0.1:8889"
```

### 日志配置

日志文件：`tdl.log`

日志轮转脚本：

- Linux/macOS: `./rotate.sh`
- Windows: `powershell -ExecutionPolicy Bypass -File rotate.ps1`

## 🔧 开发指南

### 本地开发

```bash
# 获取依赖
go mod tidy

# 运行测试
go test ./...

# 构建
go build -o fasttdl .
```

### 代码结构

主要模块说明：

- `core/tdl.go`: 核心下载逻辑
- `archive/archive.go`: 文件归档处理
- `rotate/rotate.go`: 视频旋转功能
- `discussions/discussions.go`: 评论区下载
- `sqlite/sqlite.go`: 数据库操作
- `util/`: 各种工具函数

## 🐛 常见问题

### 下载失败

1. 确认 tdl 工具已正确安装并登录
2. 检查网络连接和代理设置
3. 验证下载链接格式是否正确

### 重复下载

程序自动使用 SQLite 数据库防止重复下载：

- 数据库存储位置：程序同目录下
- 表结构：channels 和 files 表

### 代理问题

```bash
# 测试代理连通性
curl -x http://127.0.0.1:8889 http://www.google.com

# 或在代码中修改代理地址
const DEFAULT_PROXY = "your_proxy_address"
```

## 📊 性能优化

### 并发控制

- 视频处理：最大并发数 1
- 图片处理：最大并发数 1
- 下载线程：tdl 默认 8 线程

### 存储优化

归档功能可显著减少存储空间：

- 视频：H.265 编码平均节省 30-50% 空间
- 图片：AVIF 格式平均节省 20-40% 空间

## 🤝 贡献指南

欢迎提交 Issue 和 Pull Request！

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 发起 Pull Request

## 📄 许可证

本项目采用 MIT 许可证

## 🙏 致谢

- [tdl](https://github.com/iyear/tdl) - 提供 Telegram 下载核心功能
- [Cobra](https://github.com/spf13/cobra) - 命令行框架
- [GORM](https://gorm.io/) - ORM 框架
- [SQLite](https://www.sqlite.org/) - 轻量级数据库

---

**注意**: 请确保遵守 Telegram 的使用条款和服务协议。
