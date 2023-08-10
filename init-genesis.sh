HOME=$(pwd)/home
axelard=$(pwd)/bin/axelard
CHAIN_ID=scalar

# generate key for validator
$axelard keys add massbit-validator-1 --home=$HOME

$axelard init massbit-node-1 --chain-id $CHAIN_ID --home=$HOME

# add genesis account
$axelard add-genesis-account massbit-validator-1 --home=$HOME  100000000stake,1000000uaxl

# add genesis transaction
$axelard gentx massbit-validator-1 100000000stake --chain-id $CHAIN_ID --home=$HOME

# Add the gentx to the genesis file.
$axelard collect-gentxs --home=$HOME