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

## 文件处理命令
```
cd   // change directory 切换目录
pwd  // print working directory 显示当前所有目录
touch // 创建文件
ctrl+c // 终止命令
```

### mkdri [参数] 目录
```
mkdir test // 创建目录
mkdri -p test/subdir2 // 创建子目录
```

### rmdir  文件要为空
```
rmdir tset // 删除目录
rmdir -p test/subdir2 // 强制删除目录及目录下所有文件
rm -rf test // 强制删除目录及目录下所有文件
```

### touch
```
touch test.txt // 创建空文件
```

### cp
```
cp /root/test.txt /home/bbb // 复制文件到别的目录
\cp -r /home/bbb /opt/ // 强制覆盖文件
-a // 复制
-f // 先删除在复制
-i // 文件已存在，会提示
-R // 递归复制
-u // 有差异时才会被复制
```

### rm [参数] 文件名或目录名
```
rm -rf /home/bbb // 删除目录及目录下所有文件
-f // 删除文件或目录，不提提示用户
-R //递归删除目录，即字目录
```

### mv  移动文件、更名
```
mv cat.txt pig.txt // 重命名
mv pig.txt /root/ // 移动文件
mv pig.txt /root/cow.txt // 移动并重命名文件
mv bbb/ /home/ // 移动目录
```

### cat [参数] 文件名  // 显示文件内容
```
cat -n /etc/profile // 显示文件内容，并显示行号
cat -n /etc/profile | more // 显示文件内容，并显示行号，并分页显示
-b // 对输出内容中的非空行标注行号
-n // 对输出内容中的所有行标注行号
cat file1 file2 > file3 // 如果file3文件按存在，覆盖file3文件原有内容
cat file1 file2 >> file3 // 如果file3文件按存在，将1、2文件内容附加到file3
```

### more [文件名]   分页显示文件内容
```
(空格)或f  // 显示下一页
(Enter)    // 显示下一行
q或Q       // 退出
= // 输出当前的行号
```

### less
```
less [文件名] // 分页显示文件内容
/字串 // 查找字串n:向下查找; N:向上查找
```

### echo 
```
echo "hello,world" // 输出hello,world
echo $PATH // 输出PATH环境变量
echo $HOSTNAEM // 输出环境变量
echo $HOSTNAME // 输出环境变量
```

### head
```
head -5 [文件名] // 查看文件的前几行
     -5 // 显示文件的前5行
     -f // 动态显示文件内容
```

### tail
```
tail 5 [文件名] // 查看文件的后几行
     -f // 将恐动态显示文件内容
```

### > 和 >>
```
> // 写入/覆盖文件内容
>> // 追加内容到文件
echo "hello,world" > [文件名] // 追加内容到文件
ls -l /home > /home/info.txt // 将ls -l /home的输出写入到/home/info.txt
cal >> /home/mycal // 将cal的输出追加到/home/mycal
```

## ln 软链接文件
```
ln -s /root/  /home/myroot // 创建软链接
ls [文件名] 硬链接 不能跨文件系统 拷贝同步更新
ls -i i节点通过数字找文件,在任何数据里都有数字
```

## history
```
history  // 查看历史命令
history 10 // 查看近10条历史命令
!371 // 执行第371条命令
```

# 查找类
## find
```
find /home -name info.txt // 查找/home目录下所有文件名包含info.txt的文件
find /opt -user mycal // 查找mycal用户所有文件
find / -size  +200M // 查找大于200M的文件
ls -ln // 显示文件权限
```

## locate
```
sudo yum install mlocate // 安装
updatedb // 更新数据库
locate hello.txt // 查找hello.txt
```

## which
```
which ls // 查找ls命令在什么位置
```

## grep 过滤查找
```
cat /home/hello.txt | grep -n "yes" // 查找hello.txt文件中包含yes的行
grep -n "yes" /home/hello.txtn
-i // 忽略大小写
grep -v "^#" /etc/inittab | more // 查看inittab文件中没有#的行
grep shhd /var/log/messages // 查看sshd服务信息
```

# 压缩解压
## gzip
```
gzip /home/hello.txt // 压缩hello.txt
```

## zip/unzip
```
zip -r myhome.zip /home // 压缩myhome.zip
```

## tar
```
tar -zcvf pc.tar.gz /home/cat.txt /home/pig.txt // 压缩pc.tar.gz
tar -zcvf myhome.tar.gz /home/ // 压缩myhome.tar.gz
tar -zxvf pc.tar.gz // 解压pc.tar.gz
tar -zxvf /home/myhome.tar.gz -C /opt/tmp2 // 解压myhome.tar.gz
-c // 产生.tar压缩文件
-v // 显示压缩过程
-f // 指定压缩文件名
-z // 打包同时压缩
-x // 解压
```

