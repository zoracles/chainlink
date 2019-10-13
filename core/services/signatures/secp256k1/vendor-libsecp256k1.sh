#!/bin/bash

# Cribbed from
# https://github.com/ElementsProject/rust-secp256k1-zkp/blob/8e0d3/contrib/vendor-libsecp-zkp.sh
#
# This changes the symbol names in libsecp256k1 so they can be referenced via
# cgo in chainlink without conflicting with the existing symbols in go-ethereum
# (which also references libsecp256k1, but does not expose the functionality we
# need.)
#
# Expects to be invoked with cwd set to its parent.
#
# This script should only be run when updates to the libsecp256k1 dependency are
# necessary. When that happens, be sure to diff the result against the old
# version, because this is a very fragile way to change the symbol names (it
# just sed's them!)

set -e

script_name=$(basename $0)
test -f $script_name || \
    (echo "Run this script from its containing directory"; exit 1)
test -f libsecp256k1 && \
    (echo "Moving existing copy ouf libsecp256k1 out of the way first"; exit 1)

git clone https://github.com/bitcoin-core/secp256k1 libsecp256k1

# Record the HEAD commit
save_file=./libsecp256k1-HEAD-revision.txt
echo "\# This file was automatically created by $0" > $save_file
(cd libsecp256k1; git rev-parse HEAD) >> $save_file

rm -rf libsecp256k1/.git # Remove git history from libsecp256k1 clone

# Replace symbols starting with secp256k1_ with secp256k1_kyber_, except on
# #include lines, or in dot files.
find libsecp256k1 -not -path '*/\.*' -type f -print0 | \
    xargs -0 sed -i '/^#include/! s/secp256k1_/secp256k1_kyber_/g'

# Generate config files, and the 
pushd libsecp256k1
./autogen.sh
./configure
make src/ecmult_static_context.h
popd
