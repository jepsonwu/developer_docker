#!/bin/bash
docker stop alpine
docker rm alpine
docker rmi jepson/alpine

docker build -t jepson/alpine .