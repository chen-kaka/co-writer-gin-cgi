# go-gin

    my project using go gin framework.

- 1、安装 gin:

go get gopkg.in/gin-gonic/gin.v1

- 2、发现go get服务被墙, 配置lantern:

查看进程监听端口：

lsof -Pn | grep lantern | grep LISTEN

配置蓝灯代理：

在 ~/.bash_profile中添加:

    export http_proxy=http://127.0.0.1:55945

    export https_proxy=$http_proxy

    export ftp_proxy=$http_proxy

    export rsync_proxy=$http_proxy

    export no_proxy="localhost,127.0.0.1,.dade.com"

source ~/.bash_profile

重新安装 go get gopkg.in/gin-gonic/gin.v1


