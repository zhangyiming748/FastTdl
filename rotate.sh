#!/bin/bash


# 指定要查找的扩展名
extension="mp4"
root="/mnt/c/Users/zen/Downloads/media/维拉莎莎"

right=$root/toRight
# 遍历当前目录下的所有文件
for file in "$right"/*.$extension; do
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
		ffmpeg -i $file -vf transpose=1 -c:v h264_nvenc -preset medium -c:a copy -map_chapters -1 $right/$aftername
		if [ $? -eq 0 ]; then
			echo "旧文件名: $filename 新文件名: $aftername"
			echo "done"
			rm $file
		else
			echo "ffmpeg command failed for file: $file"
		fi
		echo "done"
	fi
done

left=$root/toLeft
# 遍历当前目录下的所有文件
for file in "$left"/*.$extension; do
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
		ffmpeg -i $file -vf transpose=2 -c:v h264_nvenc -preset medium -c:a copy -map_chapters -1 $left/$aftername
		if [ $? -eq 0 ]; then
			echo "旧文件名: $filename 新文件名: $aftername"
			echo "done"
			rm $file
		else
			echo "ffmpeg command failed for file: $file"
		fi
		echo "done"
	fi
done
