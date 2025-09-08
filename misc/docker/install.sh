#!/bin/bash -e

mkdir -p /www/tank/output/tmp
mkdir -p /www/tank/log

sudo docker stop tank || :
sudo docker rm tank || :
sudo docker rmi tank || :

# sudo cat /tmp/docker-tank.tar | sudo docker load

sudo docker run -d --name tank \
	--env "TMP_PATH=/output/tmp" \
	--env "OUTPUT_PATH=/output" \
	--publish "127.0.0.1:21024:80" \
	--mount type=bind,source=/www/tank/output,target=/output \
	--mount type=bind,source=/www/tank/log,target=/log \
	--restart always \
	zhengkai/tank
