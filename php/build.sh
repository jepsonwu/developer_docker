#!/bin/bash

docker stop php
docker rm php
docker rmi jepson/php

docker build -t jepson/php .
