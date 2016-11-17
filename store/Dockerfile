#base image on alpine
FROM jepson/alpine_1:latest

#arg
ARG auth
ARG email

#maintainer
MAINTAINER $auth $email

#store
RUN mkdir -pv /data

#volume
VOLUME /data

#command
ENTRYPOINT /bin/sh


