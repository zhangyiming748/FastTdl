# FastTdl
批量下载电报媒体文件

# 支持的格式
```
channle = 频道id
fid = 文件id
tag = 归档的文件夹名
fname = 下载后重命名的正式文件名(不包含扩展名)
offset = 向下继续下载n个文件

https://t.me/${channel}/${fid}
https://t.me/${channel}/${fid}#${tag}
https://t.me/${channel}/${fid}#${tag}@${fname}
https://t.me/${channel}/${fid} ${offset}
https://t.me/${channel}/${fid}#${tag} ${offset}
```

**不支持的格式 `https://t.me/${channel}/${fid}#${tag}@${fname} ${offset}`**
