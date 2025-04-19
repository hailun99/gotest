1. 确认 Nginx 是否已安装
执行以下命令检查是否已安装 Nginx：

bash
复制
rpm -qa | grep nginx
若有输出（如 nginx-1.20.1-1.el7.x86_64）：已安装，但服务未注册（跳到步骤3）。

若无输出：未安装，继续下一步。

2. 重新安装 Nginx
2.1 确保仓库配置正确
检查 /etc/yum.repos.d/nginx.repo 内容是否为官方仓库配置：

bash
复制
cat /etc/yum.repos.d/nginx.repo
标准配置（适用于 CentOS/RHEL 7/8）：

ini
复制
[nginx-stable]
name=nginx stable repo
baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
gpgcheck=1
enabled=1
gpgkey=https://nginx.org/keys/nginx_signing.key
2.2 清理缓存并安装
bash
复制
yum clean all
yum makecache
yum install nginx -y
3. 检查 systemd 服务文件
Nginx 安装后应自动生成服务文件 /usr/lib/systemd/system/nginx.service。若该文件丢失，手动创建：

bash
复制
vi /usr/lib/systemd/system/nginx.service
粘贴以下内容：

ini
复制
[Unit]
Description=The nginx HTTP and reverse proxy server
After=network.target remote-fs.target nss-lookup.target

[Service]
Type=forking
PIDFile=/run/nginx.pid
ExecStartPre=/usr/sbin/nginx -t
ExecStart=/usr/sbin/nginx
ExecReload=/usr/sbin/nginx -s reload
ExecStop=/bin/kill -s QUIT $MAINPID
PrivateTmp=true

[Install]
WantedBy=multi-user.target
保存后执行：

bash
复制
systemctl daemon-reload
4. 启动 Nginx 并验证
bash
复制
systemctl start nginx
systemctl status nginx
成功状态：显示 active (running)。

失败状态：检查日志定位问题：

bash
复制
journalctl -u nginx -xe
5. 常见问题解决
问题1：端口冲突
若 80 端口被占用（如 Apache），修改 Nginx 监听端口：

bash
复制
vi /etc/nginx/conf.d/default.conf
将 listen 80; 改为 listen 8080;，然后重启：

bash
复制
systemctl restart nginx
问题2：SELinux 或防火墙拦截
开放防火墙端口：

bash
复制
firewall-cmd --add-port=80/tcp --permanent
firewall-cmd --reload
临时禁用 SELinux（调试用）：

bash
复制
setenforce 0
6. 强制重新安装（终极方案）
若仍失败，彻底卸载后重装：

bash
复制
yum remove nginx -y
rm -rf /etc/nginx /var/log/nginx
yum install nginx -y
systemctl start nginx
操作总结
bash
复制
# 1. 确认安装状态
rpm -qa | grep nginx

# 2. 清理缓存并重装
yum clean all && yum install nginx -y

# 3. 检查服务文件
ls /usr/lib/systemd/system/nginx.service

# 4. 手动创建服务文件（若缺失）
vi /usr/lib/systemd/system/nginx.service
systemctl daemon-reload

# 5. 启动服务
systemctl start nginx



# 安装Nginx
```
yum install -y yum-utils // 安装yum-utils
```
## 手动添加,nginx的yum仓库
```
yum install vim -y // 安装vim
创建文件
vim /etc/yum.repos.d/nginx.repo

[nginx-stable]
name=nginx stable repo
baseurl=http://nginx.org/packages/centos/$releasever/$basearch/
gpgcheck=1
enabled=1
gpgkey=https://nginx.org/keys/nginx_signing.key
module_hotfixes=true

[nginx-mainline]
name=nginx mainline repo
baseurl=http://nginx.org/packages/mainline/centos/$releasever/$basearch/
gpgcheck=1
enabled=0
gpgkey=https://nginx.org/keys/nginx_signing.key
module_hotfixes=true
```
## 启动
```
yum install -y nginx // 安装nginx
systemctl start nginx // 启动nginx
systemctl stop nginx // 停止nginx
systemctl status nginx // 查看nginx状态
systemctl restart nginx // 重启nginx
systemctl enable nginx // 设置nginx服务自启动
systemctl disable nginx // 设置nginx服务不自启动
```