# vim/vi命令
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
h // 左移动一个字符
j // 下移动一行
k // 上移动一行
l // 右移动一个字符
$ // 移动行尾
0 // 移动行首
H // 移至屏幕上端
M // 移至屏幕中央
L // 移至屏幕下端
```

```
: set number // 设置行号
:set nonu // 取消行号
gg // 到第一行
G // 到最后一行
nG // 到第n行
:n // 到第n行
```

## 删除命令
```
x // 删除光标所在处字符
nx // 删除光标所在处后n个字符
dd // 删除光标所在行, ndd删除n行
dG // 删除光标所在行到末尾的内容
D // 删除从光标所在处到行尾
:n2,n2d // 删除n1行到到n2行的指定范围
```

## 复制和剪切命令
```
yy 、Y // 复制当前行
nyy 、ny // 复制当前以下n行
dd // 剪切当前行
ndd // 剪切当前行一下n行
p 、P // 粘贴在当前光标所在行下或行上
```

## 替换和取消
```
r // 取消光标所在处字符
R // 从光标所在处开始替换字符,按Esc结束
u // 取消上一步操作
```

## 搜索和替换命令
```
输入20 在输入shift+g // 
/link  // 向前搜索指定字符串stp
      // 搜索时忽略大小写 :set ic
      // 关闭忽略大小写 :set noic
n // 搜索指定字符串的下一个出现位置
:%s/stp/new/g // 全文替换指定字符串 g 强制替换 
:n1,n2s/stp/new/g \c  // 在一定范围内替换指定字符串  c 确认是否替换
```

## 保存提出
```
:wq  = ZZ
:wq!  // 强制提出(条件:管理员或者文件所有者)
:q! // 不保存退出
:w // 保存
:w /root/... // 保存到某个位置
```

## 关机&重启
```
shutdown -h now // 立刻进行关机
shutdown -r 1 // "hello 1分钟后关机了"
shutdown -r now // 立刻进行重启
halt now // 关机
reboot now // 立刻进行重启
sync // 把内存的数据同步到磁盘
```

## 
```
:r /etc/...  // 导入某个文件
:! // 在Vi中执行命令
:r !date // 把当前时间导入系统
```
## 定义快捷键   在~/.vimrc目录下编辑
Map 快捷键 触发命令
制作快捷键 ctrl+c+p  p可换作如何字母
```
:map ctrl+P I#<ESC> // 自作一个快捷键p,在第一行加入#号并保留编辑模式
:map ctrl+B 0x // 删除行首第一个字符
:unmap ctrl+c+P // 取消快捷键P
4,8s/^/#/g // 4到8行全部注释
4,8s/^#//g // 行首的#号去掉
:1,5s/^/\/\//g // 行首用//进行注释,使用转义符

:ab qin qincy@qq.com // 替换,打出qin回车自动生成qincy@qq.com
:nuab qincy@qq.com // 取消对其定义
```

# 引导流程
## 固件设置
### 查看软件时间
```
date // 查看当前时间
date +%Y // 查看当前年份
date +%m // 查看当前月份
date +%d // 查看当前日期
date +%Y-%m-%d%H:%M:%S // 查看当前年月日时分秒
date -s "2020-01-01 00:00:00" // 设置当前时间
```

修改软件时间
```
date 021617022025.30
```

### cal
```
cal // 显示当前月份的日历
cal 2020 // 显示2020年的日历
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
/etc/cr.d/init.d/sshd // 回车参看其使用方法
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

## GRUB(软连接文件)
查看存放目录
```
ls /boot/grub
```

```
default 3 // 设置启动时间为s3秒

timeout 3 // 设置缺省等待时间为3秒

splashimage // 定义GURB界面图片

hiddenmenu // 设置隐藏菜单

title // 定义菜单项名称

root // 设置GRUB的根设备即内核所在的分区

kernel // 定义内核所在位置

initrd // 命令加载镜像文件
```

### GRUB命令
功能键
```
e: // 编辑当前的启动菜单项
c: // 进入GRUB的命令行方式
b: // 启动当前的菜单项
d:// 删除当前行
ESC: // 返回GRUB启动菜单界面,取消对当前菜单项所做的任何修改
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

## 校验
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
### 安装软件包
```
yum install sudo
```

### 检测升级
```
yum check-update // 检测所有软件包
yum check-update sudo
```

### 升级
```
yum update sudo // 升级sudo
```

### 软件包查询
```
rpm -q sudo
yum list sudo

