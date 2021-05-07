FROM golang:1.16

# Configure Go 生产模式不需要安装go
ENV GOPROXY http://goproxy.cn/
ENV GO111MODULE on
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

RUN go get -u github.com/cosmtrek/air
RUN go get -u gorm.io/gorm
RUN go get -u github.com/gin-gonic/gin

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

WORKDIR /root/app
# ENTRYPOINT ["/wait-for-it.sh"]
