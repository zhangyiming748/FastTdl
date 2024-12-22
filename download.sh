#!/bin/bash

# 读取网址列表文件
url_file="urls.txt"

# 检查文件是否存在
if [[ ! -f "$url_file" ]]; then
    echo "网址列表文件 $url_file 不存在."
    exit 1
fi

# 逐行读取网址并执行命令
while IFS= read -r url; do
    # 确保网址不为空
    if [[ -n "$url" ]]; then
        echo "正在下载: $url"
        /Users/zen/Applications/tdl download --proxy http://192.168.1.2:8889 --threads 8 --url "$url"
    fi
done < "$url_file"
