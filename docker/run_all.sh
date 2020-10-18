#!/bin/sh
#set -euo pipefail

mkdir -p /root/log
touch /root/log/copyrightedphrasesd.log
./run_serverd.sh $1 >>/root/log/copyrightedphrasesd.log &

sleep 4
echo Starting Rest Server...

./run_rest_server.sh
