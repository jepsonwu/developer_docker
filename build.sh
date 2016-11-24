#!/usr/bin/env bash
#支持unix windows下请在git bash下执行

#project path
project_path='/c/Users/wjp13/projects/docker_demo_projects/'  #windows
#project_path='/data/' #unix

#docker name
docker_name='jepson'

mkdir -pv ${project_path}

#current path
current_path=`pwd`

#change to base path
cd ${project_path}


#create
mkdir -pv nginx/{conf,logs}
cp -rf ${current_path}"/projects/config/nginx/"* nginx/conf/

mkdir -pv php/{conf,logs}
cp -rf ${current_path}"/projects/config/php/"* php/conf/

mkdir -pv data

cp -rf ${current_path}"/projects/data/"* data/

#build
cd ${current_path}

#docker-compose.exe up #windows

#docker-compose up #unix

