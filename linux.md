# 错误提示
找不到执行文件
```
Command not found
```

目录敲错了
```
No such file or directory
```
# 常用命令
关闭防火墙
```
systemctl stop firewalld.service
```

查看防火墙状态
```
systemctl status firewalld.service
```

开启防火墙
```
systemctl start firewalld.service
```

查看是否安装NTP
```
rpm -qa | grep ntp
```

安装NTP
```
yum install ntp ntpdate -y
```

生成密钥
```
ssh-keygen -t rsa -C "xxx@xx.com"
```

查看网卡
```
ifconfg
```

虚拟网卡
```
loopback
```

# 命令远程管理工具
```
Putty
SecureCRT
```

# 退出系统命令
```
exit或logout
```

# Linux文件处理命令: 
ls 选项[-ald] [文件或目录]
```
-a 显示所有文件,包括隐藏文件
-A 显示所有文件,包括隐藏文件
-c 按文件修改的时间进行排序
-l 详细信息显示
-d 查看目录属性
```


## 文件类型
```
d 目录directory
- 二进制文件
l 软链接文件link
r-read读、w-write写、x-execute执行
rwx       r-x        r-x      7   5   5
r = 4     w = 2      x = 1
所有者u   所有组g     其他人o
user      group      oyhers
onwer  
```


## 软硬链接文件
```
ln -s [文件名] 软连接
ls [文件名] 硬链接 不能跨文件系统 拷贝同步更新
ls -i i节点通过数字找文件,在任何数据里都有数字
```


## 文件处理命令
```
cd   change directory 切换目录
pwd  print working directory 显示当前所有目录
touch 创建文件
ctrl+c 终止命令
```

cat [参数] 文件名  显示文件内容
```
-b 对输出内容中的非空行标注行号
-n 对输出内容中的所有行标注行号
cat file1 file2 > file3 如果file3文件按存在，覆盖file3文件原有内容
cat file1 file2 >> file3 如果file3文件按存在，将1、2文件内容附加到file3
```

more/less [文件名]   分页显示文件内容
```
    (空格)或f  显示下一页
    (Enter)    显示下一行
    q或Q       退出
```

head
```
head -5 [文件名] 查看文件的前几行
     -5 显示文件的前5行
     -f 动态显示文件内容
```

tail
```
tail 5 [文件名] 查看文件的后几行
     -f 动态显示文件内容
```

mkdri [参数] 目录
```
mkdir test 创建目录
mkdri -p test/subdir2 创建子目录
```

rmdir  文件要为空
```
rmdir tset 删除目录
rmdir -p test/subdir2
```

cp
```
-a 复制
-f 先删除在复制
-i 文件已存在，会提示
-R 递归复制
-u 有差异时才会被复制
```

mv [源文件或目录] [目的目录] 移动文件、更名
```
mv mpt /usr/test 移动文件
mv /sur/tset/mpt /tt 移动，修改文件名字
```

rm [参数] 文件名或目录名
```
-i 删除文件或目录，不提示用户
-f 删除文件或目录，提示用户
-R 递归删除目录，即字目录
```

# 权限管理命令
```
文件
r -cat、more、head、tail
w -echo、vi
x -命令、脚本
目录
r -ls
w -touch、mkdir、rm
x-cd
```

## 改变权限
```
chmod u + x [文件名] 添加权限
chmod o - r [文件名] 减去权限
chmod g = rx [文件名] 修改权限
chmod 641(rw- -wx --x) 修改权限
```

## 添加用户
```
useradd 用户名
passwd 用户名
```


## 切换用户名
```
su - 用户名
```

## 改变所有者
```
chown 用户
```

## 改变所有组
```
chgrp 组名(adm)(系统里自带的)
```

## 查看文件权限
```
umask 掩码值(777-掩码值)
umask -S
```

## 文件搜索命令
显示系统命令所在目录
```
which ls 别名信息
whereis ls 帮助信息，帮助文件所存放的信息
```

查找文件目录
```
find [搜索路径] [搜寻关键字]
```

# vim命令
## 插入命令
```
a 在光标后附加文本
A 在本行行末附加文本
i 在光标前附加文本
I 在本行开始插入文本
o 在光标下插入文本
O 在光标上插入文本
```

## 定位命令
方向键
```
h 左移动一个字符
j 下移动一行
k 上移动一行
l 右移动一个字符
$ 移动行尾
0 移动行首
H 移至屏幕上端
M 移至屏幕中央
L 移至屏幕下端
```

```
: set number 设置行号
:set nonu 取消行号
gg 到第一行
G 到最后一行
nG 到第n行
:n 到第n行
```

## 删除命令
```
x 删除光标所在处字符
nx 删除光标所在处后n个字符
dd 删除光标所在行, ndd删除n行
dG 删除光标所在行到末尾的内容
D 删除从光标所在处到行尾
:n2,n2d 删除n1行到到n2行的指定范围
```

## 复制和剪切命令
```
yy 、Y 复制当前行
nyy 、ny 复制当前以下n行
dd 剪切当前行
ndd 剪切当前行一下n行
p 、P 粘贴在当前光标所在行下或行上
```

## 替换和取消
```
r 取消光标所在处字符
R 从光标所在处开始替换字符,按Esc结束
u 取消上一步操作
```

