#base image on alpine
FROM %auth%/alpine%suffix%:latest

#maintainer
MAINTAINER %auth% %email%

RUN apk update && apk upgrade

#apk install 
RUN apk add redis

#config

#expose
EXPOSE 6379

#volume

#
ENTRYPOINT redis-server