# 下载redis
```
tar -zxvf redis-6.2.8.tar.gz 

Redis是基于c语言编写的需要安装依赖，需要安装gcc：
yum install gcc-c++
或
yum -y install gcc 

查看gcc版本：
gcc -v

make // 编译
make install // 安装

vim /etc/profile // 添加以下内容

source /etc/profile // 使配置生效

使用以下命令安装 Redis：
sudo yum install epel-release
sudo yum install redis


创建服务文件
以 root 权限创建 /etc/systemd/system/redis.service 文件，并添加以下内容：
[Unit]
Description=Redis In-Memory Data Store
After=network.target

[Service]
User=redis
Group=redis
ExecStart=/usr/local/bin/redis-server /etc/redis.conf
ExecStop=/usr/local/bin/redis-cli shutdown
Restart=always

[Install]
WantedBy=multi-user.target


重新加载 systemd 管理器配置
sudo systemctl daemon-reload

启动 Redis 服务
sudo systemctl start redis

设置开机自启
sudo systemctl enable redis

检查 Redis 安装路径
which redis-server

使用redis-cli连接测试
redis-cli -p 6379

在linux命令行关闭redis只需要
systemctl stop redis


性能测试
redis 性能测试的基本命令如下：
redis-benchmark [option] [option value]

序号	选项	  描述	                          默认值
1	    -h	    指定服务器主机名	           127.0.0.1
2	    -p	    指定服务器端口	                   6379
3	    -s	    指定服务器 socket	
4	    -c	    指定并发连接数	                    50
5	    -n	    指定请求数	                      10000
6	    -d	    以字节的形式指定 SET/GET 值的数据大小	  2
7	    -k	    1=keep alive 0=reconnect	            1
8	    -r	    SET/GET/INCR 使用随机 key, SADD 使用随机值	
9	    -P	    通过管道传输 请求	                      1
10	    -q	    强制退出 redis。仅显示 query/sec 值	
11	   –csv	    以 CSV 格式输出	
12	   *-l*（L 的小写字母）	生成循环，永久执行测试	
13	    -t	仅运行以逗号分隔的测试命令列表。	
14	    *-I*（i 的大写字母）	Idle 模式。仅打开 N 个 idle 连接并等待。	

测试100个并发链接，每个并发100000请求
redis-benchmark -h localhost -p 6379 -c 100 -n 100000
                        
原文链接：https://blog.csdn.net/qq_52227892/article/details/130649748
```



# 集群
## 设置node1
设置主机名
hostnamectl set-hostname node1

查看主机名
hostname  

如果没有vim，就安装
yum install vim 

修改node1的网卡配置
vim /etc/sysconfig/network-scripts/ifcfg-ens33

TYPE="Ethernet"
PROXY_METHOD="none"
BROWSER_ONLY="no"
BOOTPROTO="static"
DEFROUTE="yes"
IPV4_FAILURE_FATAL="no"
IPV6INIT="yes"
IPV6_AUTOCONF="yes"
IPV6_DEFROUTE="yes"
IPV6_FAILURE_FATAL="no"
IPV6_ADDR_GEN_MODE="stable-privacy"
NAME="ens33"
UUID="4e895ccd-c869-4e7a-87a6-9232880805f8"
DEVICE="ens33"
ONBOOT="yes"
IPADDR="192.168.66.181"
NETMASK="255.255.255.0"
GATEWAY="192.168.66.2"
DNS1="192.168.66.2"

保存强制退出
:wq!

重启网卡
systemctl restart network


## 设置node2
设置主机名
hostnamectl set-hostname node2

查看主机名
hostname


修改node2的网卡配置
vim /etc/sysconfig/network-scripts/ifcfg-ens33

