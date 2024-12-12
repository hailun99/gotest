# 镜像来拉取网站 
```
https://gitee.com/wanfeng789/docker-hub
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

查看版本
```
docker version

docker -v  简洁版
```

# Docker命令
命令帮助
```
docker --help
```

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

重新加载系统配置
```
systemctl daemon-reload  
```

重启docker服务
```
systemctl restart docker  重启docker服务
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
docker rim $(docker images -q) 删除本地所有镜像
docker rm -f $(docker ps -aq) 删除搜索的所有镜像

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

## 提交镜像
提交
```
docker commit -m"update index.html" mynginx  mynginx:v1.0
-a 作者
-c 有哪些改变的页表
-m 此次提交的信息
-p 打包期间暂停运行
```

保存
```
docker save -o mynginx.tar mynginx:v1.0
-o 打包成文件
```

加载
```
docker ioad -i mynginx.tar
```

## 分享社区
登录
```
docker login
```

改名
```
docker tag mynginx:v1.0 [doncker用户]/mynginx:latcst(v1.0)
```

推送镜像
```
docker push [doncker用户]/mynginx:v1.0

docker push [doncker用户]/mynginx:latcst
```


## 存储 - 目录挂载
下载官方镜像
```
ocker run -d -p 80:80 --name app01 nginx
```

进入容器
```
docker exec -it 66bb bash
```

修改文件
```
cd /usr/share/nginx/html/
ls
echo 1111 > index.html
```

目录挂载
```
docker run -d -p 80:80 -v /app/nghtml:/usr/share/nginx/html --name app01 nginx
```

## 存储 - 卷映射
```
docker run -d -p 99:80 -v /app/nghtml:/usr/share/nginx/html -v ngconf:/etc/nginx --name app03 nginx
```

存储位置
```
cd /var/lib/docker/volumes/ngconf/_data/
```

修改文件内容
```
vi nginx.conf 
```

进入容器
```
docker exec -it app03 bash
```

查看修改内容
```
cd /etc/nginx/
cat nginx.conf 
```

查看所有的卷
```
docker volume ls
```

创建卷
```
docker volume create haha
```

查看卷的详情
```
docker volume inspecot ngconf
```

## Docker网络
查看容器里的细节
```
docker container inspect app1

docker inspect app2
```

### 访问别的机器 法1
进入app1
```
docker exec -it bash
```

ip地址访问app2
```
curl http://172.17.0.3:80
```

### 访问别的机器 法2
创建网络
```
docker network create mynet
```

查看网络
```
docker network ls
```

进入app1
```
docker exec -it app1 bash
```

访问app2
```
curl http://app2:80
```

## redis主从同步集群
创建主节点
```
mkdir /app/rd1
```

授权
```
chmod -R 777 rd1
```

配置master
```
docker run -d -p 6379:6379 \
 -v /app/rd1:/bitnami/redis/data \
 -e REDIS_REPLICATION_MODE=master \
 -e REDIS_PASSWORD=123456 \
 --network mynet --name redis01 \
 bitnami/redis
```

创建从节点
```
mkdir /app/rd2
```

授权
```
chmod -R 777 rd1
```

配置slave
```
docker run -d -p 6380:6379 \
 -v /app/rd2:/bitnami/redis/data \
 -e REDIS_REPLICATION_MODE=slave \
 -e REDIS_MASTER_HOST=redis01 \
 -e REDIS_MASTER_PORT_NUMBER=6379 \
 -e REDIS_MASTER_PASSWORD=123456 \
 -e REDIS_PASSWORD=123456 \
 --network mynet --name redis02 \
 bitnami/redis
```

## 启动mysql
```
docker run -d -p 3306:3306 \
-v /app/myconf:/etc/mysql/conf.d \
-v /app/mydata:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=123456 \
mysql:8.0.37-debian
```

进入mysql容器
```
docker exec -it peaceful_driscoll bash
```

登录mysql
```
mysql -u root -p
```

定位容器位置
```
whereis mysql
```

## 命令式安装
#创建网络
```
docker network create blog
```

启动mysql
```
docker run -d -p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=123456 \
-e MYSQL_DATABASE=wordpress \
-v mysql-data:/var/lib/mysql \
-v /app/myconf:/etc/mysql/conf.d \
--restart always --name mysql \
--network blog \
mysql:8.0
```

