# 镜像来拉取网站 
```
https://gitee.com/wanfeng789/docker-hub
```

下载镜像
```
docker pull redis
```

运行
```
docker run redis
```

制作镜像
```
docker build xxx
```

推送镜像
```
docker push xxx
```


# 安装Docker
关闭防火墙
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

yum源替换

```sh
#7.无法使用yum解决方案
#查看防火墙状态
systemctl status firewalld.service
#如果是运行状态，执行命令关闭防火墙并永久禁用
systemctl stop firewalld.service & systemctl disable firewalld.service
#还是无法正常使用yum源，大概率就是yum源的问题，更换阿里和清华的镜像
#打开CentOS-Base.repo
vim /etc/yum.repos.d/CentOS-Base.repo
#将CentOS-Base.repo里的内容全部替换为下面部分
```

# 不成功去查看
```
https://www.cnblogs.com/kohler21/p/18331060
```

## 启动docker
```
systemctl start docker
```

开机自动启动
```
systemctl enable docker
```

查看状态
```
sudo systemctl status docker
```

查看进程
```
ps -ef | grep docker
```

查看版本
```
docker version

docker -v  简洁版
```

## 设置镜像
```
tee /etc/docker/daemon.json

{
  "registry-mirrors": ["https://iblvsy56.mirror.aliyuncs.com"]
}
EOF

systemctl daemon-reload  重新加载系统配置

systemctl restart docker  重启docker服务
```

# Docker镜像管理和定制





修改镜像名称
```
docker tag [要改镜像] [新名称]
docker tag docker.1ms.run/nginx:latest docker.1ms.run/nginx:test
```


通过id删除
```
docker rmi 11
```

删除仓库名为ubuntu的镜像
```
docker rim $(docker images -q ubuntu)
```

删除本地所有镜像
```
docker rim $(docker images -q)
```

导出生成tar包
```
docker save -o nginx.tar docker.1ms.run/nginx:latest
```

将tar文件包导入
```
nginx.tar
```

上传的镜像仓库
```
docker push docker.1ms.run/nginx:latest
```

创建容器网络
```
docker network create my-nework
```


## 构建私有库
使用官方Registry镜像创建私有仓库容器
```
docker run -d -p 5000:5000 --name registry registry:2
```

添加仓库地址
```
{
  "insecure-registries" : ["your-private-registry-ip:5000"]
}
```

重启docker服务
```
sudo systemctl restart docker
```

# Docker命令
删除旧版本docke
```
sudo yum remove docker \
>                   docker-client \
>                   docker-client-latest \
>                   docker-common \
>                   docker-latest \
>                   docker-latest-logrotate \
>                   docker-logrotate \
>                   docker-engine
```

配置docker下载源
```
sudo yum install -y yum-utils
```

安装Docker引擎
```
sudo yum install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

启动docker
```
sudo systemctl start docker
```

开机启动
```
ystemctl enable docker
```

查看docker运行状态
```
systemctl status docker
```

查看正在运行的应用
```
docker ps
```

配置加速
```
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
    "registry-mirrors": [
        "https://mirror.ccs.tencentyun.com",
        "https://docker.m.daocloud.io"
    ]
}
EOF
```

重启docker进程与docker
```
sudo systemctl daemon-reload
sudo systemctl restart docker
```

## 官网镜像源网站
```
http://hub.docker.com/
```

搜索镜像
```
docker search [镜像名]
```

下载
```
docker pull [镜像]
获取镜像
docker pull docker.1ms.run/library/mysql:5.7

docker pull docker.1ms.run/nginx:latest
```


查看镜像
```
docker images

-a 列出所有镜像
-q 显示镜像ID
```

删除镜像
```
docker rmi [选项] [镜像名称]

-f 强制删除
--no-prune 
```

运行镜像
```
docker run nginx
```

查看运行的docker
```
docker ps
-a :查看所有镜像(关闭与开启的)
```

停止
```
docker stop [镜像名称]
```

启动镜像
```
docker start [镜像ID]
```

重启docker
```
docker restart [镜像ID]
```

查看状态
```
docker stats [镜像ID]
```

查看日志
```
docker logs f4d4
```

## 启动容器
run命令
```
docker run -d --name myningx nging
-d 后台运行
--name 命名
```

端口映射
```
docker run -d --name mynginx -p 80:80 nginx
```

进入容器
```
docker exec -it mynginx /bin/bash
-it 交互模式(命令的发送)
```

进入nginx目录
```
cd /usr/share/nginx/html/
```

## 修改页面内容
```
echo "<h1>Hello,Docker.</h1>" > index.html 
```

确认修改内容
```
cat index.html 
```

退出
```
exit
```

## 保存镜像
提交
```
docker commit -m"update index.html" mynginx  mynginx:v1.0
-a 作者
-c 有哪些改变的页表
-m 此次提交的信息
-p 打包期间暂停运行
```
