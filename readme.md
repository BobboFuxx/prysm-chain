# Prysm Validator Setup Guide

## Prerequisites
**Install Go:**
Visit the [official Go installation page](https://golang.org/dl/) to download and install Go. Verify your installation by running `go version` in the terminal.

## Step 1: Clone Prysm Chain Repository
```
git clone https://github.com/BobboFuxx/prysm-chain/
```
## Step 2: Navigate to Prysm Chain Directory
```
cd prysm-chain
```
## Step 3: Build Prysm
```
make build
```
## Step 4: Make Prysm Binary Executable and Move It
```
chmod +x build/prysmd
mv build/prysmd ~/go/bin/prysmd
```
## Step 5: Configure Your Node
```
CHAIN_ID="prysm_100-1"
MONIKER="insert_your_moniker_here"
VALIDATOR="validatorname"
KEYRING_BACKEND="test"
```
## Step 6: Generate Keys and Save Address/Seed
```
prysmd keys add $VALIDATOR --keyring-backend test 
MY_VALIDATOR_ADDRESS=$(prysmd keys show $VALIDATOR -a --keyring-backend test)
```
## Step 7: Initialize Prysm with Moniker and Chain ID
```
prysmd init $MONIKER --chain-id $CHAIN_ID
```
## Step 8: Download genesis.json
```
wget -O ~/.prysm/config/genesis.json https://github.com/BobboFuxx/prysm-testnet/raw/main/genesis.json
```
## Step 9: Update client.toml Configuration
```
sed -i "s/chain-id = \".*\"/chain-id = \"$CHAIN_ID\"/g" ~/.prysm/config/client.toml
sed -i "s/keyring-backend = \".*\"/keyring-backend = \"$KEYRING_BACKEND\"/g" ~/.prysm/config/client.toml
```
## Step 10: Create Systemd Service for Prysm
```
SERVICE_FILE="/etc/systemd/system/prysmd.service"

echo "[Unit]
Description=Prysm Chain
After=network.target

[Service]
Type=simple
ExecStart=$(command -v prysmd) start
WorkingDirectory=$(eval echo ~$(whoami))
User=$(whoami)
StartLimitInterval=0
RestartSec=3
Restart=on-failure

[Install]
WantedBy=multi-user.target" | sudo tee $SERVICE_FILE > /dev/null

sudo systemctl daemon-reload
sudo systemctl enable prysmd.service
sudo systemctl start prysmd.service
```
## Step 11: Verify Prysm Service Status
```
sudo journalctl -f -u prysmd -o cat
```
## Step 12: Request Testnet Coins
Next, request testnet coins to your validator's wallet address 

```echo $MY_VALIDATOR_ADDRESS ```

## Step 13: Create and Run the Validator
```
# Step 1: Generate your validator key
prysmd tx staking create-validator \
  --amount=1000000000upym \
  --pubkey=$(prysmd tendermint ow-validator) \
  --moniker=$MONIKER \
  --chain-id=$CHAIN_ID \
  --commission-max-change-rate="0.01" \
  --commission-max-rate="0.1" \
  --commission-rate="0.07" \
  --min-self-delegation="1" \
  --gas="auto" \
  --gas-adjustment="1.2" \
  --gas-prices="0.025upym" \
  --from=$VALIDATOR --keyring-backend test
```
# Step 14: Start your validator and make sure it's running with no errors
```
sudo systemctl start prysmd.service
sudo journalctl -f -u prysmd -o cat
```
## Step 15: 
Verify your validator is in sync and that voting power reflects the amount of PYM
You have now set up and started your Prysm validator on the Prysm chain.
