#!/bin/bash

docker stop nginx
docker rm nginx
docker rmi jepson/nginx

docker build -t jepson/nginx .
