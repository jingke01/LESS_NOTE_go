```shell
#!/usr/bin/env bash
mcd(){
  mkdir "$1"
  cd "$1"
}
```

''是强引用原样输出$1 

""是弱引用会解析将$1转换为第一个参数

    $0 - 脚本名
    $1 到 $9 - 脚本的参数。 $1 是第一个参数，依此类推。
    $@ - 所有参数
    $# - 参数个数
    $? - 前一个命令的返回值
    $$ - 当前脚本的进程识别码
    !! - 完整的上一条命令，包括参数。常见应用：当你因为权限不足执行命令失败时，可以使用 sudo !! 再尝试一次。
    $_ - 上一条命令的最后一个参数。如果你正在使用的是交互式 shell，你可以通过按下 Esc 之后键入 . 来获取这个值。

```shell
false || echo "Oops, fail"
# Oops, fail

true || echo "Will not be printed"
#

true && echo "Things went well"
# Things went well

false && echo "Will not be printed"
#

false ; echo "This will always run"
# This will always run
```
|| 和 $$ 或和与

; 在命令后的同一行接另一语句

$()把结果变成字符串

<()把结果写入一个虚拟文件地址

```shell
#!/usr/bin/env bash
echo "Starting program at $(date)"
echo "Running program $0 with $# arguments with pid $$"
for file in "$@";do
  grep foobar "$file" > /dev/null 2> /dev/null
  #if not find , grep exit at statue 1
  #stdin>null stdout>null, we don't care these
  if [ $? -ne 0 ]; then
      echo "File $file does not have any foobar ,adding one"
      echo "# foobar" >> "$file"
  fi
done
```
通配符 ? 匹配一个 * 匹配多个

{} 有一系列公共子串可以用{}把不同的地方括起来 系统会自动展开

```shell
convert image.{png,jpg}
# 会展开为
convert image.png image.jpg

cp /path/to/project/{foo,bar,baz}.sh /newpath
# 会展开为
cp /path/to/project/foo.sh /path/to/project/bar.sh /path/to/project/baz.sh /newpath

# 也可以结合通配使用
mv *{.py,.sh} folder
# 会移动所有 *.py 和 *.sh 文件

mkdir foo bar

# 下面命令会创建 foo/a, foo/b, ... foo/h, bar/a, bar/b, ... bar/h 这些文件
touch {foo,bar}/{a..h}
touch foo/x bar/y
# 比较文件夹 foo 和 bar 中包含文件的不同
diff <(ls foo) <(ls bar)
# 输出
# < x
# ---
# > y
```

下载tldr
```shell
# 下载二进制文件 (适用于大多数 64 位 Linux)
wget https://github.com/dbrgn/tealdeer/releases/latest/download/tealdeer-linux-x86_64-musl -O /usr/local/bin/tldr

# 赋予执行权限
chmod +x /usr/local/bin/tldr

# 刷新 bash 缓存
hash -r

# 更新缓存（tealdeer 的更新命令是 --update）
tldr --update

# 测试
tldr tar
```
