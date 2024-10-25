## Envoy sidecar proxying gRPC with tls and header based routing

We use envoy to proxy gRPC requests. where client points at a localhost port and send rpc with header. envoy client sidecar will route to different upstream based on header.

`client --> envoy(9999) -tls-> envoy(4443) -> server:50051`

1. start docker
```
docker compose pull
docker compose up --build -d
docker compose ps
```



3. send gRPC requests
find client container id:
```
docker container ls | grep client
```
shell into container:
```
docker exec -it <id> /bin/bash

root@id:/app# go run src/grpc_client.go --hostname client-egress:9999 --originalHost server-2
2024/10/25 05:49:40 hostname: client-egress:9999, originalHost: server-2
2024/10/25 05:49:40 Echo response: what's up


root@id:/app# go run src/grpc_client.go --hostname client-egress:9999 --originalHost server-1
2024/10/25 05:49:42 hostname: client-egress:9999, originalHost: server-1
2024/10/25 05:49:42 Echo response: what's up

```
