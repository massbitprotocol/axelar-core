make build
make docker-image
docker network create -d bridge --gateway "172.24.97.1" --subnet "172.24.97.0/24" massbit