>我自己的用法如下：
主要搞清频道id和消息id就可以了
频道id TG里面设置--高级--实验功能--显示id。随后点击频道右上方三点，显示频道信息。复制频道id
消息id 例子如下
消息链接：https://t.me/c/1492447836/251011/269724（251011 是主题 ID）
https://t.me/soqkkqossmsn/48334 48334是消息id soqkkqossmsn是频道id
第一步
tdl chat export -c 频道id或者用户名 --reply 回复的消息id
例子：tdl.exe chat export -c 1924435530 --reply 48334
或者
tdl.exe chat export -c soqkkqossmsn --reply 48334
默认导出json文件到tdl程序安装目录下
经测试--topic和--reply 效果类似，作者可以详细说明下作用
第二步
tdl.exe dl -f C:\tdl\tdl-export.json -d C:\TG\ -t 64 -l 5 --continue
-f 是json文件所在目录
-d为下载目录
ok，坐等下载完毕

----

以`https://t.me/DWdmic/3402`评论区为例
```bash
tdl chat export -c DWdmic --proxy http://127.0.0.1:8889 --reply 3402 
# 自动保存到用户家目录 不能也不需要更改
```
```bash
tdl download --proxy http://127.0.0.1:8889 --threads 8  --file tdl-export.json
# 自动在用户家目录寻找
```