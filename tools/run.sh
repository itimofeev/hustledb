#!/usr/bin/env bash

docker load -i hustlesa.img
docker load -i nginxhustlesa.img

tar -jxvf frontend.tar.bz2

docker-compose -p hustlesa -f prod.docker-compose.yml up -d
