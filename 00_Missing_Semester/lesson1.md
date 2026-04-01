```shell
cat << 'EOF' > filename
这是内容 
可以将任意文本以原格式写入filename(filename是一个文件) 
直到下一个EOF要单独另起一行
EOF
```
```shell
cat anything
```
<< >> < > |重定向和管道可以将文件串联起来
