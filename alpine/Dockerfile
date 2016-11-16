#base image on alpine
FROM alpine:latest

#maintainer
MAINTAINER jepson wjp13671142513@163.com

#china mirrors
RUN echo "http://mirrors.ustc.edu.cn/alpine/edge/main" >> /etc/apk/repositories
RUN echo "http://mirrors.ustc.edu.cn/alpine/edge/community" >> /etc/apk/repositories
RUN echo "http://mirrors.ustc.edu.cn/alpine/edge/testing" >> /etc/apk/repositories
RUN apk update && apk upgrade

#timezone
ENV TIMEZONE Asia/Shanghai
RUN apk add tzdata
RUN ln -snf /usr/share/zoneinfo/$TIMEZONE /etc/localtime
RUN echo $TIMEZONE > /etc/timezone

#install base packages
RUN apk add vim
#tar \
#wget \


#exec command when start container
#这里不能用 ENTRYPOINT 后面继承改镜像的 cmd会当作参数传入 无法执行
CMD ["sh"]
