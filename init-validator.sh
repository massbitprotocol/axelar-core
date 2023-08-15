#/bin/bash
VALIDATOR_ADDRESS=$(yes $KEYRING_PASSWORD | axelard keys show $VALIDATOR_NAME -a --bech val)
BROADCASTER_ADDR=$(yes $KEYRING_PASSWORD | axelard keys show $BROADCASTER_NAME -a)

if axelard q snapshot proxy $VALIDATOR_ADDRESS; then
echo "Proxy exists!"
else
echo "Register proxy..."
yes $KEYRING_PASSWORD | axelard tx snapshot register-proxy $BROADCASTER_ADDR --from $VALIDATOR_NAME --chain-id $CHAIN_ID
fi


if axelard query staking validators | grep -q $VALIDATOR_ADDRESS; then
echo "Validator exists!"
else
echo "==> Creating validator..."
yes $KEYRING_PASSWORD | axelard tx staking create-validator \
  --amount=10000uaxl \
  --pubkey="$(axelard tendermint show-validator)" \
  --moniker=$VALIDATOR_NAME \
  --chain-id=scalar \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1000" \
  --from=$VALIDATOR_NAME
fi

echo "Delegate to validator..."
yes $KEYRING_PASSWORD | axelard tx staking delegate $VALIDATOR_ADDRESS 1000uaxl --from $VALIDATOR_NAME

# get staking amount
axelard q staking validator $VALIDATOR_ADDRESS