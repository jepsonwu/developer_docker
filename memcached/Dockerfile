#base image on alpine
FROM jepson/alpine:latest

#maintainer
MAINTAINER jepson wjp13671142513@163.com

#apk install 
RUN apk add memcached

#config

#expose
EXPOSE 11211

#volume

#exec user
USER daemon

#
ENTRYPOINT memcached

#exec command when start container
CMD ["-m","128"]
