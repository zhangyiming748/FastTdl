#!/bin/zsh

# 定义包含 URL 的文件名
link_file="post.link"
proxy="http://127.0.0.1:8889"
# 检查文件是否存在
if [ ! -f "$link_file" ]; then
  echo "错误: 文件 '$link_file' 不存在。"
  exit 1
fi

# 使用 for 循环逐行读取文件内容
for url in $(cat "$link_file"); do
  # 移除 URL 前后可能的空白字符 (可选，如果需要)
  url=$(echo "$url" | tr -d '[:space:]')

  # 检查 URL 是否为空或注释行 (可选，如果需要)
  if [ -z "$url" ] || [[ "$url" == \#* ]]; then
    continue  # 跳过空行和注释行
  fi

  # 打印 URL
  echo "正在处理 URL: $url"

  # 在这里添加你的 URL 处理逻辑
  #  示例：打印 URL 长度
  echo "URL长度: ${#url}"
  tdl download --url $url --proxy $proxy

  echo "------------------------" # 分隔符
done

echo "处理完成。"
exit 0

