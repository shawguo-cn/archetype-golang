#!/usr/bin/env bash


#Ubuntu/Debian (clean box)
#
#sudo apt-get install build-essential libssl-dev git -y
#git clone https://github.com/wg/wrk.git
#cd wrk
#make
## move the executable to somewhere in your PATH, ex:
#sudo cp wrk /usr/local/bin


#This runs a benchmark for 30 seconds, using 12 threads, and keeping 400 HTTP connections open.
wrk -t12 -c50000 -d60s http://10.253.11.216:9091/ping

#wrk -t12 -c50000 -d20s http://10.253.11.216:9092/bench/mongo-get