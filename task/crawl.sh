#!/bin/bash

curl localhost:21024/task/crawl

echo
rsync --partial -vzrtopg --exclude="tmp/**" /www/tank/output doll:/www/tank
