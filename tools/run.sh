#!/usr/bin/env bash

docker load -i hustlesa.img
docker load -i nginxhustlesa.img

docker-compose -p hustlesa up -d
