#!/bin/bash

docker stop store
docker rm store
docker rmi jepson/store

docker build -t jepson/store .
