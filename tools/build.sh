#!/usr/bin/env bash

export GOPATH=/Users/ilyatimofee/prog/axxonsoft/

PROJECT_PATH=/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustledb
FRONTEND_PROJECT_PATH=/Users/ilyatimofee/prog/js/hustlesa-ui

rm ${PROJECT_PATH}/target/*
mkdir ${PROJECT_PATH}/target


docker build --force-rm=true -t nginxhustledb -f ${PROJECT_PATH}/tools/nginx.Dockerfile .
docker save -o "$PROJECT_PATH/target/nginxhustledb.img" "nginxhustledb"

export GOOS=linux
export GOARCH=amd64
go build -v github.com/itimofeev/hustledb/main/hustledb



docker build --force-rm=true -t hustledb -f ${PROJECT_PATH}/tools/hustledb.Dockerfile .
docker save -o "$PROJECT_PATH/target/hustledb.img" "hustledb"


rm hustledb

cp ${PROJECT_PATH}/tools/prod.docker-compose.yml ${PROJECT_PATH}/tools/run.sh ${PROJECT_PATH}/tools/prod.env ${PROJECT_PATH}/tools/postgres.env ${PROJECT_PATH}/target/

echo 'building frontend'

#npm build ${FRONTEND_PROJECT_PATH}
cp -r ${FRONTEND_PROJECT_PATH}/build ${PROJECT_PATH}/target/frontend
cd target
tar -jcvf ${PROJECT_PATH}/target/frontend.tar.bz2 frontend

rm -r frontend
