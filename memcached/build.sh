#!/bin/bash

docker stop memcached
docker rm memcached
docker rmi jepson/memcached

docker build -t jepson/memcached .