TYPE="Ethernet"
PROXY_METHOD="none"
BROWSER_ONLY="no"
BOOTPROTO="static"
DEFROUTE="yes"
IPV4_FAILURE_FATAL="no"
IPV6INIT="yes"
IPV6_AUTOCONF="yes"
IPV6_DEFROUTE="yes"
IPV6_FAILURE_FATAL="no"
IPV6_ADDR_GEN_MODE="stable-privacy"
NAME="ens33"
UUID="4e895ccd-c869-5e7a-87a6-9232880805f9"
DEVICE="ens33"
ONBOOT="yes"
IPADDR="192.168.66.182"
NETMASK="255.255.255.0"
GATEWAY="192.168.66.2"
DNS1="192.168.66.2"

保存强制退出
:wq!

重启网卡
systemctl restart network

## 设置node3
设置主机名
hostnamectl set-hostname node3

查看主机名
hostname


修改node2的网卡配置
vim /etc/sysconfig/network-scripts/ifcfg-ens33

TYPE="Ethernet"
PROXY_METHOD="none"
BROWSER_ONLY="no"
BOOTPROTO="static"
DEFROUTE="yes"
IPV4_FAILURE_FATAL="no"
IPV6INIT="yes"
IPV6_AUTOCONF="yes"
IPV6_DEFROUTE="yes"
IPV6_FAILURE_FATAL="no"
IPV6_ADDR_GEN_MODE="stable-privacy"
NAME="ens33"
UUID="4e895ccd-c869-6e7a-87a6-9232880805f9"
DEVICE="ens33"
ONBOOT="yes"
IPADDR="192.168.66.183"
NETMASK="255.255.255.0"
GATEWAY="192.168.66.2"
DNS1="192.168.66.2"

保存强制退出
:wq!

重启网卡
systemctl restart network

## 集群配置
vim /etc/hosts 

添加映射
192.168.66.181 node1
192.168.66.182 node2
192.168.66.183 node3

# 配置SSH 免密登
在每台机器上执行:
ssh-keygen -t rsa -b 4096 // 生成密钥，一路回车即可

在每台机器上执行:完成免密登录
ssh-copy-id node1
ssh-copy-id node2
ssh-copy-id node3

连接:
ssh node1 

退出
exit


# 安装jdk
创建目录存放jdk
mkdir -p /export/server

解压jdk到指定目录:
tar -zxvf jdk-8u451-linux-x64.tar.gz -C /export/server/ 

创建软连接
ln -s /export/server/jdk1.8.0_451/ /export/server/jdk

配置jdk环境变量
vim /etc/profile
添加以下内容:
export JAVA_HOME=/export/server/jdk
export PATH=$PATH:$JAVA_HOME/bin

保存退出
wq

使配置生效
source /etc/profile 

查看是否生效
echo $JAVA_HOME
echo $PATH

删除系统自带的java
rm -f /usr/bin/java

替换自己安装的java
ln -s /export/server/jdk/bin/java /usr/bin/java

查看jkd版本
java -version

# 关闭防火墙
关闭每台机器的防火墙
systemctl stop firewalld // 关闭防火墙
systemctl disable firewalld // 设置禁止启动


修改每台机器的selinux配置文件
vim /etc/sysconfig/selinux 

# This file controls the state of SELinux on the system.
# SELINUX= can take one of these three values:
#     enforcing - SELinux security policy is enforced.
#     permissive - SELinux prints warnings instead of enforcing.
#     disabled - No SELinux policy is loaded.
SELINUX=disabled
# SELINUXTYPE= can take one of three values:
#     targeted - Targeted processes are protected,
#     minimum - Modification of targeted policy. Only selected processes are protected. 
#     mls - Multi Level Security protection.
SELINUXTYPE=targeted

关机:
init 0


# 下载Zookeeper
wget https://archive.apache.org/dist/zookeeper/zookeeper-3.5.9/apache-zookeeper-3.5.9-bin.tar.gz

解压Zookeepertar包
tar -zxvf apache-zookeeper-3.5.9-bin.tar.gz -C /export/server/

构建软连接
ln -s /export/server/apache-zookeeper-3.5.9-bin/ /export/server/zookeeper

切换路径，修改文件名称
cd zoookeeper/conf/
mv zoo_sample.cfg zoo.cfg

