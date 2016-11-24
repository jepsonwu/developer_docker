#!/bin/bash

docker stop memcached%suffix%
docker rm memcached%suffix%
docker rmi %auth%/memcached%suffix%

docker build -t %auth%/memcached%suffix% .
