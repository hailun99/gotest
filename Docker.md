# 安装Docker
查看版本
```
uname -u
```

关闭防火墙s
```
systemctl stop firewalld  
```

设置开机禁用防火墙
```
systemctl disable firewalld
```

查看防火墙状态
```
systemctl status firewalld
```

修改文件
```
setenforce 0

vi /etc/selinux/config

SELINUX=disabled
```

修改网卡配置信息
```
vi /etc/sysconfig/network-scripts/ficfg-eno16777736

TYPE=Ethernet
BOOTPROTO=static
IPADDR=192.168.66.137
NETMASK=255.255.255.0
GATEWAY=192.168.66.1
DNS1=114.114.114.114
ONBOOT=yes
```

重启网卡
```
systemctl restart network
```

测试链接
```
ping -c 4 www.baidu.com
```

# 不成功去查看
```
https://www.cnblogs.com/kohler21/p/18331060
```
