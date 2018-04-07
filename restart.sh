#! /bin/bash
rm -rf XiaodiTest/XiaodiServer
rm -rf XiaodiTest/XiaodiServer.tar.gz
rm -rf XiaodiTest/conf
rm -rf XiaodiTest/restart.sh
rm -rf XiaodiTest/sftp-config.json
rm -rf XiaodiTest/xiaodi.sh
mv XiaodiServer.tar.gz XiaodiTest/
cd XiaodiTest/
tar -zxvf XiaodiServer.tar.gz
kill -9 $(lsof -i | grep 1688 | awk '{print $2}')
nohup ./XiaodiServer &