启动wordpress
```
docker run -d -p 8080:80 \
-e WORDPRESS_DB_HOST=mysql \
-e WORDPRESS_DB_USER=root \
-e WORDPRESS_DB_PASSWORD=123456 \
-e WORDPRESS_DB_NAME=wordpress \
-v wordpress:/var/www/html \
--restart always --name wordpress-app \
--network blog \
wordpress:latestdebian
```

## Docker Compose
(可查考Docker官方文档)

修改文件
```
vi Compose.yaml
```

配置Compose.yaml
```
name: myblog
services:
  mysql:
    container_name: mysql
    image: mysql:8.0
    ports:
      - "3306:3306"
    environment:
      - MYSQL_ROOT_PASSWORD=123456
      - MYSQL_DATABASE=wordpress
    volumes:
      - mysql-data:/var/lib/mysql
      - /app/myconf:/etc/mysql/conf.d
    restart: always
    networks:
      - blog

  wordpress:
    image: wordpress
    ports:
      - "8080:80"
    environment:
      WORDPRESS_DB_HOST: mysql
      WORDPRESS_DB_USER: root
      WORDPRESS_DB_PASSWORD: 123456
      WORDPRESS_DB_NAME: wordpress
    volumes:
      - wordpress:/var/www/html
    restart: always
    networks:
      - blog
    depends_on:
      - mysql

volumes:
  mysql-data:
  wordpress:

networks:
  blog:
```

启动文件
```
docker compose -f compose.yaml up -d
```

删除应用
```
docker compose -f compose.yaml down

docker compose -f compose.yaml down --rmi -v

docker compose -f compose.yaml down --rmi all -v 本地的都删除
--rmi 移除镜像
-v 移除卷
```

特性
```
 增量更新
   修改 Docker Compose 文件。重新启动应用。只会触发修改项的重新启动。
 数据不删
   默认就算down了容器，所有挂载的卷不会被移除。比较安全
```

## 多文件安装
关闭内存分页
```
sudo swapoff -a
```

打开文件
```
sudo vi /etc/sysctl.conf
```

加入内容
```
vm.max_map_count=262144
```

是文件生效
```
sudo sysctl -p
```

查看修改
```
cat /proc/sys/vm/max_map_count
```

修改文件
```
vi Compose.yaml
```

