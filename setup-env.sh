make prereqs
# build node binary and image
make build
make docker-image

# build tofnd binary and image
cd ../tofnd
docker build -t axelar/tofnd .
sh -x ./scripts/copy-binaries-from-image.sh
cd ../axelar-core
cp ../tofnd/bin/tofnd ./bin/tofnd

# build evm local image
cd ../axelar-examples
docker build -t evm-local .

# create docker network
docker network create -d bridge --gateway "172.24.97.1" --subnet "172.24.97.0/24" massbit