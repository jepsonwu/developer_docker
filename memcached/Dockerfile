#base image on alpine
FROM jepson/alpine_1:latest

#arg
ARG auth
ARG email

#maintainer
MAINTAINER $auth $email

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
