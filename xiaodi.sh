#! /bin/bash

echo 'pack'
bee pack -be GOOS=linux
echo 'upload'
scp XiaodiServer.tar.gz ubuntu@119.29.164.153:
rm XiaodiServer.tar.gz
echo 'restart'
ssh ubuntu@119.29.164.153 'bash -s' < restart.sh
