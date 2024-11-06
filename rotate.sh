#!/bin/bash


# 指定要查找的扩展名
extension="mp4"
folder="/mnt/c/Users/zen/Downloads/telegram/dance/toRight"
# 遍历当前目录下的所有文件
for file in "$folder"/*.$extension; do
    # 检查文件是否具有指定的扩展名
    if [[ $file == *.$extension ]]; then
        # 提取文件名和扩展名
        filename=$(basename "$file")
        basename_without_ext="${filename%.*}"
        ext="${filename##*.}"
        aftername=${basename_without_ext}_rotate.mp4
        # 打印分开的文件名和扩展名
        echo "文件名： $basename_without_ext"
        echo "扩展名： $ext"
        ffmpeg -i ${file} -vf transpose=1 -c:v h264_nvenc -preset medium -c:a copy -map_chapters -1 ${folder}/${aftername}
        echo "旧文件名: ${filename} 新文件名: ${aftername}"
        echo "done"
    fi
done
