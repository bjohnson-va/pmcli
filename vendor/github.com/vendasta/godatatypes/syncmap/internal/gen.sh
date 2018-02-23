#!/usr/bin/env bash

cat generic.go | genny gen "KeyType=$1 ValueType=$2" | cat > ../$1_$2_syncmap.go
