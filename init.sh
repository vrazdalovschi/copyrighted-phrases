#!/usr/bin/env bash

rm -rf ~/.copyrightedphrasesd
rm -rf ~/.copyrightedphrasescli

copyrightedphrasesd init test --chain-id=namechain

copyrightedphrasescli config output json
copyrightedphrasescli config indent true
copyrightedphrasescli config trust-node true
copyrightedphrasescli config chain-id namechain
copyrightedphrasescli config keyring-backend test

copyrightedphrasescli keys add user1
copyrightedphrasescli keys add user2

copyrightedphrasesd add-genesis-account $(copyrightedphrasescli keys show user1 -a) 1000nametoken,100000000stake
copyrightedphrasesd add-genesis-account $(copyrightedphrasescli keys show user2 -a) 1000nametoken,100000000stake

copyrightedphrasesd gentx --name user1 --keyring-backend test

echo "Collecting genesis txs..."
copyrightedphrasesd collect-gentxs

echo "Validating genesis file..."
copyrightedphrasesd validate-genesis
