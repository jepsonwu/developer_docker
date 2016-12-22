#base image on alpine
FROM %auth%/alpine%suffix%:latest

#maintainer
MAINTAINER %auth% %email%

#install base
RUN apk add rabbitmq-c

#install php
RUN apk add php7 \
php7-zip \
#php7-cgi \
#php7-embed \
php7-xmlrpc \
php7-sysvshm \
php7-mysql \
#php7-imap \
php7-amqp \
php7-doc \
php7-imagick \
php7-zlib \
php7-calendar \
php7-mysqli \
php7-soap \
php7-shmop \
#php7-wddx \
php7-cli \
php7-fpm \
#php7-phpdbg \
#php7-embed \
php7-bz2 \
php7-sockets \
#php7-phalcon \
#php7-memcache \
#php7-apache2 \
php7-pdo_mysql \
#php7-mailparse \
php7-xdebug \
php7-sysvmsg \
#php7-pspell \
php7-iconv \
## pecl required
php7-dev \
php7-ftp \
php7-gettext \
#php7-mssql \
php7-mcrypt \
#php7-exif \
php7-xmlreader \
#php7-xcache \
php7-gd \
php7-xml \
php7-pcntl \
php7-phpmailer \
php7-phar \
php7-apcu \
php7-ctype \
#php7-intl \
php7-pdo \
php7-openssl \
php7-memcached \
php7-common \
php7-sysvsem \
php7-posix \
php7-dom \
php7-redis \
php7-curl \
#php7-xsl \
#php7-ldap \
php7-json \
#php7-enchant \
php7-bcmath \
php7-pear
#php7-opcache
#php7-sqlite3 \
#php7-gmp \
#php7-snmp

#config php
RUN mkdir -pv /app/{conf,logs}
RUN mkdir -pv /data
RUN touch /etc/php7/fpm.d/none.conf
RUN touch /var/log/php_errors.log
RUN chmod 777 -R /var/log/
#去掉ip绑定限制
RUN sed -i -e "s/listen\s*=\s*127.0.0.1:9000/listen = 9000/g" /etc/php7/php-fpm.conf
RUN mv /etc/php7/php.ini /etc/php7/php.ini_bak
RUN ln -s /app/conf/php.ini /etc/php7/php.ini
RUN cd /data/web/
RUN curl -sS https://getcomposer.org/installer |php7

##install swoole
#RUN apk add linux-headers ##install from source need head
#RUN pecl install swoole ##1.9
##you should install from source if need 2.0
##add ini extension = swoole.so

##install libevent  ## pecl 去掉-n 否则报错
#RUN apk add autoconf gcc g++ libevent libevent-dev make
#RUN pecl install libevent-0.1.0
#RUN chmod +x /usr/lib/php7/modules/libevent.so
## add ini extension = sockets.so  extension=libevent.so

#expose
EXPOSE 9000

#volume
VOLUME /app

#
#ENTRYPOINT php-fpm

#exec command when start container
CMD ["php-fpm","-F"]
