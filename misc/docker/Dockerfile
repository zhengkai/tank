# builder stage

FROM golang:latest as builder

ARG DOCKER_RUNNING=yes
ARG DEBIAN_FRONTEND=noninteractive
RUN apt update && apt install -yq protobuf-compiler 
RUN go get google.golang.org/protobuf/cmd/protoc-gen-go

COPY . /project

RUN cd /project/proto && make

RUN /project/server/build/build-server.sh prod

RUN go version

# clean stage

FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/

RUN apk add --no-cache gzip brotli

COPY --from=builder /project/server/dist/prod/tank-server-next /tank-server

CMD ["/tank-server"]
