#base image on alpine
FROM %auth%/alpine%suffix%:latest

#maintainer
MAINTAINER %auth% %email%

RUN apk update && apk upgrade

#store
RUN mkdir -pv /data

#volume
VOLUME /data

#command
ENTRYPOINT /bin/sh


