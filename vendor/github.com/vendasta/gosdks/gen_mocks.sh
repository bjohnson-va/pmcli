#!/bin/bash
if [[ "$#" -ne 1 ]]; then
    echo "Usage: ./gen_mocks <dirname>"
    exit
fi
DIR="$1"
echo "generating mocks for $DIR"
(cd "$DIR" || exit; mockery -all  -dir . -output ./mocks -case underscore -inpkg)
