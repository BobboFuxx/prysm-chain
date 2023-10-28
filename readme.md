<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Prysm Validator Setup Guide</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            margin: 20px;
            padding: 20px;
        }

        h1, h2 {
            color: #007bff;
        }

        pre {
            background-color: #f8f9fa;
            padding: 10px;
            border-radius: 5px;
            overflow-x: auto;
        }
    </style>
</head>

<body>

    <h1>Prysm Validator Setup Guide</h1>

    <h2>Prerequisites</h2>
    <p><strong>Install Go:</strong> Visit the <a href="https://golang.org/dl/">official Go installation page</a> to
        download and install Go. Verify your installation by running <code>go version</code> in the terminal.</p>

    <h2>Step 1: Clone Prysm Chain Repository</h2>
    <pre><code>git clone https://github.com/BobboFuxx/prysm-chain/</code></pre>

    <h2>Step 2: Navigate to Prysm Chain Directory</h2>
    <pre><code>cd prysm-chain</code></pre>

    <h2>Step 3: Build Prysm</h2>
    <pre><code>make build</code></pre>

    <h2>Step 4: Make Prysm Binary Executable and Move It</h2>
    <pre><code>chmod +x build/prysmd<br>mv build/prysmd ~/go/bin/prysmd</code></pre>

    <h2>Step 5: Configure Your Validator</h2>
    <pre><code>
CHAIN_ID="prysm_100-1"
MONIKER="insert_your_moniker_here"
VALIDATOR="validatorname"
KEYRING_BACKEND="test"
    </code></pre>

    <h2>Step 6: Generate Keys and Save Address/Seed</h2>
    <pre><code>
prysmd keys add $VALIDATOR --keyring-backend test 
MY_VALIDATOR_ADDRESS=$(prysmd keys show $VALIDATOR -a --keyring-backend test)
    </code></pre>

    <h2>Step 7: Initialize Prysm with Moniker and Chain ID</h2>
    <pre><code>prysmd init $MONIKER --chain-id $CHAIN_ID</code></pre>

    <h2>Step 8: Update Stake Denomination in genesis.json</h2>
    <pre><code>
sed -i 's/stake/upym/g' ~/.prysm/config/genesis.json
    </code></pre>

    <h2>Step 9: Update client.toml Configuration</h2>
    <pre><code>
sed -i "s/chain-id = \".*\"/chain-id = \"$CHAIN_ID\"/g" ~/.prysm/config/client.toml
sed -i "s/keyring-backend = \".*\"/keyring-backend = \"$KEYRING_BACKEND\"/g" ~/.prysm/config/client.toml
    </code></pre>

    <h2>Step 10: Create Systemd Service for Prysm</h2>
    <pre><code>
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
    </code></pre>

    <h2>Step 11: Verify Prysm Service Status</h2>
    <pre><code>
sudo journalctl -f -u prysmd -o cat
    </code></pre>

    <h2>Step 12: Request Testnet Coins</h2>
    <p>Next, request testnet coins to your validator's wallet address (<code>$MY_VALIDATOR_ADDRESS</code>).</p>

    <h2>Step 13: Create and Run the Validator</h2>
    <pre><code>
# Step 1: Generate your validator key
prysmd tx staking create-validator \
  --amount=1000000000upym \
  --pubkey=$(prysmd tendermint show-validator) \
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

# Step 2: Start your validator
prysmd tx staking start-validator \
  --from=$VALIDATOR --keyring-backend test
    </code></pre>

    <p>You have now set up and started your Prysm validator on the Prysm chain. Remember
