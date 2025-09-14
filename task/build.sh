#!/bin/bash -ex

curl -s --max-time 10 --noproxy --ipv4 http://localhost:21024/task/build
