SHELL:=/bin/bash

build: git
	echo "tank 的 docker 网络有问题，暂时屏蔽"
	exit 1
	sudo docker build -t tank -f Dockerfile ../..

run: build
	sudo docker run --env "TANK_MYSQL=tank:tank@tcp(172.17.0.1:3306)/tank" \
		--mount type=bind,source=/www/tank/output,target=/output \
		--mount type=bind,source=/www/tank/tmp,target=/tmp \
		--mount type=bind,source=/www/tank/server/dist/prod/log,target=/log \
		tank

install: build
	sudo docker save tank > docker-tank.tar
	scp docker-tank.tar freya:/tmp
	scp install.sh freya:/tmp
	ssh freya "chmod +x /tmp/install.sh && /tmp/install.sh && rm /tmp/install.sh"

git:
	../../server/build/git-hash.sh > ../../server/build/.git-hash

clean:
	sudo docker stop tank || :
	sudo docker rm tank || :
	sudo docker rmi tank || :
	sudo docker image prune --force