yum list | grep sudo
```

### 软件包信息
```
yum info sudo
```

### 卸载
```
yum remove sudo
```

### 帮助
```
yum -help
man yum
```

### 查询选项
```
-a //查询所有已安装的软件包
-f //查询文件所属软件包
-p //查询软件包
-i //查询软件包信息
-l //显示软件包中的文件列表
-d //显示被标注为文档的文件列表
-c //显示被标注为配置文件的文件列表
```

## APT包管理
```
apt-cache search //搜索软件包
apt-cache show //查看软件包信息
apt-get install (reinstall、-f) //安装软件包
apt-get remove (autoremove、-purge)//卸载软件包
apt-get update //更新软件包列表
apt-get upgrade //升级软件包
```

# Liunx用户管理

## 账户管理
```
logout // 退出当前用户
```

### 添加新的用户
```
useradd 选项 用户名
参数
-c // comment 指定一段注释性描述
-d // home 指定用户主目录，目录不存在,使用-m创建主目录
-g // group 指定用户所属组
-G // group 指定用户所属组的附加组
-s // shell 指定用户的登录shell
-u // uid 指定用户UID
```

### 删除用户
```
userdel wangwu // 删除用户,保留家目录
userdle -r 张三 // 删除用户,及主目录
find // 查找用户文件手动删除
```

## 清屏
```
clear // 清屏
```

## 添加用户
```
<!-- passwd wangwu // 添加用户 -->
passwd -d /home/test king // 创建用户的指定目录
```
useradd 设置选项 用户名 -D 查看缺省参数
```
u: // UID
g: //缺省组所属用户UID
G: //指定用户所属多个组
d: // 宿主目录
s: // 命令解释器
c: // 描述信息
e: 指定用户失效时间
suermod -G sys 用户名 // 把用户添加到组中
gpasswd -a 用户名 组名 // 给用户组添加成员
```

### 显示目录所在位置
```
pwd
```

### 退出
```
logout // 退出当前用户
exit // 退出当前用户
```

### 查看当前用户/登录用户
```
who ma i
```

## 配置文件

### 用户信息文件/etc/passwd文件格式
```
root:x:0:0:root:/root:/bin/bash
用户名:密码位:UID:GID:描述信息：宿主目录:shell

用户名 // 用户登录系统时使用的用户名
密码 // 密码位
UID // 用户标识号
GID // 缺省组标识号
注释性描述 //例如存放用户全名等信息
宿住目录 // 拥护登录系统后的缺省目录
命令解释器 //用户使用的Shell,默认为bash
```

### 生成一个加密密码
```
echon "12345" | mad5sum
```
### etc/shadow文件格式
```
root:$6$GtDEJN14P6he8VqH$SL/uHvSTEfMGwvqE6QUZwIuUpA3KBBWGONejeif9sbqehy39DMu9QbM/eeQnMNMTgZe80YguHO49pmt6D2rFL/::0:99999:7:::
用户名:加密密码:最后一次修改密码时间:最小时间间隔:最大时间间隔:警告时间:失效时间:标志

