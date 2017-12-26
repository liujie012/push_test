#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/push_test/
git pull https://github.com/liujie012/push_test.git
cd webserver/
./webserver &