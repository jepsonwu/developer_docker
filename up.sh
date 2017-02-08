#!/usr/bin/env bash
#support unix,windows system,please execute under git bash if you on windows

#=========================you need config=================================#
#project path,like "/c/Users/wjp13/projects/docker_demo_projects/"
project_path=''
if [ "${project_path}X" = "X" ];then
    echo "You must be config project path"
    exit 0
fi

#docker auth jepson wjp13671142513@163.com
auth=''
email=''
if [ "${auth}X" = "X" ] || [ "${email}X" = "X" ];then
    echo "You must be config auth and email"
    exit 0
fi

#container suffix,like _1
suffix='' #empty

#system
read -p "what's the current operating system?unix:1,windows:2" type
case ${type} in
2) os="windows"
;;
*) os="unix"
;;
esac

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
lock_file="build_lock_file"

if [ ! -f ${lock_file} ];then
    sed -irn "s/%auth%/${auth}/g" `grep %auth% -rl ./ |grep -v up`
    sed -irn "s/%email%/${email}/g" `grep %email% -rl ./ |grep -v up`
    sed -irn "s/%suffix%/${suffix}/g" `grep %suffix% -rl ./ |grep -v up`

    sed -in "s#%store_data%#${project_path}data#" docker-compose.yml
    sed -in "s#%php_app%#${project_path}php#" docker-compose.yml
    sed -in "s#%nginx_app%#${project_path}nginx#" docker-compose.yml

    rm -rf `grep -rl ./ |grep rn`
    touch ${lock_file}
fi

if [ "${os}X" = 'windowsX' ];then
    docker-compose.exe up
else
    docker-compose up
fi

