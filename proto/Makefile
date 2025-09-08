SHELL:=/bin/bash

server:
	protoc --go_out=../server/src *.proto

client:
	../client/node_modules/protobufjs-cli/bin/pbjs -t static-module -w es6 --es6 --force-number -o ../client/src/pb/pb.js *.proto
	../client/node_modules/protobufjs-cli//bin/pbts -o ../client/src/pb/pb.d.ts ../client/src/pb/pb.js
	sed 's#^import.*$$#const $$protobuf = require("protobufjs/minimal");#g' -i ../client/src/pb/pb.js
	#tar zcvf ../client/src/assets/proto.tar.gz *.proto
