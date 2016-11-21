#base image on alpine
FROM jepson/alpine_1:latest

#arg
ARG auth
ARG email

#maintainer
MAINTAINER $auth $email

#install nginx version
RUN mkdir -pv /app/{conf,logs}
RUN mkdir -pv /data

#apk install nginx
RUN apk add nginx

#config nginx
RUN mv /etc/nginx/nginx.conf /etc/nginx/nginx.conf_bak
RUN ln -s /app/conf/nginx.conf /etc/nginx/nginx.conf

#expose
EXPOSE 80

#volume
VOLUME /app

#exec command when start container
CMD ["nginx", "-g", "daemon off;"]
