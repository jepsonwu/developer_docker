#base image on alpine
FROM %auth%/alpine%suffix%:latest

#maintainer
MAINTAINER %auth% %email%

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
