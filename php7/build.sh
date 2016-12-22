#!/bin/bash

docker stop php%suffix%
docker rm php%suffix%
docker rmi %auth%/php%suffix%

docker build -t %auth%/php%suffix% .
