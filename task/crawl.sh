#!/bin/bash -ex

unset HTTP_PROXY
unset HTTPS_PROXY

curl -s --max-time 10 --noproxy "*" --ipv4 localhost:21024/task/crawl
