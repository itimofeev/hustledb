#!/usr/bin/env bash

docker load -i hustledb.img
docker load -i nginxhustledb.img

tar -jxvf frontend.tar.bz2

docker-compose -p hustledb -f prod.docker-compose.yml up -d
