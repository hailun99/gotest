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



