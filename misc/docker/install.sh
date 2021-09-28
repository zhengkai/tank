#!/bin/bash

if [ "$HOSTNAME" != "Gnar" ]; then
	>&2 echo only run in server Gnar
	# exit 1
fi

sudo docker stop tank
sudo docker rm tank
sudo docker rmi tank

sudo cat /tmp/docker-tank.tar | sudo docker load
# sudo cat /www/gnar/docker-demo-image.tar | sudo docker load

sudo docker run -d --name tank \
	--env "TANK_MYSQL=tank:tank@tcp(172.17.0.1:3306)/tank" \
	--env "TMP_PATH=/tmp" \
	--env "OUTPUT_PATH=/output" \
	--mount type=bind,source=/www/tank/output,target=/output \
	--mount type=bind,source=/www/tank/log,target=/log \
	--mount type=bind,source=/www/tank/tmp,target=/tmp \
	--restart always \
	tank
