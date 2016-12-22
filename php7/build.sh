#!/bin/bash

docker stop php7%suffix%
docker rm php7%suffix%
docker rmi %auth%/php7%suffix%

docker build -t %auth%/php7%suffix% .
