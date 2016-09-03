#!/usr/bin/env bash

export GOPATH=/Users/ilyatimofee/prog/axxonsoft/

PROJECT_PATH=/Users/ilyatimofee/prog/axxonsoft/src/github.com/itimofeev/hustlesa

rm $PROJECT_PATH/target/*


docker build --force-rm=true -t nginxhustlesa -f $PROJECT_PATH/tools/nginx.Dockerfile .
docker save -o "$PROJECT_PATH/target/nginxhustlesa.img" "nginxhustlesa"

export GOOS=linux
export GOARCH=amd64
go build -v github.com/itimofeev/hustlesa



docker build --force-rm=true -t hustlesa -f $PROJECT_PATH/tools/hustlesa.Dockerfile .
docker save -o "$PROJECT_PATH/target/hustlesa.img" "hustlesa"


rm hustlesa

cp tools/docker-compose.yml tools/run.sh tools/prod.env tools/postgres.env $PROJECT_PATH/target/
