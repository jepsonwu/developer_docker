# Docker tool

## Summary
This is a docker build tool for developers.

## System
Base on alpine.

## What you'll need
- Install docker

## All of containers here
- alpine
- store
- memcached
- reids
- rebitmq
- ssdb
- mysql
- mongodb
- php
- nginx

## Build from sh

```
./build.sh
```

## Manual build

```
cd alpine && ./build.sh
```

```
docker run --name store -d -it -v  /data_path:/data jepson/store

docker run --name memcached -d jepson/memcached

docker run --name php -d -it -v  /data_path:/app --volumes-from store --link memcached jepson/php

docker run --name nginx -d -p 80:80 -v  /data_path:/app --volumes-from store --link php jepson/nginx
```

## More
To learn more about all the features of [docker compose](https://docs.docker.com/compose/compose-file/) see the list of features.

## Notice

### docker
- 环境记得配置docker允许共享目录

### docker-compose.yml

- version: '2' 冒号之后需要有空格
-  alpine: 选项之前只能是空格
