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
- php5.6
- php7
- nginx

## Build from exec
```
sudo ./up -a auth -p project_path
```

```
up.exe -a auth -p project_path
```

## Build from sh

```
./up.sh
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
- You should shared local drivers you want to be available to your containers

### docker-compose.yml

- {version: '2'} There is a space after the colon
- { alpine:} Only is a space before option
