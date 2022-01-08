#!/usr/bin/env bash
set -e

dir=/tmp/test-strong
rm -rf $dir
mkdir $dir

cd $dir
rm -rf ./a
strong new a
cd ./a/cmd && go build
if [ $? -ne 0 ]; then
  echo "Failed: all"
  exit 1
else
  rm -rf ../../a
fi

cd $dir
rm -rf ./b
strong new b --grpc
cd ./b/cmd && go build
if [ $? -ne 0 ];then
  echo "Failed: --grpc"
  exit 1
else
  rm -rf ../../b
fi

cd $dir
rm -rf ./c
strong new c --http
cd ./c/cmd && go build
if [ $? -ne 0 ]; then
  echo "Failed: --http"
  exit 1
else
  rm -rf ../../c
fi

rm -rf $dir
