#!/bin/bash

BASTION_WORK_DIR=$(pwd)
export BASTION=${BASTION_WORK_DIR}

ID=$(ps -ef | grep bastion | grep -v grep | grep -v PPID | awk '{ print $2}')
echo "old pid: $ID"

echo "-------1--------"
for id in $ID; do
  kill -9 "$id"
  echo "killed $id"
done
echo "-------2--------"

echo "restart..."
chmod +x bastion
nohup ./bastion -conf "./secret.toml" >out.log 2>&1 &

echo "new pid："
ID=$(ps -ef | grep bastion | grep -v grep | grep -v PPID | awk '{ print $2}')
echo "$ID"