修改配置文件
vim /export/server/zookeeper/conf/zoo.cfg

tickTime=2000
initLimit=5
syncLimit=2
// zookeeper数据存储目录
dataDir=/export/server/zookeeper/data
clientPort=2181
server.1=node1:2888:3888
server.2=node2:2888:3888
server.3=node3:2888:3888


## 配置zookeeper
配置环境变量
vi /etc/profile

export ZOOKEEPER_HOME=/export/server/zookeeper
export PATH=$PATH:$ZOOKEEPER_HOME/bin

## 创建zookeeper数据目录
mkdir /export/server/zookeeper/data

创建文件，并填入1
vim /export/server/zookeeper/data/myid

把zookeeper复制到另两台机器上
scp -r /export/server/apache-zookeeper-3.5.9-bin node2:`pwd`/
scp -r /export/server/apache-zookeeper-3.5.9-bin node3:$PWD

刷新环境变量
source /etc/profile

构建软连接
ln -s /export/server/apache-zookeeper-3.5.9-bin/ /export/server/zookeeper

分别修改node2和node3的myid文件
node2的修改为2
vim /export/server/zookeeper/data/myid
node3修改we3
vim /export/server/zookeeper/data/myid

## 启动zookeeper
文件中启动
cd /export/server/zookeeper/bin
./zkServer.sh start 

绝对路径启动
/export/server/zookeeper/bin/zkServer.sh start 

操作控制台
./zkCli.sh

三台主机都输入
zkServer.sh start
都输入完成之后，在输入

zkServer.sh status
查看当前主机状态

zkServer.sh stop // 停止zookeeper


# 安装kafka
wget http://archive.apache.org/dist/kafka/2.4.1/kafka_2.12-2.4.1.tgz


解压
tar -zxvf kafka_2.12-2.4.1.tgz -C /export/server/

软连接
ln -s /export/server/kafka_2.12-2.4.1/ /export/server/kafka

## 编辑文件
vim /export/server/kafka/config/server.properties

指定broker的id
broker.id=1

指定kafka监听端口
/listeners
listeners=PLAINTEXT://node1:9092

指定kafka数据位置
/dirs
log.dirs=/export/server/kafka/data

指定zookeeper的三个节点
zookeeper.connect=node1:2181,node2:2181,node3:2181

### 编辑noed2
vim /export/server/kafka/config/server.properties

broker.id=2
listeners=PLAINTEXT://node2:9092

### 编辑noed3
vim /export/server/kafka/config/server.properties

broker.id=3
listeners=PLAINTEXT://node3:9092

## 配置kafka环境变量
vim /etc/profile

export KAFKA_HOME=/export/server/kafka
export PATH=$PATH:$KAFKA_HOME/bin

刷新环境变量
source /etc/profile


## 启动
前台执行
/export/server/kafka/bin/kafka-server-start.sh /export/server/kafka/config/server.properties 

后台执行
nohup /export/server/kafka/bin/kafka-server-start.sh /export/server/kafka/config/server.properties 2>&1 >> /export/server/kafka/kafka.log &

日志存放位置
cd /export/server/kafka
tail -f kafka.log // 查看日志

## 测试kafka能否正常使用
/export/server/kafka/bin/kafka-topics.sh --create --zookeeper node1:2181 --replication-factor 3 --partitions 1 --topic test

启动一个生产者
/export/server/kafka/bin/kafka-console-producer.sh --broker-list node1:9092 --topic test

启动一个消费者
/export/server/kafka/bin/kafka-console-consumer.sh --bootstrap-server node1:9092 --topic test --from-beginning


# 安装Hadoop
法一:
windows下载在传入liunx

法二:
wget https://archive.apache.org/dist/hadoop/common/hadoop-3.3.0/hadoop-3.3.0.tar.gz

解压
tar -zxvf hadoop-3.3.0.tar.gz

构建软连接
ln -s /export/server/hadoop-3.3.0 /export/server/hadoop


编辑env.sh文件
vim hadoop-env.sh

