# co-writer-gin-cgi
  
    my project using go gin framework.
    co writer gin cgi project.

- 1、安装 gin:

go get gopkg.in/gin-gonic/gin.v1

- 2、发现go get服务被墙, 配置lantern:



### 配置蓝灯代理：

查看进程监听端口：

lsof -Pn | grep lantern | grep LISTEN

在 ~/.bash_profile中添加:

    export http_proxy=http://127.0.0.1:55945
    export https_proxy=$http_proxy
    export ftp_proxy=$http_proxy
    export rsync_proxy=$http_proxy
    export no_proxy="localhost,127.0.0.1,.baidu.com"

source ~/.bash_profile

### 配置Shadowsocks代理：

#### 安装 polipo：

参考:

http://qichunren.github.io/tool/2014/07/15/Convert-shadowsocks-into-http-proxy-on-mac/

brew install polipo

#### 配置polipo代理：

    <array>
        <string>/usr/local/opt/polipo/bin/polipo</string>
        <string>socksParentProxy=localhost:1080</string>
      </array>

#### 启动服务：

brew services start polipo

#### 配置软链：
ln -sfv /usr/local/opt/polipo/*.plist ~/Library/LaunchAgents launchctl load ~/Library/LaunchAgents/homebrew.mxcl.polipo.plist

#### 在 ~/.bash_profile中添加:

    export https_proxy=http://localhost:8123
    export ftp_proxy=http://localhost:8123
    export no_proxy="localhost,127.0.0.1,.baidu.com"

source ~/.bash_profile


#### 重新安装

go get gopkg.in/gin-gonic/gin.v1


## 使用gin-boilerplate作为系统模板


    use gin-boilerplate for project template.
    github: https://github.com/Massad/gin-boilerplate

### 构建系统环境:

- 安装 postgreSQL

- 安装redis

- 进入psql环境

sudo -u kakachan psql postgres

press \q to exit psql command line environment.

- 创建 postgres用户

CREATE USER postgres SUPERUSER;

- 创建postgres数据库

CREATE DATABASE postgres WITH OWNER postgres;

- 导入DB模板:

psql -U postgres -h localhost < ./db/database.sql

- 运行服务:

go run main.go

- 访问浏览器:

http://localhost:9000

### 配置mongodb:

- 安装mongo:

go get gopkg.in/mgo.v2

- 参考如下demo project 构建mongodb mgo 连接

https://github.com/madhums/go-gin-mgo-demo