配置Compose.yaml
```
name: devsoft
services:
  redis:
    image: bitnami/redis:latest
    restart: always
    container_name: redis
    environment:
      - REDIS_PASSWORD=123456
    ports:
      - '6379:6379'
    volumes:
      - redis-data:/bitnami/redis/data
      - redis-conf:/opt/bitnami/redis/mounted-etc
      - /etc/localtime:/etc/localtime:ro

  mysql:
    image: mysql:8.0.31
    restart: always
    container_name: mysql
    environment:
      - MYSQL_ROOT_PASSWORD=123456
    ports:
      - '3306:3306'
      - '33060:33060'
    volumes:
      - mysql-conf:/etc/mysql/conf.d
      - mysql-data:/var/lib/mysql
      - /etc/localtime:/etc/localtime:ro

  rabbit:
    image: rabbitmq:3-management
    restart: always
    container_name: rabbitmq
    ports:
      - "5672:5672"
      - "15672:15672"
    environment:
      - RABBITMQ_DEFAULT_USER=rabbit
      - RABBITMQ_DEFAULT_PASS=rabbit
      - RABBITMQ_DEFAULT_VHOST=dev
    volumes:
      - rabbit-data:/var/lib/rabbitmq
      - rabbit-app:/etc/rabbitmq
      - /etc/localtime:/etc/localtime:ro
  opensearch-node1:
    image: opensearchproject/opensearch:2.13.0
    container_name: opensearch-node1
    environment:
      - cluster.name=opensearch-cluster # Name the cluster
      - node.name=opensearch-node1 # Name the node that will run in this container
      - discovery.seed_hosts=opensearch-node1,opensearch-node2 # Nodes to look for when discovering the cluster
      - cluster.initial_cluster_manager_nodes=opensearch-node1,opensearch-node2 # Nodes eligibile to serve as cluster manager
      - bootstrap.memory_lock=true # Disable JVM heap memory swapping
      - "OPENSEARCH_JAVA_OPTS=-Xms512m -Xmx512m" # Set min and max JVM heap sizes to at least 50% of system RAM
      - "DISABLE_INSTALL_DEMO_CONFIG=true" # Prevents execution of bundled demo script which installs demo certificates and security configurations to OpenSearch
      - "DISABLE_SECURITY_PLUGIN=true" # Disables Security plugin
    ulimits:
      memlock:
        soft: -1 # Set memlock to unlimited (no soft or hard limit)
        hard: -1
      nofile:
        soft: 65536 # Maximum number of open files for the opensearch user - set to at least 65536
        hard: 65536
    volumes:
      - opensearch-data1:/usr/share/opensearch/data # Creates volume called opensearch-data1 and mounts it to the container
      - /etc/localtime:/etc/localtime:ro
    ports:
      - 9200:9200 # REST API
      - 9600:9600 # Performance Analyzer

  opensearch-node2:
    image: opensearchproject/opensearch:2.13.0
    container_name: opensearch-node2
    environment:
      - cluster.name=opensearch-cluster # Name the cluster
      - node.name=opensearch-node2 # Name the node that will run in this container
      - discovery.seed_hosts=opensearch-node1,opensearch-node2 # Nodes to look for when discovering the cluster
      - cluster.initial_cluster_manager_nodes=opensearch-node1,opensearch-node2 # Nodes eligibile to serve as cluster manager
      - bootstrap.memory_lock=true # Disable JVM heap memory swapping
      - "OPENSEARCH_JAVA_OPTS=-Xms512m -Xmx512m" # Set min and max JVM heap sizes to at least 50% of system RAM
      - "DISABLE_INSTALL_DEMO_CONFIG=true" # Prevents execution of bundled demo script which installs demo certificates and security configurations to OpenSearch
      - "DISABLE_SECURITY_PLUGIN=true" # Disables Security plugin
    ulimits:
      memlock:
        soft: -1 # Set memlock to unlimited (no soft or hard limit)
        hard: -1
      nofile:
        soft: 65536 # Maximum number of open files for the opensearch user - set to at least 65536
        hard: 65536
    volumes:
      - /etc/localtime:/etc/localtime:ro
      - opensearch-data2:/usr/share/opensearch/data # Creates volume called opensearch-data2 and mounts it to the container

  opensearch-dashboards:
    image: opensearchproject/opensearch-dashboards:2.13.0
    container_name: opensearch-dashboards
    ports:
      - 5601:5601 # Map host port 5601 to container port 5601
    expose:
      - "5601" # Expose port 5601 for web access to OpenSearch Dashboards
    environment:
      - 'OPENSEARCH_HOSTS=["http://opensearch-node1:9200","http://opensearch-node2:9200"]'
      - "DISABLE_SECURITY_DASHBOARDS_PLUGIN=true" # disables security dashboards plugin in OpenSearch Dashboards
    volumes:
      - /etc/localtime:/etc/localtime:ro
  zookeeper:
    image: bitnami/zookeeper:3.9
    container_name: zookeeper
    restart: always
    ports:
      - "2181:2181"
    volumes:
      - "zookeeper_data:/bitnami"
      - /etc/localtime:/etc/localtime:ro
    environment:
      - ALLOW_ANONYMOUS_LOGIN=yes

  kafka:
    image: 'bitnami/kafka:3.4'
    container_name: kafka
    restart: always
    hostname: kafka
    ports:
      - '9092:9092'
      - '9094:9094'
    environment:
      - KAFKA_CFG_NODE_ID=0
      - KAFKA_CFG_PROCESS_ROLES=controller,broker
      - KAFKA_CFG_LISTENERS=PLAINTEXT://:9092,CONTROLLER://:9093,EXTERNAL://0.0.0.0:9094
      - KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092,EXTERNAL://192.168.66.137:9094
      - KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,EXTERNAL:PLAINTEXT,PLAINTEXT:PLAINTEXT
      - KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=0@kafka:9093
      - KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER
      - ALLOW_PLAINTEXT_LISTENER=yes
      - "KAFKA_HEAP_OPTS=-Xmx512m -Xms512m"
    volumes:
      - kafka-conf:/bitnami/kafka/config
      - kafka-data:/bitnami/kafka/data
      - /etc/localtime:/etc/localtime:ro
  kafka-ui:
    container_name: kafka-ui
    image: provectuslabs/kafka-ui:latest
    restart: always
    ports:
      - 8080:8080
    environment:
      DYNAMIC_CONFIG_ENABLED: true
      KAFKA_CLUSTERS_0_NAME: kafka-dev
      KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS: kafka:9092
    volumes:
      - kafkaui-app:/etc/kafkaui
      - /etc/localtime:/etc/localtime:ro

  nacos:
    image: nacos/nacos-server:v2.3.1
    container_name: nacos
    ports:
      - 8848:8848
      - 9848:9848
    environment:
      - PREFER_HOST_MODE=hostname
      - MODE=standalone
      - JVM_XMX=512m
      - JVM_XMS=512m
      - SPRING_DATASOURCE_PLATFORM=mysql
      - MYSQL_SERVICE_HOST=nacos-mysql
      - MYSQL_SERVICE_DB_NAME=nacos_devtest
      - MYSQL_SERVICE_PORT=3306
      - MYSQL_SERVICE_USER=nacos
      - MYSQL_SERVICE_PASSWORD=nacos
      - MYSQL_SERVICE_DB_PARAM=characterEncoding=utf8&connectTimeout=1000&socketTimeout=3000&autoReconnect=true&useUnicode=true&useSSL=false&serverTimezone=Asia/Shanghai&allowPublicKeyRetrieval=true
      - NACOS_AUTH_IDENTITY_KEY=2222
      - NACOS_AUTH_IDENTITY_VALUE=2xxx
      - NACOS_AUTH_TOKEN=SecretKey012345678901234567890123456789012345678901234567890123456789
      - NACOS_AUTH_ENABLE=true
    volumes:
      - /app/nacos/standalone-logs/:/home/nacos/logs
      - /etc/localtime:/etc/localtime:ro
    depends_on:
      nacos-mysql:
        condition: service_healthy
  nacos-mysql:
    container_name: nacos-mysql
    build:
      context: .
      dockerfile_inline: |
        FROM mysql:8.0.31
        ADD https://raw.githubusercontent.com/alibaba/nacos/2.3.2/distribution/conf/mysql-schema.sql /docker-entrypoint-initdb.d/nacos-mysql.sql
        RUN chown -R mysql:mysql /docker-entrypoint-initdb.d/nacos-mysql.sql
        EXPOSE 3306
        CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci"]
    image: nacos/mysql:8.0.30
    environment:
      - MYSQL_ROOT_PASSWORD=root
      - MYSQL_DATABASE=nacos_devtest
      - MYSQL_USER=nacos
      - MYSQL_PASSWORD=nacos
      - LANG=C.UTF-8
    volumes:
      - nacos-mysqldata:/var/lib/mysql
      - /etc/localtime:/etc/localtime:ro
    ports:
      - "13306:3306"
    healthcheck:
      test: [ "CMD", "mysqladmin" ,"ping", "-h", "localhost" ]
      interval: 5s
      timeout: 10s
      retries: 10
  prometheus:
    image: prom/prometheus:v2.52.0
    container_name: prometheus
    restart: always
    ports:
      - 9090:9090
    volumes:
      - prometheus-data:/prometheus
      - prometheus-conf:/etc/prometheus
      - /etc/localtime:/etc/localtime:ro

  grafana:
    image: grafana/grafana:10.4.2
    container_name: grafana
    restart: always
    ports:
      - 3000:3000
    volumes:
      - grafana-data:/var/lib/grafana
      - /etc/localtime:/etc/localtime:ro

volumes:
  redis-data:
  redis-conf:
  mysql-conf:
  mysql-data:
  rabbit-data:
  rabbit-app:
  opensearch-data1:
  opensearch-data2:
  nacos-mysqldata:
  zookeeper_data:
  kafka-conf:
  kafka-data:
  kafkaui-app:
  prometheus-data:
  prometheus-conf:
  grafana-data:
```

启动
```
 在 compose.yaml 文件所在的目录下执行
docker compose up -d
 等待启动所有容器
```