找不到执行文件
```
Command not found
```

目录敲错了
```
No such file or directory
```

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

## 文件处理命令
```
cd   change directory 切换目录
pwd  print working directory 显示当前所有目录
touch 创建文件
mkdir/[文件名] 创建目录
cp -R [源文件或目录] [目的目录]  复制文件或目录
ctrl+c 终止命令
mv [源文件或目录] [目的目录] 移动文件、更名
rm -r [文件或目录] 删除目录(文件或目录要为空)
cat [文件名] 显示文件内容
more [文件名]   分页显示文件内容
    (空格)或f  显示下一页
    (Enter)    显示下一行
    q或Q       退出
head -5 [文件名] 查看文件的前几行
     -5 显示文件的前5行
     -f 动态显示文件内容
tail 5 [文件名] 查看文件的后几行
     -f 动态显示文件内容
```

## 软硬链接文件
```
ln -s [文件名] 软连接
ls [文件名] 硬链接 不能跨文件系统 拷贝同步更新
ls -i i节点通过数字找文件,在任何数据里都有数字
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