# Query worker

Worker for uploading meow messages to elastic

## Run nats docker

`docker network create nats`
`docker run -d --name nats --network nats --rm -p 4222:4222 -p 8222:8222 nats`

## Run elasticsearch docker

`docker network create elasticsearch`
`docker run -d --name elasticsearch --network elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.14.2`

## Build docker

`docker build --tag cyber-meower-query-worker .`
