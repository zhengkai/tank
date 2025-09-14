#!/bin/bash -e

if [ ! -d /www/tank ]; then
	exit
fi

mkdir -p /www/tank/output/tmp
mkdir -p /www/tank/output/history
mkdir -p /www/tank/log

sudo docker stop tank || :
sudo docker rm tank || :

# sudo cat /tmp/docker-tank.tar | sudo docker load

sudo docker run -d --name tank \
	--publish "127.0.0.1:21024:80" \
	--mount type=bind,source=/www/tank/static,target=/static \
	--mount type=bind,source=/www/tank/log,target=/log \
	--restart always \
	tank
