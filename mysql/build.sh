#!/bin/bash

docker stop mysql%suffix%
docker rm mysql%suffix%
docker rmi %auth%/mysql%suffix%

docker build -t %auth%/mysql%suffix% .
