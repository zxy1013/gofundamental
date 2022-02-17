##### **Docker安装开发环境** 

###### **Docker安装Nginx** 

Nginx 是一个高性能的 HTTP 和反向代理 web 服务器，同时也提供了 IMAP/POP3/SMTP 服务 。 

**1查看可用的Nginx版本** 

访问 Nginx 镜像库地址： 

`https://hub.docker.com/_/nginx?tab=tags` 

**2、取最新版的Nginx镜像** 

`docker pull nginx:latest `

**3、查看本地镜像** 

使用以下命令来查看是否已安装了 nginx： 

`docker images`

**4、运行容器** 

安装完成后，我们可以使用以下命令来运行 nginx 容器 

`docker run -it --name nginx-80 --rm -d -p 80:80 nginx`

> 参数说明： 
>
> **--name nginx-test**：容器名称。 
>
> **-p 80:80**：端口进行映射，将本地 8080 端口映射到容器内部的 80 端口。 
>
> **-d nginx**： 设置容器在后台一直运行。 

**5、安装成功** 

最后我们可以通过浏览器可以直接访问 80 端口的 nginx 服务 

**6、自定义配置** 

`mkdir -p /usr/local/docker/nginx ` //创建指定目录配置 

`cd /usr/local/docker/nginx ` 

`mkdir conf`

`mkdir html`

`mkdir logs`

> **示例1静态资源配置：** 
>
> 静态资源直接访问。直接将容器内部静态资源文件复制到/usr/local/docker/nginx下即可。 
>
> `docker cp f4:/etc/nginx/nginx.conf /usr/local/docker/nginx/conf`
>
> `docker cp f4:/etc/nginx/conf.d/ /usr/local/docker/nginx/conf`
>
> \# 启动Nginx 静态资料配置 
>
> `docker run -it --name nginx-80 --rm -d -p 80:80 -v /usr/local/docker/nginx/html:/usr/share/nginx/html -v /usr/local/docker/nginx/nginx.conf:/etc/nginx/nginx.conf -v /usr/local/docker/nginx/conf.d/default.conf:/etc/nginx/conf.d/default.conf -v /usr/local/docker/nginx/logs:/var/log/nginx nginx` 
>
> `cd ..`
>
> `cd html`
>
> `vi index.html` 加入页面 访问 ip192.168.20.136:80
>
> **集群配置**
>
> **conf.d/default.conf** 
>
> ```
> server_name exam_qf;
> location / { 
> 	proxy_pass http://nginxCluster; 
> }
> ```
>
> **nginx.conf** 
>
> ```shell
> upstream nginxCluster{ 
>     server 192.168.20.136:8080; 
>     server 192.168.20.136:8081; 
> }
> server { 
>     listen 80; 
>     server_name localhost; 
>     # charset koi8-r; 
>     # access_log /var/log/nginx/host.access.log main; 	  location /{ 
>         proxy_pass http://nginxCluster; 
>     }
> }
> ```
>
> `docker restart nginx-80`



**Docker安装MySQL** 

###### **安装MySQL 5.\*版本** 

**1、搜索镜像** 

`docker search mysql `

**2、下载镜像** 

`docker pull mysql:5.6 `

**3、创建并启动MySQL容器**  -e 密码

`docker run -d --name mysql5.6-3306 -p 3306:3306 -e MYSQL_ROOT_PASSWORD='guoweixin' mysql:5.6 `

**4、访问测试** 进入到容器内部

`docker exec -it mysql5.6-3306 bash` 

连接mysql数据库： 

`mysql -u root -p `

输入数据库密码guoweixin即可完成。 

**5、授权其他机器登陆**

> 1、授权主机外部访问： 
>
> MySQL>`GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'guoweixin' WITH GRANT OPTION;` 
>
> 2、刷新权限： 
>
> MySQL>`FLUSH PRIVILEGES; `
>
> 3、退出： 
>
> MySQL>`EXIT;`
>
> 开启防火墙端口
>
> 开启端口: 
>
> `firewall-cmd --zone=public --add-port=3306/tcp --permanent `
>
> `firewall-cmd --reload` # 重启 



