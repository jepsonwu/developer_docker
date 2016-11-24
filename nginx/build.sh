#!/bin/bash

docker stop nginx%suffix%
docker rm nginx%suffix%
docker rmi %auth%/nginx%suffix%

docker build -t %auth%/nginx%suffix% .
