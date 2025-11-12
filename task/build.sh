#!/bin/bash -ex

unset HTTP_PROXY
unset HTTPS_PROXY

curl -s --max-time 10 --noproxy "*" --ipv4 http://localhost:21024/task/build

sleep 100

echo
cd /www/tank/static/data || exit 1

sudo find . -type f \( -name '*.gz' -o -name '*.br' \) -delete
sudo find . -type f \( -name '*.json' -o -name '*.pb' \) -size +1k -exec sh -c '
  for f; do
    gzip -c "$f" > "$f.gz"
    brotli -q 11 -o "$f.br" "$f"
  done
' sh {} +

rsync --partial -vzrtopg ./ doll:/www/tank/data
