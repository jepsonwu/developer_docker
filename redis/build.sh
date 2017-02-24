#!/bin/bash

docker stop redis%suffix%
docker rm redis%suffix%
docker rmi %auth%/redis%suffix%

docker build -t %auth%/redis%suffix% .