# 配置java安装路径
export JAVA_HOME=/export/server/jdk
# 配置hadoop安装路径
export HADOOP_HOME=/export/server/hadoop
# Hadoop hdfs配置文件路径
export HADOOP_CONF_DIR=$HADOOP_HOME/etc/hadoop
# Hadoop YARN配置文件路径
export YARN_CONF_DIR=$HADOOP_HOME/etc/hadoop
# Hadoop YARN日志文件夹
export YARN_LOG_DIR=$HADOOP_HOME/logs/yarn
# Hadoop HDFS日志文件夹
export HDFS_LOG_DIR=$HADOOP_HOME/logs/hdfs

# Hadoop的使用启动用户配置
export HDFS_NAMENODE_USER=root
export HDFS_DATANODE_USER=root
export HDFS_SECONDARYNAMENODE_USER=root
export YARN_RESOURCEMANAGER_USER=root
export YARN_NODEMANAGER_USER=root
export YARN_PROXYSERVER_USER=root


编辑core-site.xml文件
vim core-site.xml 

 <property>
    <name>fs.defaultFS</name>
    <value>hdfs://node1:8020</value>
    <description></description>
 </property>

 <property>
    <name>io.file.buffer.size</name>
    <value>131072</value>
    <description></description>
 </property>


编辑hdfs-site.xml文件
vim hdfs-site.xml

 <property>
    <name>dfs.datanode.data.dir.perm</name>
    <value>700</value>
 </property>

<property>
  <name>dfs.namenode.name.dir</name>
  <value>/data/nn/</value>
  <description>Path on the local filesystem where the NameNode stores the namespace an ransactions logs persistently.</description>
</property>

<property>
  <name>dfs.namenode.hosts</name>
  <value>node1,node2,node3</value>
  <description>List of permitted DataNodes.</description>
</property>

<property>
  <name>dfs.blocksize</name>
  <value>268435456</value>
  <description></description>
</property>

<property>
  <name>dfs.namenode.handler.count</name>
  <value>100</value>
  <description></description>
</property>

<property>
  <name>dfs.datanode.data.dir</name>
  <value>/data/dn</value>
</property> 


编辑mapred-env.sh文件
vim mapred-env.sh

export JAVA_HOME=/export/server/jdk
export HADOOP_JOB_HISTORYSERVER_HEAPSIZE=1000
export HADOOP_MAPRED_ROOT_LOGGER=INFO,RFA


编辑mapred-site.xml文件
vim mapred-site.xml

<configuration>
  <property>
    <name>mapreduce.framework.name</name>
    <value>yarn</value>
    <description></description>
  </property>

  <property>
    <name>mapreduce.jobhistory.address</name>
    <value>node1:10020</value>
    <description></description>
  </property>

  <property>
    <name>mapreduce.jobhistory.webapp.address</name>
    <value>node1:19888</value>
    <description></description>
  </property>

  <property>
    <name>mapreduce.jobhistory.intermediate-done-dir</name>
    <value>/data/mr-history/tmp</value>
    <description></description>
  </property>

  <property>
    <name>mapreduce.jobhistory.done-dir</name>
    <value>/data/mr-history/done</value>
    <description></description>
  </property>

  <property>
    <name>yarn.app.mapreduce.am.env</name>
    <value>HADOOP_MAPRED_HOME=$HADOOP_HOME</value>
  </property>

  <property>
  <name>mapreduce.reduce.env</name>
  <value>HADOOP_MAPRED_HOME=$HADOOP_HOME</value>
</property> 
</configuration> 


编辑yarn-env.sh文件
vim yarn-env.sh 

export JAVA_HOME=/export/server/jdk
export HADOOP_HOME=/export/server/hadoop
export HADOOP_CONF_DIR=$HADOOP_HOME/etc/hadoop
export YARN_CONF_DIR=$HADOOP_HOME/etc/hadoop
export YARN_LOG_DIR=$HADOOP_HOME/logs/yarn
export HADOOP_LOG_DIR=$HADOOP_HOME/logs/hdfs 




编辑yarn-site.xml文件
vim yarn-site.xml