## 搜索和替换命令
```
/link  向前搜索指定字符串stp
      搜索时忽略大小写 :set ic
      关闭忽略大小写 :set noic
n 搜索指定字符串的下一个出现位置
:%s/stp/new/g 全文替换指定字符串 g 强制替换 
:n1,n2s/stp/new/g  c 在一定范围内替换指定字符串  c 确认是否替换
```

## 保存提出
```
:wq  = ZZ
:wq!  强制提出(条件:管理员或者文件所有者)
:q! 不保存退出
:w 保存
:w /root/... 保存到某个位置
```


## 
```
:r /etc/... 导入某个文件
:! 在Vi中执行命令
:r !date 把当前时间导入系统
```
## 定义快捷键   在~/.vimrc目录下编辑
Map 快捷键 触发命令
制作快捷键 ctrl+c+p  p可换作如何字母
```
:map ctrl+P I#<ESC> 自作一个快捷键p,在第一行加入#号并保留编辑模式
:map ctrl+B 0x 删除行首第一个字符
:unmap ctrl+c+P 取消快捷键P
4,8s/^/#/g 4到8行全部注释
4,8s/^#//g 行首的#号去掉
:1,5s/^/\/\//g 行首用//进行注释,使用转义符

:ab qin qincy@qq.com 替换,打出qin回车自动生成qincy@qq.com
:nuab qincy@qq.com 取消对其定义
```

# 引导流程
## 固件设置
### 查看软件时间
```
date
```

查看使用方法
```
man date
```

修改软件时间
```
date 021617022025.30
```

### 查看硬件时间
```
hwclock
```

查看使用方法
```
hwclock --help
man hwclock
```

修改硬件时间
```
hwclock --set --date="02/16/2025 17:07:05"
```

## inittab
查看当前允许级别
```
runlevel
```

切换级别
```
init[012345Ss](0-6级别)
```

### 指定运行级别，可以多个
```
run-levels
```

### 指定状态
```
action
```

### 指定要运行的脚本/命令
```
process
```

修改脚本名称
```
mv S10network s10network
```

启动/结束脚本命令
```
/etc/cr.d/init.d/sshd 回车参看其使用方法
/etc/cr.d/init.d/sshd start(stop)
```


## 查看设置自启动命令
```
ln -s

man chkconfig

man ntsysv
```

查看内核启动信息
```
dmesg

dmesg | grep rtho
```

日志信息存放位置
```
ls /var/log
```

### initdefault
```
1 单用户模式
2 3 字符界面多用户模式
5 xWindo图形多用户模式
0 关机
6 重启
4 预留的
```

## grep
查看开头没有#的数据
```
grep -v "^#" /etc/inittab | more 
```

查看服务(日志)信息
```
grep shhd /var/log/messages

grep syslog /var/log/messages
```

## GRUB(软连接文件)
查看存放目录
```
ls /boot/grub
```

```
default 3 设置启动时间为s3秒

timeout 3 设置缺省等待时间为3秒

splashimage 定义GURB界面图片

hiddenmenu 设置隐藏菜单

title 定义菜单项名称

root 设置GRUB的根设备即内核所在的分区

kernel 定义内核所在位置

initrd 命令加载镜像文件
```

### GRUB命令
功能键
```
e: 编辑当前的启动菜单项
c: 进入GRUB的命令行方式
b: 启动当前的菜单项
d:删除当前行
ESC:返回GRUB启动菜单界面,取消对当前菜单项所做的任何修改
```
设置GRUB密码
```
grub-md5-crypt
```

# 软件包管理 
## RPM包管理
卸载软件包
```
rpm -e 软件包
rpm -e --nodeps 软件包 //强制卸载
```

安装软件包
```
rpm -ivh 软件包
rpm -ivh --excludedocs 软件包 //不安装软件包中的文档文件
rpm -vih --test 软件包 //测试安装
rpm -ivh --replacepkgs 软件包 //覆盖安装
rpm -ivh --replacefiles 软件包 //覆盖安装
rmp -Uvh 软件包 //升级软件包
```

查询
```
rpm -q 软件包 
rpm -qa grep  软件包  //查询所有的软件包
rpm -qa | grep vim //查询vim相关软件包
```

挂载光盘
```
mkdir /mnt/cdrom
mount /dev/cdrom /mnt/cdrom
```

## RPM查询
```
rpm -qf //查询文件隶属的软件包
rpm -qi //查看已安装软件包的信息
rpm -qip sudo //查看未安装软件包信息
rpm -ql //查看已安装软件包是干什么的
rpm -qlp sudo //查看未安装软件包是做什么的
rpm -qd sudo //查看软件包帮助文档
rmp -qc sudo/ /查看软件包设置文件
```

校验
rpm -V sudo
```
5 //文件的md5校验值
S //文件大小
L //链接文件
T //文件的创键时间
D //设备文件
U //文件的用户
G //文件的用户组
M //文件权限
```

## yum命令
安装软件包
```
yum install sudo
```

检测升级
```
yum check-update // 检测所有软件包
yum check-update sudo
```

升级
```
yum update sudo
```

软件包查询
```
rpm -q sudo
yum list sudo

yum list | grep sudo
```

软件包信息
```
yum info sudo
```

卸载
```
yum remove sudo
```

帮助
```
yum -help
man yum
```

查询选项
```
-a //查询所有已安装的软件包
-f //查询文件所属软件包
-p //查询软件包
-i //查询软件包信息
-l //显示软件包中的文件列表
-d //显示被标注为文档的文件列表
-c //显示被标注为配置文件的文件列表
```







