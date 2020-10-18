[![Build Status](https://travis-ci.com/vrazdalovschi/copyrighted-phrases.svg?branch=main)](https://travis-ci.com/vrazdalovschi/copyrighted-phrases)
# copyrighted-phrases

A blockchain application using cosmos-sdk that is a 'copyrighted phrases' registry.

App enforces uniqueness of all registered phrases.

## Quick Start

```
make install
make test
```

## Dockerized

Provided a docker image to help with test setups.

Build: 
```shell script
docker build -t github.com/vrazdalovschi/copyrighted-phrases:latest .
```

Run:
```shell script
# This will start both copyrightedphrasesd and copyrightedphrasescli rest-server, 
# only copyrightedphrasescli output is shown on the screen
docker run --rm -it -p 26657:26657 -p 26656:26656 -p 1317:1317 github.com/vrazdalovschi/copyrighted-phrases:latest
```
## Api examples

* Register Copyrighted Phrase
```shell script
curl -X POST -s http://localhost:1317/copyrightedphrases/text --data-binary '{"base_req":{"from":"'$(copyrightedphrasescli keys show user1 -a)'","chain_id":"namechain"},"text":"EeeCosmos","owner":"'$(copyrightedphrasescli keys show user1 -a)'"}' > unsignedTx.json

copyrightedphrasescli tx sign unsignedTx.json --from user1 --offline --chain-id namechain --sequence 1 --account-number 2 > signedTx.json

copyrightedphrasescli tx broadcast signedTx.json
```

* Remove Copyrighted Phrase
```shell script
curl -X DELETE -s http://localhost:1317/copyrightedphrases/text --data-binary '{"base_req":{"from":"'$(copyrightedphrasescli keys show user1 -a)'","chain_id":"namechain"},"text":"EeeCosmos","owner":"'$(copyrightedphrasescli keys show user1 -a)'"}' > unsignedTx.json

copyrightedphrasescli tx sign unsignedTx.json --from user1 --offline --chain-id namechain --sequence 1 --account-number 2 > signedTx.json

copyrightedphrasescli tx broadcast signedTx.json
```

* Show all Copyrighted Phrases for the sdk.AccAddress
```shell script
copyrightedphrasescli query copyrightedphrases list-copyrighted-texts cosmos1m8dyn2rfagx85hmes68sm5270flc4sk75mymn2
```

* Describe the owner of Copyrighted Phrase
```shell script
copyrightedphrasescli query copyrightedphrases get-copyrighted-text 'CosmoeEEE'
```