#!/bin/bash
docker stop alpine%suffix%
docker rm alpine%suffix%
docker rmi %auth%/alpine%suffix%

docker build -t %auth%/alpine%suffix% .