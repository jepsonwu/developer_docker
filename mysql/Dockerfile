#base image on alpine
FROM %auth%/alpine%suffix%:latest

#maintainer
MAINTAINER %auth% %email%

#must exec before volume
RUN mkdir -pv /data
#RUN mkdir -pv /data/mysql

WORKDIR /data

VOLUME /data

COPY startup.sh /startup.sh
RUN chmod +x /startup.sh

RUN apk add --update mysql mysql-client && rm -f /var/cache/apk/*

COPY my.cnf /etc/mysql/my.cnf

EXPOSE 3306

CMD ["/startup.sh"]