<property>
  <name>yarn.log.server.url</name>
  <value>http://node1:19888/jobhistory/logs/</value>
  <description></description>
</property>

<property>
  <name>yarn.web-proxy.address</name>
  <value>node1:8089</value>
  <description>proxy server hostname and port</description>
</property>

<property>
  <name>yarn.log-aggregation-enable</name>
  <value>true</value>
  <description>Configuration to enable or disable log aggregation</description>
</property>

<property>
  <name>yarn.nodemanager.remote-app-log-dir</name>
  <value>/tmp/logs</value>
  <description>Configuration to enable or disable log aggregation</description>
</property>

<!-- Site specific YARN configuration properties -->
<property>
  <name>yarn.resourcemanager.hostname</name>
  <value>node1</value>
  <description></description>
</property>

<property>
  <name>yarn.resourcemanager.scheduler.class</name>
  <value>org.apache.hadoop.yarn.server.resourcemanager.scheduler.fair.FairScheduler</value>
  <description></description>
</property> 

<property>
  <name>yarn.nodemanager.local-dirs</name>
  <value>/data/nm-local</value>
  <description>Comma-separated list of paths on the local filesystem where intermediate data is written.</description>
</property>

<property>
  <name>yarn.nodemanager.log-dirs</name>
  <value>/data/nm-logs</value>
  <description>Comma-separated list of paths on the local filesystem where logs are written.</description>
</property>

<property>
  <name>yarn.nodemanager.log.retain-seconds</name>
  <value>10800</value>
  <description>Default time (in seconds) to retain log files on the NodeManager. Only applicable if log-aggregation is disabled.</description>
</property>

<property>
  <name>yarn.nodemanager.aux-services</name>
  <value>mapreduce_shuffle</value>
  <description>Shuffle service that needs to be set for Map Reduce applications.</description>
</property>

修改workers文件
vim workers 

node1
node2
node3

scp -r hadoop-3.3.0 node2:`pwd`/ // 将hadoop-3.3.0文件复制到node2节点
scp -r hadoop-3.3.0 node3:`pwd`/ // 将hadoop-3.3.0文件复制到node3节点

ln -s /export/server/hadoop-3.3.0 /export/server/hadoop // 在每个节点上创建软链接

## 创建所需目录
在node1节点上创建所需目录
mkdir -p /data/nn
mkdir -p /data/dn
mkdir -p /data/nm-log
mkdir -p /data/nm-local


在node2节点上创建所需目录
mkdir -p /data/dn
mkdir -p /data/nm-log
mkdir -p /data/nm-local

在node3节点上创建所需目录
mkdir -p /data/dn
mkdir -p /data/nm-log
mkdir -p /data/nm-local


## 在三个节点上,配置hadoop
vim /etc/profile

export HADOOP_HOME=/export/server/hadoop
export PATH=$PATH:$HADOOP_HOME/bin:$HADOOP_HOME/sbin

刷新环境变量
source /etc/profile


## 格式化
hadoop namenode -format

启动hadoop
start-dfs.sh // 启动hadoop
stop-dfs.sh // 关闭hadoop

启动yarn
start-yarn.sh // 启动yarn
stop-yarn.sh // 关闭yarn

启动历史服务器
mapred --daemon start historyserver // 启动历史服务器
mapred --daemon stop historyserver // 关闭历史服务器

代理服务启动
yarn-daemon.sh start proxyserver // 启动代理服务
yarn-daemon.sh stop proxyserver // 关闭代理服务


浏览器验证
http://192.168.66.181:9870 // 验证HDFS

http://192.168.66.181:8088 // 验证YARM 


验证

vim test.txt // 创建文件

hadoop fs -put test.txt /test.txt // 上传文件
hadoop fs -ls / // 查看hadoop中的文件


验证YARM是否正常使用
vim words.txt // 创建文件

hadoop fs -put words.txt /words.txt 

hadoop jar /export/server/hadoop/share/hadoop/mapreduce/hadoop-mapreduce-examples-3.3.0.jar wordcount -Dmapred.job.queue.name=root.root /words.txt /output 












