# Docker容器技术实践

>数据科学与计算机学院
>
>17343012 陈泰霖

## 准备docker环境

docker环境的安装可以从[官方教程](https://docs.docker.com/install/linux/docker-ce/ubuntu/https://docs.docker.com/install/linux/docker-ce/ubuntu/)中获得

在命令行中输入以下命令

```
root@ubuntu:~# sudo apt-get update
root@ubuntu:~# sudo apt-get install \
>     apt-transport-https \
>     ca-certificates \
>     curl \
>     gnupg-agent \
>     software-properties-common
root@ubuntu:~# curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
root@ubuntu:~# sudo apt-key fingerprint 0EBFCD88
```

核对fingerprint是否与以下一致

```
pub   rsa4096 2017-02-22 [SCEA]
      9DC8 5822 9FC7 DD38 854A  E2D8 8D81 803C 0EBF CD88
uid           [ unknown] Docker Release (CE deb) <docker@docker.com>
sub   rsa4096 2017-02-22 [S]
```

输入以下命令下载docker社区版

```
root@ubuntu:~# apt-get install docker-ce docker-ce-cli containerd.io
```

查看docker版本

```
root@ubuntu:~# docker version
Client: Docker Engine - Community
 Version:           19.03.5
 API version:       1.40
 Go version:        go1.12.12
 Git commit:        633a0ea838
 Built:             Wed Nov 13 07:29:52 2019
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          19.03.5
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.12.12
  Git commit:       633a0ea838
  Built:            Wed Nov 13 07:28:22 2019
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.2.10
  GitCommit:        b34a5c8af56e510852c35414db4c1f4fa6172339
 runc:
  Version:          1.0.0-rc8+dev
  GitCommit:        3e425f80a8c931f88e6d94a8c831b9d5aa481657
 docker-init:
  Version:          0.18.0
  GitCommit:        fec3683

```

## 运行第一个容器

```
root@ubuntu:~# docker run hello-world

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```

可以看到第一个容器`hello-world`运行成功

## mysql与容器化

##### 拉取mysql容器

```
root@ubuntu:~# docker pull mysql:5.7
5.7: Pulling from library/mysql
Digest: sha256:5779c71a4730da36f013a23a437b5831198e68e634575f487d37a0639470e3a8
Status: Image is up to date for mysql:5.7
docker.io/library/mysql:5.7
```

##### 构建docker镜像

首先创建Dockerfile，录入以下内容

```
FROM ubuntu
ENTRYPOINT ["top", "-b"]
CMD ["-c"]
```

构建镜像

```
docker build . -t hello
```

运行镜像

```
docker run -it --rm hello -H
```

实践

```
root@ubuntu:/home/linux/Documents/DockerTest# cat Dockerfile 
FROM ubuntu
ENTRYPOINT ["top", "-b"]
CMD ["-c"]
root@ubuntu:/home/linux/Documents/DockerTest# docker build . -t hello
Sending build context to Docker daemon  3.584kB
Step 1/3 : FROM ubuntu
 ---> 775349758637
Step 2/3 : ENTRYPOINT ["top", "-b"]
 ---> Using cache
 ---> 1825d2c6e06c
Step 3/3 : CMD ["-c"]
 ---> Using cache
 ---> a83fd9a32660
Successfully built a83fd9a32660
Successfully tagged hello:latest
root@ubuntu:/home/linux/Documents/DockerTest# docker run -it --rm hello -H  

top - 13:49:00 up 21 min,  0 users,  load average: 0.39, 0.25, 0.29
Threads:   1 total,   1 running,   0 sleeping,   0 stopped,   0 zombie
%Cpu(s):  3.7 us,  4.3 sy,  0.1 ni, 91.5 id,  0.2 wa,  0.0 hi,  0.2 si,  0.0 st
KiB Mem :  5044900 total,   199184 free,  2748364 used,  2097352 buff/cache
KiB Swap:   969960 total,   968668 free,     1292 used.  1863644 avail Mem 

   PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND
     1 root      20   0   36480   2924   2572 R  0.0  0.1   0:00.16 top

```

##### 使用mysql服务器

启动服务器

```
root@ubuntu:/home/linux/Documents/DockerTest# docker run -p 3306:3306 --name mysql2 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
980b4ed17fb5d43e19f88003b69d07ce50a20d529a83728fe3d9086a1e82f822
root@ubuntu:/home/linux/Documents/DockerTest# docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                               NAMES
980b4ed17fb5        mysql:5.7           "docker-entrypoint.s…"   5 seconds ago       Up 4 seconds        0.0.0.0:3306->3306/tcp, 33060/tcp   mysql2
0adc42d42a1e        adminer             "entrypoint.sh docke…"   24 hours ago        Up 25 minutes       0.0.0.0:8080->8080/tcp              comptest_adminer_1

```

```
root@ubuntu:/home/linux/Documents/DockerTest# docker run -it --net host mysql:5.7 "sh"
# mysql -h127.0.0.1 -P3306 -uroot -proot
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 2
Server version: 5.7.28 MySQL Community Server (GPL)

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
```

mysql的数据库文件位置

```
root@ubuntu:/home/linux/Documents/DockerTest# docker exec -it mysql2 bash
root@980b4ed17fb5:/# ls var/lib/mysql
auto.cnf    ca.pem	     client-key.pem  ib_logfile0  ibdata1  mysql	       private_key.pem	server-cert.pem  sys
ca-key.pem  client-cert.pem  ib_buffer_pool  ib_logfile1  ibtmp1   performance_schema  public_key.pem	server-key.pem
```

## Docker网络

用`docker network ls`获得当前容器网络

```
root@ubuntu:/home/linux/Documents/DockerTest# docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
ebc903d583a6        bridge              bridge              local
a15e595360c4        comptest_default    bridge              local
96ec5f27f745        host                host                local
f6b94d408ddf        mynet               bridge              local
f2f397c88143        none                null                local
```

创建none网络

```
root@ubuntu:/home/linux/Documents/DockerTest# docker run -it --network=none busybox
Unable to find image 'busybox:latest' locally
latest: Pulling from library/busybox
322973677ef5: Pull complete 
Digest: sha256:1828edd60c5efd34b2bf5dd3282ec0cc04d47b2ff9caa0b6d4f07a21d1c08084
Status: Downloaded newer image for busybox:latest
/ # ifconfig
lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

/ # exit
```

创建host网络

```
root@ubuntu:/home/linux/Documents/DockerTest# docker run -it --network=host busybox
/ # ip l
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: ens33: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel qlen 1000
    link/ether 00:0c:29:e3:f1:5f brd ff:ff:ff:ff:ff:ff
3: br-f6b94d408ddf: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue 
    link/ether 02:42:71:9c:59:ce brd ff:ff:ff:ff:ff:ff
4: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue 
    link/ether 02:42:01:bc:a6:7c brd ff:ff:ff:ff:ff:ff
5: br-a15e595360c4: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue 
    link/ether 02:42:00:60:49:0c brd ff:ff:ff:ff:ff:ff
9: veth0f6ff18@if8: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue master br-a15e595360c4 
    link/ether 76:c9:23:bd:67:62 brd ff:ff:ff:ff:ff:ff
17: veth7d670d9@if16: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue master docker0 
    link/ether 36:00:db:9a:53:28 brd ff:ff:ff:ff:ff:ff
/ # hostname
ubuntu
```

创建bridge网络

```
root@ubuntu:/home/linux/Documents/DockerTest# docker network create --driver bridge --subnet 172.22.16.0/24 --gateway 172.22.16.1 my_net2
db223fed5a616039874c537191cb2a1d33022977bd00f92bccdad263e0ad2773
root@ubuntu:/home/linux/Documents/DockerTest# brctl show
bridge name	bridge id		STP enabled	interfaces
br-144585457ab7		8000.02426bdfdd3c	no		
br-a15e595360c4		8000.02420060490c	no		veth0f6ff18
br-db223fed5a61		8000.024200a25d10	no		
br-f6b94d408ddf		8000.0242719c59ce	no		
docker0		8000.024201bca67c	no		veth7050a93
							veth7d670d9
							veth87163d9
root@ubuntu:/home/linux/Documents/DockerTest# docker run -it --network=my_net2 busybox
/ # ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
24: eth0@if25: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue 
    link/ether 02:42:ac:16:10:02 brd ff:ff:ff:ff:ff:ff
    inet 172.22.16.2/24 brd 172.22.16.255 scope global eth0
       valid_lft forever preferred_lft forever
/ # 
```

设置了my_net2网络，子网掩码是172.22.16.0/24，网关是172.22.16.1，可以看到分配的ip地址是

172.22.16.2

```
root@ubuntu:/home/linux/Documents/DockerTest# docker run --name unet -it --rm ubuntu bash
root@c2051e639771:/# apt-get update
root@c2051e639771:/# apt-get install net-tools
root@c2051e639771:/# apt-get install iputils-ping -y
root@c2051e639771:/# ifconfig
root@c2051e639771:/# ping 192.168.58.130
PING 192.168.58.130 (192.168.58.130) 56(84) bytes of data.
64 bytes from 192.168.58.130: icmp_seq=1 ttl=64 time=0.119 ms
64 bytes from 192.168.58.130: icmp_seq=2 ttl=64 time=0.106 ms
64 bytes from 192.168.58.130: icmp_seq=3 ttl=64 time=0.105 ms
64 bytes from 192.168.58.130: icmp_seq=4 ttl=64 time=0.110 ms
```

发现可以ping通主机

## Docker容器

搭建私有容器仓库

创建本地仓库并将复制镜像到仓库

```
root@ubuntu:~# docker run -d -p 5000:5000 --restart=always --name registry registry:2
Unable to find image 'registry:2' locally
2: Pulling from library/registry
c87736221ed0: Pull complete 
1cc8e0bb44df: Pull complete 
54d33bcb37f5: Pull complete 
e8afc091c171: Pull complete 
b4541f6d3db6: Pull complete 
Digest: sha256:8004747f1e8cd820a148fb7499d71a76d45ff66bac6a29129bfdbfdc0154d146
Status: Downloaded newer image for registry:2
1d488816202230a80a789515f1309e5a1e37d3f7d45784a25392c9905d6690d1

root@ubuntu:~# docker pull ubuntu:16.04
16.04: Pulling from library/ubuntu
976a760c94fc: Pull complete 
c58992f3c37b: Pull complete 
0ca0e5e7f12e: Pull complete 
f2a274cc00ca: Pull complete 
Digest: sha256:e10375c69cf9e22989c82b0a87c932a21e33619ee322d6c7ce6a61456f54c30c
Status: Downloaded newer image for ubuntu:16.04
docker.io/library/ubuntu:16.04

root@ubuntu:~# docker tag ubuntu:16.04 localhost:5000/my-ubuntu
root@ubuntu:~# docker push localhost:5000/my-ubuntu
The push refers to repository [localhost:5000/my-ubuntu]
aa7f8c8d5f39: Pushed 
48817fbd6c92: Pushed 
1b039d138968: Pushed 
7082d7d696f8: Pushed 
```

将本地的`ubuntu:16.04`删除后，可以从仓库pull下来

```
root@ubuntu:~# docker image remove ubuntu:16.04
Untagged: ubuntu:16.04
Untagged: ubuntu@sha256:e10375c69cf9e22989c82b0a87c932a21e33619ee322d6c7ce6a61456f54c30c

root@ubuntu:~# docker image remove localhost:5000/my-ubuntu
Untagged: localhost:5000/my-ubuntu:latest
Untagged: localhost:5000/my-ubuntu@sha256:b1c268ca7c73556456ffc3318eb2a8e7ac6ad257ef5788d50dc1db4a3e3bd2ac
Deleted: sha256:56bab49eef2ef07505f6a1b0d5bd3a601dfc3c76ad4460f24c91d6fa298369ab
Deleted: sha256:384da4653311be75d58f084c09c91d73a66d8a259c951dcea49bb3daa3bc0923
Deleted: sha256:ea1e2a1c0567349cfcbeb1026e71a5cb7af452cf0aada89a8e97653ccaeac24d
Deleted: sha256:beae6f0d13ffdc05c24401888e6cba79cdbfd4beeaaee7c80e631cf52fab721d
Deleted: sha256:7082d7d696f8489bc9030e119acc56e210b6314bc8ac91aa69ed11c57c9243ba

root@ubuntu:~# docker pull localhost:5000/my-ubuntu
Using default tag: latest
latest: Pulling from my-ubuntu
976a760c94fc: Pull complete 
c58992f3c37b: Pull complete 
0ca0e5e7f12e: Pull complete 
f2a274cc00ca: Pull complete 
Digest: sha256:b1c268ca7c73556456ffc3318eb2a8e7ac6ad257ef5788d50dc1db4a3e3bd2ac
Status: Downloaded newer image for localhost:5000/my-ubuntu:latest
localhost:5000/my-ubuntu:latest

```