/etc/group // 用户组文件
/etc/gshadow // 用户密码文件
/etc/login.defs  /etc/default/useradd // 用户配置文件:js文件
/etc/skel // 新用户信息文件
/etc/motd  /etc/issue // 登录信息 
```

## etc/group文件格式
```
vi etc/group // 查找gruop文件
```

## steUID的定义
将touch命令授权setUID权限
```
chmod g+s // 添加setUID权限
chmo g-s // 删除setUID权限
chnod u+S // 添加setGID权限
chnod u-S // 删除setGID权限
chmod o+t // 添加setTTY权限
chmod o-t // 删除setTTY权限
```

### 查找SetUId程序
```
find / -perm -4000 -o -perm -2000
```

### 组的组成部分
```
sys::3:root,bin,adm
组名:组密码位:GID:组成员
```

## 用户组
```
groupadd wudang // 添加用户组
```

### 查看用户组id
```
id wzj // 查看用户组id
```

### 删除组
```
groupdel wudang // 删除组
```

### 增加用户时直接加上组
```
useradd -g wudang zwj // 添加用户到wudang组
```

### 修改用户组
```
usermod -g mojiao zwj // 修改用户组
```

## 用户组管理命令
gpasswd 设置组密码及管理组内成员
```
-a // 添加用户到用户组
-d // 从用户组中删除用户
-A // 设置用户组管理员
-r // 删除用户组组密码
-R //禁止用户切换为改组
gpasswd 组名 // 给组设置密码
groups // 查看用户属于那些用户组
nwegep // 切换用户组
grpck // 修改文件所属组
cigr //编辑/etc/group文件
groupdel // 删除用户组
groupmod // 修改用户信息
pwck //检测/etc/paswwd文件
vipw // 编辑/etc/passwd文件
id // 查看用户id和组信息
finger 用户 // 查看用户详细信息
su // 切换用户(su - 环境变量切换)
passwd -S // 查看用户密码状态
who/w // 查看当前登录用户信息
```


## 运行级别
```
0 // 关机
1 // 单用户模式
2 // 多用户状态没有网络模式
3 // 多用户状态有网络模式
4 // 系统未使用保留用户
5 // xWindo图形多用户模式
6 // 重启
```

### 切换级别
```
init 3 // 切换到运行级别3
```

### 查看当前运行级别
```
systemctl get-default // 查看当前运行级别
```

# 找回root密码步骤
```
init=/bin/sh
快捷键 Ctrl+x 回车
mount -o remount,rw / 
passwd
输入密码 root
touch /.autorelabel
exec /sbin/init
```

# 帮助指令
```
man [命令或配置文件] // 查看帮助
help // 查看帮助
```

### chage
```
-l // 查看用户密码设置
-m // 修改密码最小天数
-M // 修改密码最大天数
-d // 密码最后修改的日期
-I // 密码失效后,锁定账号的天数
-E // 设置密码的过期时间,0代表密码立刻过期,-1代表密码永不过期
-W // 密码过期前多少天发送邮件警告
```

## 批量添加用户
```
newusers // 批量添加用户
pwunconv // 批量删除用户
chpasswd // 导入密码文件
pwconv // 将密码写入shadow文件
```


# 获取命令帮助
```
help 内置命令 // 使用help命令查看指定的Shell内置命令的使用方法
命令名 --help // 使用命令名后加--help查看命令的用法摘要和参数列表
whatis 命令名 // 使用whatis命令获取指定的简要功能描述
man 命令名 // 使用man命令查看指定的手册 (按Q键退出)
info/pinfo 命令名 // 使用info或pinfo命令查看指定的GNU项目文档
man -k <关键字>/apropos <关键字> // 列出所有与 <关键字>匹配的手册页
```

# 所属组
## 文件 / 目录 所有组
### 查看文件所属组
```
ls -ahl l //查看文件所属组
```

### 修改文件所属组
```
chown tom apple.txt // 修改文件所属组
```

### 组的创建
```
groupadd monster // 创建组
useradd -g monster fox // 添加用户到组
id fox // 查看用户所属组
```

### chgrp
```
chgrp fruit orang.txt // 修改文件所属组
```

## 其他组
### 改变用户所在组
```
usermod -g wudang zwj // 改变所有组
usermod -d 目录 用户名 // 改变用户登录地初始目录
```

# 权限管理命令
```
-rwxrwxrwx  1 root root  718 12月 11 20:43 compose.yaml // 文件权限
d // 目录文件
- // 普通文件
l // 链接文件
c // 字符设备文件,鼠标键盘
b // 块设备文件,硬盘等
第一组rwx: 文件拥有者的权限是读写和执行
第二组rwx: 所有组的用户拥有读写和执行权限
第三组rwx: 其他用户拥有读写和执行权限
1 // 文件:硬链接数或 目录:子目录数
root // 用户
root // 组
718 // 文件大小
12月 11 20:43 // 文件修改时间
compose.yaml // 文件名
```

## 文件类型
```
r-read读、w-write写、x-execute执行
rwx       r-x        r-x 
 7         5          5
r = 4     w = 2      x = 1
所有者u   所有组g     其他人o
user      group      oyhers
onwer  
```

### 权限
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
u:所有者 g:所有组 o:其他用户 a:所有用户
chmod u-x,g+w apple.txt // 减去权限/添加权限
chmod u=rwx,g=rx,o=rx apple.txt // 修改权限
chmod a+r apple.txt // 添加权限
chmod 641(rw- -wx --x) // 修改权限
chmod 751 /home/info.txt  // 修改权限
```

## 修改文件所有者
```
chown tom /home/info.txt // 修改文件所有者
chown -R tom /home/test //目录下的所有文件和目录的所有者都修改成tom
```

