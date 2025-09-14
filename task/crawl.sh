#!/bin/bash -ex

curl localhost:21024/task/crawl

echo
cd /www/tank/static/data || exit 1

find . -type f \( -name '*.gz' -o -name '*.br' \) -delete
find . -type f \( -name '*.json' -o -name '*.pb' \) -size +1k -exec sh -c '
  for f; do
    gzip -c "$f" > "$f.gz"
    brotli -q 11 -o "$f.br" "$f"
  done
' sh {} +

rsync --partial -vzrtopg ./ doll:/www/tank/data
