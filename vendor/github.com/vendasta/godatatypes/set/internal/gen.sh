#!/usr/bin/env bash

if [[ $# -eq 0 ]]; then
    cat generic.go | genny gen "Type=BUILTINS" | cat > ../set.go
else
    cat generic.go | genny gen "Type=$1" | cat > ../$1set.go
fi

