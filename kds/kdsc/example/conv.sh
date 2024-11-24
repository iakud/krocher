#!/usr/bin/env bash

[[ "$TRACE" ]] && set -x
pushd `dirname "$0"` > /dev/null
trap `popd > /dev/null` EXIT

colorful=false
tput setaf 7 > /dev/null 2>&1
if [[ $? -eq 0 ]]; then
    colorful=true
fi

function printError() {
    $colorful && tput setaf 1
    >&2 echo "Error: $@"
    $colorful && tput setaf 7
}

function printImportantMessage() {
    $colorful && tput setaf 3
    >&2 echo "$@"
    $colorful && tput setaf 7
}

function printUsage() {
    $colorful && tput setaf 3
    >&2 echo "$@"
    $colorful && tput setaf 7
}

NEW_PATH=`../../../privatePath.sh`
[[ $? -ne 0 ]] && exit 1
PATH="$NEW_PATH"

rm -rf examplekds
mkdir -p examplekds

kdsc --tpl_path=tpls --out=examplekds *.kds
[[ $? -ne 0 ]] && exit 1

rm -rf examplepb
mkdir -p examplepb

cp test.proto examplekds/

protoc -I=./examplekds -I=../../../local/protoc/include \
	--go_out=paths=source_relative:./examplepb \
	`find examplekds -name "*.proto"`