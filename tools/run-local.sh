#!/usr/bin/env bash

TOOLS_PATH=/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustledb/tools
TARGET_PATH=/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustledb/target

${TOOLS_PATH}/build.sh

docker load -i ${TARGET_PATH}/hustledb.img
docker load -i ${TARGET_PATH}/nginxhustledb.img
docker-compose -p hustledb -f ${TOOLS_PATH}/local.docker-compose.yml up -d