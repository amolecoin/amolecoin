#!/bin/sh
COMMAND="amolecoin --data-dir $DATA_DIR --wallet-dir $WALLET_DIR $@"

adduser -D -u 10000 amolecoin

if [[ \! -d $DATA_DIR ]]; then
    mkdir -p $DATA_DIR
fi
if [[ \! -d $WALLET_DIR ]]; then
    mkdir -p $WALLET_DIR
fi

chown -R amolecoin:amolecoin $( realpath $DATA_DIR )
chown -R amolecoin:amolecoin $( realpath $WALLET_DIR )

su amolecoin -c "$COMMAND"
