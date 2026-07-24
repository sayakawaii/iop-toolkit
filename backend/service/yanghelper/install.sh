#!/usr/bin/bash
set -e
APP_PATH=/repo/app

#pull repo
rm -rf yanghelper
git clone git@gitlabe1.ext.net.nokia.com:iop_toolkit/yanghelper.git
if [ $? -eq 0 ]; then
    echo "\nclone yanghelper succeeded\n"
else
    echo "\nclone yanghelper failed\n"
fi

#yanghelper image build
cd yanghelper
chmod +x run.sh
docker build --build-arg http_proxy=http://10.158.100.1:8080/ --build-arg https_proxy=http://10.158.100.1:8080/ -t yanghelper .
if [ $? -eq 0 ]; then
    echo "\nbuild yanghelper succeeded\n"
else
    echo "\nbuild yanghelper failed\n"
fi
cd -

#image run
mkdir -p "${APP_PATH}/yanghelper"
cp docker-compose-yanghelper.yaml "${APP_PATH}/yanghelper"
cd "${APP_PATH}/yanghelper"
docker compose -f docker-compose-yanghelper.yaml up -d
if [ $? -eq 0 ]; then
    echo "\nstart yanghelper succeeded\n"
else
    echo "\nstart yanghelper failed\n"
fi
cd -