###### **安装MySQL 8.\*版本**

```
docker pull mysql

# 启动 
docker run -d --name mysql8 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=guoweixin mysql 

# 进入容器 
docker exec -it mysql8 bash 

# 登录mysql 
mysql -u root -p 
ALTER USER 'root'@'localhost' IDENTIFIED BY 'guoweixin'; 
```

参数说明： 

> **-p 3302:3306** ：映射容器服务的 3306 端口到宿主机的 3302 端口，外部主机可以直接通过 宿主机ip:3302访问到 MySQL 的服务。 
>
> **MYSQL_ROOT_PASSWORD=guoweixin**：设置 MySQL 服务 root 用户的密码。 



###### **Docker安装Redis** 

**安装单机版Redis** 

**1、搜索redis** 

`docker search redis `

**2、下载镜像** 

`docker pull redis:4.0.1 `

**3、创建并运行容器** 

`docker run --rm -d --name redis6379 -p 6379:6379 redis:4.0.1 --requirepass "guoweixin" `

**4、测试Redis** 

`docker exec -it redis6379 bash` // 进入redis命令 

`redis-cli` // 开启客户端功能 





##### **Docker定制镜像** 

当我们从docker镜像仓库中下载的镜像不能满足我们的需求时，我们可以通过以下两种方式对镜像进行更改。

> 1.从已经创建的容器中更新镜像，并且提交这个镜像
>
> 2.使用 Dockerfile 指令来创建一个新的镜像 

Dockerfile定制镜像 

> 1、对于开发人员，可以为开发团队提供一个完全一致的开发环境
>
> 2、对于测试人员，可以直接拿开发时所构建的镜像测试。 
>
> 3、对于运维人员，在部署时，可以实现快速部署、移值。 



###### **Dockerfile定制镜像** 

镜像的定制实际上就是定制每一层所添加的配置、文件。**如果我们可以把每一层修改、安装、构建、操作的命令都写入一个脚本，用这个脚本来构建、定制镜像，**那么之前提及的无法重复的问题、镜像构建透明性的问题、体积的问题就都会解决。这个脚本就是 Dockerfile。 

Dockerfile 是一个文本文件，其内包含了一条条的指令(Instruction)，每一条指令的内容，就在描述该层应当如何构建。 

###### **Dockerfile常用命令** 

