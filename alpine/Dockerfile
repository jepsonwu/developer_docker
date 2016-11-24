#base image on alpine
FROM alpine:latest

#maintainer
MAINTAINER %auth% %email%

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

#exec command when start container
CMD ["sh"]