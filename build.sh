#!/usr/bin/env bash
#support unix,please exec in git bash if you are windows

#project path
project_path='~/projects/docker_demo_projects/'

mkdir -pv ${project_path}

#current path
current_path=`pwd`

#change to base path
cd ${project_path}

#create
mkdir -pv nginx/{conf,logs}
cp -rf ${current_path}."/config/nginx/*" nignx/conf/

mkdir -pv php/{conf,logs}
cp -rf