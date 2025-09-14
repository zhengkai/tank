# builder stage

FROM golang:latest AS builder

ARG DOCKER_RUNNING=yes
ARG DEBIAN_FRONTEND=noninteractive
RUN apt update && apt install -yq protobuf-compiler tzdata ca-certificates
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.8

WORKDIR /project

COPY server/src/go.mod server/src/go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	go mod graph && go mod download

COPY . .

RUN cd /project/proto && make

RUN --mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	/project/server/build/build-server.sh prod

# clean stage

FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/
COPY --from=builder /project/server/dist/prod/tank-server-next /tank-server

ENV TZ="Asia/Shanghai"

ENV TANK_MYSQL="tank:tank@tcp(172.17.0.1:3306)/tank"

EXPOSE 80

CMD ["/tank-server"]
