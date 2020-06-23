#!/bin/bash

DEJA_WORK_DIR=$(pwd)
export DEJA=${DEJA_WORK_DIR}

ID=$(ps -ef | grep deja | grep -v grep | grep -v PPID | awk '{ print $2}')
echo "old pid: $ID"

echo "-------1--------"
for id in $ID; do
  kill -9 "$id"
  echo "killed $id"
done
echo "-------2--------"

echo "restart..."
chmod +x deja
nohup ./deja -conf "./conf/prod.toml" >out.log 2>&1 &

echo "new pid："
ID=$(ps -ef | grep deja | grep -v grep | grep -v PPID | awk '{ print $2}')
echo "$ID"