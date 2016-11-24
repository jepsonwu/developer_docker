#!/bin/bash

docker stop store%suffix%
docker rm store%suffix%
docker rmi %auth%/store%suffix%

docker build -t %auth%/store%suffix% .