## 修改文件/目录所在组
```
chgrp shaoli /home/info.txt // 修改文件/目录所在组
chgrp -R shaoli /home/test // 修改目录下所有文件和目录的组
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

# crond 任务调度
```
crontab -e // 编辑crontab定时任务
crontab -l // 查看crontab任务
crontab -r // 删除当前用户所有的crontab任务
service crond start // 启动crond服务
*/1 * * * * ls -l /etc/ > /tmp/test.txt
0 0 1,15 * * // 每月1号和15号执行 
35 8 * * 1 // 周一的8:35执行
第一个* // 表示分钟 0-59
第二个* // 表示小时 0-23
第三个* // 表示日期 1-31
第四个* // 表示月份 1-12
第五个* // 表示星期 0-7(0和7都表示星期天)
```

## 编写crontab定时任务
```
vim ym.txt
data >> /tmp/ym.txt
cal >> /tmp/ym.txt 
chmod u+x my.sh // 添加权限
./my.sh // 执行
crontab -e // 编辑crontab定时任务
1 * * * * /home/ym.sh // 1分钟执行
0 2 * * * mysqldump -u root -p123456 testdb > /tmp/ym.sql // 备份数据库
```

# at 定时任务
```
ps -ef | grep atd // 查看atd服务
-m // 邮件通知
-I // atq别名
-d // atrm别名
-v // 显示任务将被执行时间
-c // 打印任务的内容到标准输出
-V // 显示版本信息
-q <列队> // 列出队列
-f <文件> // 从文件中读取任务
-t <时间参数> // 指定执行时间
Ctrl+d // 按两下退出
at 5pm + 2 days   /bin/ls /home // 两天后下午五点执行/bin/ls /home
atq // 查看系统中没有执行的任务
at 5pm tomorrow // 明天下午五点执行
at now + 2 minutes // 2分钟后执行
at 24.03.2025 18:00 // 2025年3月24日18:00执行
atrm 4 // 删除第四个任务
```

# 磁盘管理
```
lablk // 查看磁盘信息
lablk -f // 显示磁盘分区信息
```

虚拟机添加硬盘
```
fdisk /dev/sdb // 创建硬盘
m  // 显示命令列表
p // 显示磁盘分区信息
n // 新建分区
d // 删除分区
w // 写入并退出
```
## 挂载
```
mkfs -t ext4 /dev/sdb1 // 格式化
mkdir newdisk // 创建目录
mount /dev/sdb1 /newdisk/ // 挂载
umount /dev/sdb1 // 卸载挂载
vim /etc/fstab // 配置挂载文件
/dev/sdb1   /newdisk   xfs  defaults   0 0 // 自动挂载
mount -a // 即可生效挂载的文件
```

## 查看磁盘使用情况
```
du -h // 显示磁盘使用情况
-s // 显示当前目录下文件大小
-h // 带计量单位
-a // 显示所有文件
--max-depth=1 // 只显示当前目录下文件大小
-c // 显示总大小
du -h --max-depth=1 /opt // 显示opt目录下文件大小
```

## 食用命令 
```
ls -l /opt | grep "^-" | wc -l // 显示opt目录下文件数量
ls -lR /opt | grep "^b" | wc -l // 显示opt目录下块设备数量
ls -l /opt | grep "^d" | wc -l // 显示opt目录下目录数量
tree /opt // 显示目录树结构
```

# 网络配置
```
yum install net-tools // 安装ifconfig
ifconfig // 查看网络信息
ipconfig // Windows查看网络信息
```

## 指定设置ip
```
vim /etc/sysconfig/network-scripts/ifcfg-ens33 // 配置网络环境文件
BOOTPROTO="static" // 静态ip
ONBOOT="yes" // 启动
IPADDR=192.168.1.100 // ip
GATEWAY=192.168.1.1 // 网关
DNS1 = 114.114.114.114 // 域名解析器
reboot // 重启
service network restart // 重启网络服务
```

# 主机和hosts映射
```
vim /etc/hostname // 修改主机名
```

## Windows
```
C:\Windows\System32\drivers\etc\hosts // 修改hosts文件
192.168.1.100 hspedu100 // 映射主机名
ping hspedu100 // ping测试
```

## Linux
```
vim /etc/hosts // 修改hosts文件
192.168.1.100 hspedu100 // 映射主机名
ping hspedu100 // ping测试
```


# shell
字符
```
* //匹配任意字串
? //匹配任意字符
/ //代表根目录或作为路径间隔符使用
\ //转义字符
\<Enter> //续行符
$ //变换值位置
| //管道
& //后台执行
```