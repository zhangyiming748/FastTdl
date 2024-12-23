#!/bin/bash
# 文件名的格式如下
# 1746575178_116770_aaaaa.mp4
# 1746575178_116770_啊啊啊啊啊

# 1. 利用find命令查找并打印全部符合条件的文件的绝对路径 
# 2. 利用sed命令重命名文件保留从第一个不是数字或下划线的文字内容到结尾为止

# 比如 文件名

# `1746575178_116770_aaaaa.mp4`重命名之后是`1746575178_116770_aaaaa.mp4`
# `1746575178_116770_啊啊啊啊啊`重命名之后是`1746575178_116770_啊啊啊啊啊`

# shell脚本实现
# dir=/mnt/c/Users/zen/Downloads/media/lice
# 查找符合条件的文件并打印绝对路径
find ${dir} -type f -name '*_*' | while read -r file; do
    # 获取文件名
    filename=$(basename "$file")
    
    # 使用 sed 提取从第一个不是数字或下划线的字符开始到结尾的部分
    new_filename=$(echo "$filename" | sed -E 's/^[0-9_]+(.*)/\1/')
    
    # 如果 new_filename 不为空，进行重命名
    if [[ -n "$new_filename" ]]; then
        # 获取文件的目录
        dir=$(dirname "$file")
        
        # 重命名文件
        mv "$file" "$dir/$new_filename"
        
        # 打印重命名后的文件路径
        echo "Renamed: $file -> $dir/$new_filename"
    fi
done
