#!/bin/bash

VERSION=2.0.0

wget "https://github.com/sneuder/tig/releases/download/v$VERSION/wkspace-$VERSION.tar.gz" -O "tig-$VERSION.tar.gz"

tar -xzf "tig-$VERSION.tar.gz"

sudo mv ./tig /usr/local/bin/tig

rm "tig-$VERSION.tar.gz"

echo "Installation completed for tig version $VERSION."