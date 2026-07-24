#!/usr/bin/bash
set -e
APP_PATH=/repo/app

#image run
mkdir -p "${APP_PATH}/kafka"
cp docker-compose.yaml "${APP_PATH}/kafka"
cd "${APP_PATH}/kafka"
docker compose -f docker-compose.yaml up -d
if [ $? -eq 0 ]; then
    echo "\nstart kafka succeeded\n"
else
    echo "\nstart kafka failed\n"
fi
cd -