> **FROM** 
>
> **--指定基础镜像** 
>
> 基础镜像不存在会在Docker Hub上拉去(一般会是文件的第一个指令) 使用格式： 
>
> FROM <镜像>:[tag] 
>
> FROM <镜像>@digest[校验码] 
>
> 当前主机没有此镜像时，会自动去官网HUB下载 
>
> 
>
> **ENV**
>
> ENV指令可以用于为docker容器设置环境变量 ENV设置的环境变量，可以使用 docker inspect命令来查看。同时还可以使用docker run --env =来修改环境变量。 
>
> ```shell
> 具体用法： 
> ENV JAVA_HOME /usr/local/jdk ENV JRE_HOME $JAVA_HOME/jre ENV CLASSPATH $JAVA_HOME/lib/:$JRE_HOME/lib/ ENV PATH $PATH:$JAVA_HOME/bin/
> ```
>
> 
>
> **USER** 
>
> 用来切换运行属主身份的。Docker 默认是使用 root，但若不需要，建议切换使用者身分，毕竟 root 权限太大了，使用上有安全的风险。 
>
> 
>
> **WORKDIR** 
>
> **WORKDIR用来切换工作目录的。** 
>
> Docker 默认的工作目录是/，如果想让其他指令在指定的目录下执行，就得靠 WORKDIR。相当于cd，WORKDIR 动作的目录改变是持久的，不用每个指令前都使用一次 WORKDIR。 
>
> `WORKDIR /usr/local/tomcat/`
>
> 
>
> **VOLUME** 
>
> 创建一个可以从本地主机或其他容器挂载的挂载点，一般用来存放数据库和需要保持的数据等。 
>
> --卷 
>
> 只能定义docker管理的卷： `VOLUME /data/mysql`运行的时候会随机在宿主机的目录下生成一个卷目录！ 
>
> 
>
> **COPY** 
>
> **--把宿主机中的文件复制到镜像中去！** 
>
> 文件要在Dockerfile工作目录 
>
> src 原文件 --支持通配符 --通常相对路径 
>
> dest 目标路径 --通常绝对路径 
>
> 
>
> **ADD** 
>
> **类似COPY命令** 
>
> ADD 将文件从路径 复制添加到容器内部路径 。 
>
> 必须是相对于源文件夹的一个文件或目录，也可以是一个远程的url。 
>
> 是目标容器中的绝对路径。 所有的新文件和文件夹都会创建UID 和 GID。事实上如果是一个远程文件URL，那么目标文件的权限将会是600。 
>
> 
>
> **EXPOSE** 
>
> **为容器打开指定要监听的端口以实现与外部通信** 
>
> 使用格式： 
>
> `EXPOSE 80/tcp 23/udp `
>
> 不加协议默认为tcp 
>
> 使用-P选项可以暴露这里指定的端口！但是宿主的关联至这个端口的端口是随机的！外部命令可以指定
>
> 
>
> **RUN** 
>
> RUN 指令是用来执行命令行命令的。由于命令行的强大能力，RUN 指令在定制镜像时是最常用的指令之一。其格式有两种： 
>
> > • shell 格式：`RUN <命令>`，就像直接在命令行中输入的命令一样。刚才写的 Dockerfile 中的 RUN 指令就是这种格式。 
> >
> > • exec 格式：RUN ["可执行文件", "参数1", "参数2"]，这更像是函数调用中的格式。使用格式： `RUN["","",""] `
>
> RUN 就像 Shell 脚本一样可以执行命令，那么我们是否就可以像 Shell 脚本一样把每个命令对应一个 RUN 呢？比如这样： 
>
> ```shell
> RUN apt-get update 
> RUN apt-get install -y gcc libc6-dev make 
> RUN wget http://download.redis.io/releases/redis-4.0.1.tar.gz 
> RUN tar xzf redis-4.0.1.tar.gz 
> RUN cd redis-4.0.1 # 错误 用WORKDIR
> ```
>
> Dockerfile 中每一个指令都会建立一层，RUN 也不例外。每一个 RUN 的行为，和刚才我们手工建立镜像的过程一样： 新建立一层，在其上执行这些命令，执行结束后，commit 这一层的修改，构成新的镜像。 而上面的这种写法，创建了多层镜像。这是完全没有意义的，而且很多运行时不需要的东西，都被装进了镜像里，比如编译环境、更新的软件包等等。结果就是产生非常臃肿、非常多层的镜像，不仅仅增加了构建部署的时间，也很容易出错。 这是很多初学 Docker 的人常犯的一个错误。 
>
> Union FS 是有最大层数限制的，比如 AUFS，曾经是最大不得超过 42 层，现在是不得超过 127 层。上面的 Dockerfile正确的写法应该是这样：
>
> ```
> FROM centos 
> RUN apt-get update \ 
>     && apt-get install -y gcc libc6-dev make \
>     && wget http://download.redis.io/releases/redis-4.0.1.tar.gz \ 
>     && tar xzf redis-4.0.1.tar.gz \ 
>     && cd redis-4.0.1
> ```
>
> 首先，之前所有的命令只有一个目的，就是编译、安装 redis 可执行文件。因此没有必要建立很多层，这只是一层的事情。因此，这里没有使用很多个 RUN 对一一对应不同的命令，而是仅仅使用一个 RUN 指令，并使用 && 将各个所需命令串联起来。将之前的 7 层，简化为了 1 层。在撰写 Dockerfile 的时候，要经常提醒自己，这并不是在写 Shell脚本，而是在定义每一层该如何构建。 并且，这里为了格式化还进行了换行。Dockerfile 支持 Shell 类的行尾添加 \ 的命令换行方式，以及行首 # 进行注释的格式。良好的格式，比如换行、缩进、注释等，会让维护、排障更为容易，这 是一个比较好的习惯。 



###### **案例1** 

