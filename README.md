## None

The current result(use docker) is not what I expected, maybe some thing wrong with network emulation, or docker ..


### In Real World (multi-datacenter)

Is there some thing wrong with Golang(IO model/runtime)?

pprof, perf_event and other profile tools doesn't give a shit..

```
$ python client.py --addr 10.xx.xx.xx:8020
0.621856927872 s
0.215888977051 s
0.106964111328 s
0.129124879837 s
0.132770061493 s
0.132673025131 s
0.162160873413 s
0.181607961655 s
0.180653095245 s
0.180568933487 s
[AVG] 0.204426884651 s
$ ./client -addr 10.xx.xx.xx:8020
228.383141ms
96.566201ms
107.525007ms
121.94838ms
127.089427ms
126.70438ms
126.264553ms
126.128879ms
125.931539ms
125.36633ms
[AVG] 131.190783ms
$ ./client -addr 10.xx.xx.xx:8020 -rbufsz 4194304
896.057939ms
843.960815ms
843.428539ms
844.362365ms
843.698511ms
843.693909ms
843.272623ms
843.934291ms
843.394633ms
843.344434ms
[AVG] 848.914805ms
$ python client.py --addr 10.xx.xx.xx:8020 --rbufsz 4194304
0.227913141251 s
0.0962738990784 s
0.102230072021 s
0.12353181839 s
0.141967058182 s
0.170171976089 s
0.172912120819 s
0.172124862671 s
0.171636104584 s
0.171622037888 s
[AVG] 0.155038309097 s
```

### Use docker

```
# build go programs
make build-linux

# build and run containers
docker-compose up -d
docker-compose ps

# download https://github.com/alexei-led/pumba/releases
# delay 25ms
pumba netem --duration 1h delay --time 25 --jitter 0 have-fun-1_server_1

# show results
docker attach have-fun-1_client_1
ping server
sh run_client.sh

# stop containers
docker-compose stop
```
