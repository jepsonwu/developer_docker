#!/usr/bin/env bash
#支持unix windows下请在git bash下执行

#=========================you need config=================================#
#project path
project_path='/c/Users/wjp13/projects/docker_demo_projects/'

#docker auth
auth='jepson'
email='wjp13671142513@163.com'

#container suffix
suffix='_1' #empty

#system
os='windows' #unix

#===========================start build=================================#
#start
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

sed -irn "s/%auth%/${auth}/g" `grep %auth% -rl ./ |grep -v up`
sed -irn "s/%email%/${email}/g" `grep %email% -rl ./ |grep -v up`
sed -irn "s/%suffix%/${suffix}/g" `grep %suffix% -rl ./ |grep -v up`

sed -in "s#%store_data%#${project_path}data#" docker-compose.yml
sed -in "s#%php_app%#${project_path}php#" docker-compose.yml
sed -in "s#%nginx_app%#${project_path}nginx#" docker-compose.yml

rm -rf `grep -rl ./ |grep rn`

if [ "${os}"x = 'windows'x ];then
    docker-compose.exe up
else
    docker-compose up
fi