需求：创建一个镜像（基于tomcat）里面要有一个index.html，并写入Hello qfnj Docker 

**1、在宿主机创建一空白目录** 

`mkdir -p /usr/local/docker/demo1 `

**2、在该目录下，创建一文件Dockerfile** 

`vim Dockerfile`

**3、其内容为：** 

```dockerfile
FROM tomcat //指定tomcat最新版本镜像

RUN mkdir -p /usr/local/tomcat/webapps/ROOT/
RUN echo 'Hello qfnj Docker'>/usr/local/tomcat/webapps/ROOT/index.html

WORKDIR /usr/local/tomcat/webapps/
```

**这个Dockerfile很简单，一共就4行。涉及到了两条指令，FROM和RUN。** 

**4、构建镜像** 

`docker build -t demo1 .`

**5、运行镜像所在容器** 

`docker run --rm --name demo1-8080 -p 8080:8080 -d demo1 `

访问浏览器即可成功访问



**构建镜像Build** 

回到之前定制的 tomcat 镜像的 Dockerfile 来。现在我们明白了这个Dockerfile 的内容，那么让我们来构建这个镜像吧。在 Dockerfile 文件所在目录执行： 

`docker build -t demo1 .` 

从命令的输出结果中，我们可以清晰的看到镜像的构建过程。 

docker build [选项] <上下文路径/URL/-> 

>  **-t** ：指定要创建的目标镜像名 
>
> **.** ：Dockerfile 文件所在目录，可以指定Dockerfile 的绝对路径 

在这里我们指定了最终镜像的名称，构建成功后，我们可以像之前运行 tomcat 那样来运行这个镜像，`docker run --rm --name demo1-8080 -p 8080:8080 -d demo1 `，其结果会和 tomcat 一样。



###### **案例2**

案例：基于上一个镜像（基于tomcat）将ROOT内多余的文件都删除。只保留index.html 

**1** **基于如上修改Dockerfile** 

```dockerfile
FROM tomcat //指定tomcat最新版本镜像 

// WORKDIR 用来切换工作目录的。而不是用RUN cd。 
WORKDIR /usr/local/tomcat/webapps/ROOT/ //切换到该目录下

RUN rm -rf * //将当前目录的文件都删掉 
RUN echo 'Hello qfnj Docker'>/usr/local/tomcat/webapps/ROOT/index.html 
```

**2、 构建镜像** 

`docker build –t 镜像名 . `//Dockerfile上下文路径 

**3、查看镜像列表docker images** 

如果镜像名称有 

**4、删除虚拟镜像** 

`docker image prune`



###### **案例3** 

案例：基于上一个镜像（基于tomcat）外部复制一个文件(图片)，并复制到容器中并能访问 

**1** **基于如上修改Dockerfile** 

```dockerfile
FROM tomcat //指定tomcat最新版本镜像 

WORKDIR /usr/local/tomcat/webapps/ROOT/ //切换到该目录下 

RUN rm -rf * //将当前目录的文件都删掉

// COPY <源路径>... <目标路径> 
COPY 1.png /usr/local/tomcat/webapps/ROOT/ 

RUN echo 'Hello qfnj Docker'>/usr/local/tomcat/webapps/ROOT/index.html 
```

**2、 构建镜像** 

`docker build –t 镜像名 .` //Dockerfile上下文路径COPY 格式： 

> • COPY <源路径>... <目标路径> 
>
> • COPY ["<源路径1>",... "<目标路径>"] 

和 RUN 指令一样，也有两种格式，一种类似于命令行，一种类似于函数调用。 

<源路径> 的文件/目录复制到新的一层的镜像内的 <目标路径> 位置。比如： `COPY qfjy.png /usr/local/tomcat/webapps/ROOT/ `

<目标路径> 可以是容器内的绝对路径，也可以是相对于工作目录的相对路径（工作目录可以用 WORKDIR指令来指定）。目标路径不需要事先创建，如果目录不存在会在复制文件前先行创建缺失目录。 

