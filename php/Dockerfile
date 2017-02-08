#base image on alpine
FROM %auth%/alpine%suffix%:latest

#maintainer
MAINTAINER %auth% %email%

RUN apk update && apk upgrade

#install base
RUN apk add rabbitmq-c

#install php
RUN apk add php5 \
php5-zip \
#php5-cgi \
#php5-embed \
php5-xmlrpc \
php5-sysvshm \
php5-mysql \
#php5-imap \
php5-amqp \
php5-doc \
php5-imagick \
php5-zlib \
php5-calendar \
php5-mysqli \
php5-soap \
php5-shmop \
#php5-wddx \
php5-cli \
php5-fpm \
#php5-phpdbg \
#php5-embed \
php5-bz2 \
php5-sockets \
#php5-phalcon \
php5-memcache \
#php5-apache2 \
php5-pdo_mysql \
#php5-mailparse \
php5-xdebug \
php5-sysvmsg \
#php5-pspell \
php5-iconv \
## pecl required
php5-dev \
php5-ftp \
php5-gettext \
#php5-mssql \
php5-mcrypt \
#php5-exif \
php5-xmlreader \
#php5-xcache \
php5-gd \
php5-xml \
php5-pcntl \
php5-phpmailer \
php5-phar \
php5-apcu \
php5-ctype \
#php5-intl \
php5-pdo \
php5-openssl \
#php5-memcached \
php5-common \
php5-sysvsem \
php5-posix \
php5-dom \
php5-redis \
php5-curl \
#php5-xsl \
#php5-ldap \
php5-json \
#php5-enchant \
php5-bcmath \
php5-pear
#php5-opcache
#php5-sqlite3 \
#php5-gmp \
#php5-snmp

#config php
RUN mkdir -pv /app/conf && mkdir -pv /app/logs
RUN mkdir -pv /data
RUN touch /etc/php5/fpm.d/none.conf
RUN touch /var/log/php_errors.log
RUN chmod 777 -R /var/log/
#去掉ip绑定限制
RUN sed -i -e "s/listen\s*=\s*127.0.0.1:9000/listen = 9000/g" /etc/php5/php-fpm.conf
RUN mv /etc/php5/php.ini /etc/php5/php.ini_bak
RUN ln -s /app/conf/php.ini /etc/php5/php.ini
#RUN cd /data/web/
#RUN curl -sS https://getcomposer.org/installer |php5

##install swoole
#RUN apk add linux-headers ##install from source need head
#RUN pecl install swoole ##1.9
##you should install from source if need 2.0
##add ini extension = swoole.so

##install libevent  ## pecl 去掉-n 否则报错
#RUN apk add autoconf gcc g++ libevent libevent-dev make
#RUN pecl install libevent-0.1.0
#RUN chmod +x /usr/lib/php5/modules/libevent.so
## add ini extension = sockets.so  extension=libevent.so

##install sshd
RUN apk add openssh
RUN ssh-keygen -t rsa -b 2048 -f /etc/ssh/ssh_host_rsa_key
RUN sed -i -e "s/#HostKey \/etc\/ssh\/ssh_host_rsa_key/HostKey \/etc\/ssh\/ssh_host_rsa_key/g" /etc/ssh/sshd_config
#RUN ssh-keygen -t rsa
#RUN touch ~/.ssh/authorized_keys

#expose
EXPOSE 9000

#volume
VOLUME /app

#
#ENTRYPOINT php-fpm

#exec command when start container
CMD ["php-fpm","-F"]
