#!/bin/zsh

# 指定要查找的扩展名
extension="mp4"
root="/Users/zen/Downloads/media/dance"

process_files() {
    local dir=$1
    local transpose_value=$2

    # 遍历指定目录下的所有文件
    for file in "$dir"/*.$extension; do
        # 检查文件是否具有指定的扩展名
        if [[ -f $file ]]; then
            # 提取文件名和扩展名
            filename=$(basename "$file")
            temp_output="${file}.temp.mp4"

            # 打印处理信息
            echo "处理文件：$filename"

            # 执行ffmpeg命令，输出到临时文件
            ffmpeg -i "$file" -vf "transpose=$transpose_value" -c:v libx265 -tag:v hvc1 -c:a aac -map_chapters -1 "$temp_output"
            
            if [ $? -eq 0 ]; then
                echo "转换成功"
                # 删除原文件
                rm "$file"
                # 将临时文件重命名为原文件名
                mv "$temp_output" "$file"
                echo "完成：$filename"
            else
                echo "转换失败：$filename"
                # 如果转换失败，清理临时文件
                rm -f "$temp_output"
            fi
        fi
    done
}

# 处理右侧和左侧的文件
process_files "$root/toRight" 1
process_files "$root/toLeft" 2