此外，还需要注意一点，使用 COPY 指令，源文件的各种元数据都会保留。比如读、写、执行权限、文件变更时间等。这个特性对于镜像定制很有用。特别是构建相关文件都在使用 Git 进行管理的时候。 

`docker run --rm --name demo1-8080 -p 8080:8080 -d 镜像名`



##### **Docker仓库搭建** 

**问题描述**

在我们普通的 docker pull 的过程,都是从hub.docker.com 进行镜像的拉取。但是这个有一个问题,在公司的内部项目中如果push 上去,那么就会被其他的人看到，这个显然是不允许的。 

就好比很多公司不会把项目代码放到github上面一样,他们会在自己的内网搭建gitlab服务器，好在docker已经考虑到这一点,提供好了。

###### **Docker官方Registry** 

**拉取镜像仓库** 

`docker pull registry `

**查看所有镜像** 

`docker images `

**启动镜像服务器registry** 

首先在主机上新建一个目录，供存储镜像。由于Registry是一个镜像，运行后若我们删除了容器，容器里面的资源就会丢失，所以我们在运行时，指定一个资源的挂载目录，映射到宿主的一个目录下，这样资源就不会丢失了。 

`cd /usr/local/ `

`mkdir docker_registry `

`docker run -d -p 5000:5000 --name=guoweixinregistry --restart=always --privileged=true -v /usr/local/docker_registry:/var/lib/registry docker.io/registry`

> -p 5000:5000 端口 
>
> --name=guoweixinregistry 运行的容器名称
>
> --restart=always 自动重启 
>
> --privileged=true centos7中的安全模块selinux把权限禁止了，加上这行是给容器增加执行权限 
>
> -v /usr/local/docker_registry:/var/lib/registry 把主机的/usr/local/docker_registry 目录挂载到registry容器的/var/lib/registry目录下，假如有删除容器操作，我们的内容也不会被删除 
>
> docker.io/registry 镜像名称



###### **测试** 

**从公有镜像仓库中下载一个镜像下来，或本地构建镜像。然后push到私有仓库进行测试** 

```shell
# 利用tag 标记一个新镜像 
docker tag exam 127.0.0.1:5000/exam 
# 推送镜像
docker push 127.0.0.1:5000/exam
```

**此时，访问浏览器私有仓库地址：http://宿主IP:5000/v2/_catalog， 即可看见推送的镜像信息了** `http://192.168.20.135:5000/v2/_catalog`



###### **IP地址提交** 

直接使用127.0.0.1或者local时，是没有进行安全检验的。 

当我们使用外部的ip地址推送时，Registry为了安全性考虑，默认是需要https证书支持的。 

```shell
# 利用tag 标记一个新镜像 
docker tag demo1 192.168.20.135:5000/demo1 
# 通过IP地址进行push提交 
docker push 192.168.20.135:5000/demo1
```

**错误提示：** 

```
The push refers to repository [192.168.20.135:5000/demo1] 

Get https://192.168.20.135:5000/v2/: http: server gave HTTP response to HTTPS client 
```

**解决方案:** 

一种是通过daemon.json配置一个insecure-registries属性； 

另一种就直接配置一个https的证书了。 

在 /etc/docker/daemon.json 中写入如下内容 

```shell
"insecure-registries": ["实际的ip:端口"] 
{
"insecure-registries": ["192.168.20.135:5000"], 
"registry-mirrors": ["https://gxeo3yz7.mirror.aliyuncs.com"] 
} 
```

insecure-registries----->开放注册https协议 

registry-mirrors----->仓库源 

**重启Docker:** 

```shell
sudo systemctl daemon-reload 
sudo systemctl restart docker 
```

**拉取私有仓库镜像** 

`docker pull IP地址:端口号/镜像名称 `



##### **Harbor介绍及实践** 

docker 官方提供的私有仓库 registry，用起来虽然简单 ，但在管理的功能上存在不足。 

Harbor是一个用于存储和分发Docker镜像的企业级Registry服务器，harbor使用的是官方的docker registry(v2命名是distribution)服务去完成。 

harbor在docker distribution的基础上增加了一些安全、访问控制、管理的功能以满足企业对于镜像仓库的需求。