mkdir ./home
INDEX=1
HOME=$(pwd)/home/node
axelard=$(pwd)/bin/axelard
CHAIN_ID=scalar

$axelard init massbit-node-1 --chain-id $CHAIN_ID --home=$HOME"1"
# generate key for validator
$axelard keys add massbit-validator-1 --home=$HOME"1"
# add genesis account
$axelard add-genesis-account massbit-validator-1 --home=$HOME"1"  100000000stake,1000000uaxl

# add genesis transaction
$axelard gentx massbit-validator-1 100000000stake --chain-id $CHAIN_ID --home=$HOME"1"

# Add the gentx to the genesis file.
$axelard collect-gentxs --home=$HOME"1"

# seeds
for i in 2 3 4
do
$axelard init massbit-node-i --chain-id $CHAIN_ID --home=$HOME"$i"
cp $HOME"1/config/genesis.json" $HOME"$i/config/genesis.json"
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