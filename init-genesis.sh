#/bin/bash
mkdir ./home
HOME=$(pwd)/home/node
axelard=$(pwd)/bin/axelard
tofnd=$(pwd)/bin/tofnd
CHAIN_ID=scalar
EVM_CONFIG="
[[axelar_bridge_evm]]
name = \"Ethereum\"
rpc_addr = \"http://172.24.97.100:8500/2\"
start-with-bridge = true

[[axelar_bridge_evm]]
name = \"Avalanche\"
rpc_addr = \"http://172.24.97.100:8500/0\"
start-with-bridge = true

[[axelar_bridge_evm]]
name = \"Fantom\"
rpc_addr = \"http://172.24.97.100:8500/1\"
start-with-bridge = true
"
PASSPHRASE="12345678"

echo "Setup genesis file and keys..."

# seeds
for i in 1 2 3 4
do
# init node
$axelard init massbit-node-$i --chain-id $CHAIN_ID --home=$HOME"$i"
# generate key for validator and broadcaster
yes $PASSPHRASE | $axelard keys add massbit-validator-$i --home=$HOME$i
yes $PASSPHRASE | $axelard keys add massbit-broadcaster-$i --home=$HOME$i
yes $PASSPHRASE | $tofnd -m create --directory=$HOME$i
# BROADCASTER_ADDR=$($axelard keys show massbit-broadcaster-$i -a --home=$HOME"$i")
# $axelard tx snapshot register-proxy $BROADCASTER_ADDR --from massbit-validator-$i --chain-id $CHAIN_ID --home=$HOME"$i" --gas auto --gas-adjustment 1.4
if [ $i = 1 ]; then
$axelard set-genesis-mint --blocks-per-year 3153600 --goal-bonded 0.67 --inflation-min 0.05 --inflation-max 0.1 --inflation-max-rate-change 0.05 --mint-denom uaxl --home=$HOME$i
$axelard set-genesis-staking --bond-denom uaxl --max-validators 1000 --unbonding-period "1h" --home=$HOME$i
# $axelard add-genesis-evm-chain "Ethereum Goerli" --home=$HOME$i
# $axelard add-genesis-evm-chain "Avalanche" --home=$HOME$i
# $axelard add-genesis-evm-chain "Fantom" --home=$HOME$i
$axelard set-genesis-crisis --constant-fee=1000uaxl --home=$HOME$i
$axelard set-genesis-gov --minimum-deposit 1000000uaxl --max-deposit-period 172800s --voting-period 172800s --home=$HOME$i
# add genesis account
yes $PASSPHRASE | $axelard add-genesis-account massbit-validator-$i --home=$HOME$i  1000000000uaxl
yes $PASSPHRASE | $axelard add-genesis-account massbit-broadcaster-$i --home=$HOME$i  1000000000uaxl
# add genesis transaction
yes $PASSPHRASE | $axelard gentx massbit-validator-$i 1000000uaxl --chain-id $CHAIN_ID --home=$HOME$i
# # Add the gentx to the genesis file.
$axelard collect-gentxs --home=$HOME$i
else
cp $HOME"$pre/config/genesis.json" $HOME"$i/config/genesis.json"
yes $PASSPHRASE | $axelard add-genesis-account massbit-validator-$i --home=$HOME$i  1000000000uaxl
yes $PASSPHRASE | $axelard add-genesis-account massbit-broadcaster-$i --home=$HOME$i  1000000000uaxl
fi
pre=$i
if [ $i = 4 ]; then
for y in 1 2 3
do
cp $HOME"$i/config/genesis.json" $HOME"$y/config/genesis.json"
done
fi
# setup EVM
echo "$EVM_CONFIG" >> $HOME"$i/config/config.toml"
done

# setup peers
PEER1=$($axelard tendermint show-node-id --home=$HOME"1")@172.24.97.11:26656
PEER2=$($axelard tendermint show-node-id --home=$HOME"2")@172.24.97.12:26656
PEER3=$($axelard tendermint show-node-id --home=$HOME"3")@172.24.97.13:26656
PEER4=$($axelard tendermint show-node-id --home=$HOME"4")@172.24.97.14:26656

sed -i "s/seeds = \"\"/seeds = \"$PEER2,$PEER3,$PEER4\"/g" $HOME"1/config/config.toml"
sed -i "s/seeds = \"\"/seeds = \"$PEER1,$PEER3,$PEER4\"/g" $HOME"2/config/config.toml"
sed -i "s/seeds = \"\"/seeds = \"$PEER1,$PEER2,$PEER4\"/g" $HOME"3/config/config.toml"
sed -i "s/seeds = \"\"/seeds = \"$PEER1,$PEER2,$PEER3\"/g" $HOME"4/config/config.toml"

docker compose up -d --force-recreate
echo "Waiting node and validator ready, please wait..."
sleep "10s"
echo "Init validators and proxcies..."
for i in 1 2 3 4
do
docker exec -it axelar-node-$i /bin/sh /init-validator.sh
done

docker compose down
echo "==> Finished init data, run \`docker compose up\` to start nodes"
