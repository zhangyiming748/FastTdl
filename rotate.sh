#!/bin/bash

# 指定要查找的扩展名
extension="mp4"
root="/Users/zen/github/FastTdl/dance"

process_files() {
    local dir=$1
    local transpose_value=$2

    # 遍历指定目录下的所有文件
    for file in "$dir"/*.$extension; do
        # 检查文件是否具有指定的扩展名
        if [[ -f $file ]]; then
            # 提取文件名和扩展名
            filename=$(basename "$file")
            basename_without_ext="${filename%.*}"
            aftername="${basename_without_ext}_rotate.mp4"

            # 打印分开的文件名和扩展名
            echo "文件名： $basename_without_ext"
            echo "扩展名： $extension"

            # 执行ffmpeg命令
            ffmpeg -i "$file" -vf "transpose=$transpose_value" -c:v libx265 -c:a libmp3lame -tag:v hvc1 -map_chapters -1 "$dir/$aftername"
            # ffmpeg -i "$file" -vf "transpose=$transpose_value" -c:v h264_nvenc -preset medium -c:a copy -map_chapters -1 "$dir/$aftername"
            if [ $? -eq 0 ]; then
                echo "旧文件名: $filename 新文件名: $aftername"
                echo "done"
                rm "$file"
            else
                echo "ffmpeg command failed for file: $file"
            fi
            echo "done"
        fi
    done
}

# 处理右侧和左侧的文件
process_files "$root/toRight" 1
process_files "$root/toLeft" 2
