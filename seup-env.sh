make prereqs
make build
make docker-image
docker network create -d bridge --gateway "172.24.97.1" --subnet "172.24.97.0/24" massbit
cd ../tofnd
docker build -t axelar/tofnd .
sh -x ./scripts/copy-binaries-from-image.sh
cd ../axelar-core
cp ../tofnd/bin/tofnd ./bin/tofnd