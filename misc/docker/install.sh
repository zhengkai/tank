#!/bin/bash -e

TARGET="Eirena"

if [ "$HOSTNAME" != "$TARGET" ]; then
	>&2 echo "only run in server $TARGET"
	exit
fi

mkdir -p /www/tank/output
mkdir -p /www/tank/log
mkdir -p /www/tank/tmp

sudo docker stop tank || :
sudo docker rm tank || :
sudo docker rmi tank || :

sudo cat /tmp/docker-tank.tar | sudo docker load

sudo docker run -d --name tank \
	--env "TANK_MYSQL=tank:tank@tcp(172.17.0.1:3306)/tank" \
	--env "TMP_PATH=/tmp" \
	--env "OUTPUT_PATH=/output" \
	--publish "127.0.0.1:21024:80" \
	--mount type=bind,source=/www/tank/output,target=/output \
	--mount type=bind,source=/www/tank/log,target=/log \
	--mount type=bind,source=/www/tank/tmp,target=/tmp \
	--restart always \
	tank
