# Test Swarm DNSRR (DNS Round Robin)

## Environment
* `docker version`:
```
Client: Docker Engine - Community
 Version:           19.03.14
 API version:       1.40
 Go version:        go1.13.15
 Git commit:        5eb3275d40
 Built:             Tue Dec  1 19:20:17 2020
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          19.03.14
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.13.15
  Git commit:       5eb3275d40
  Built:            Tue Dec  1 19:18:45 2020
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.3.9
  GitCommit:        ea765aba0d05254012b0b9e595e995c09186427f
 runc:
  Version:          1.0.0-rc10
  GitCommit:        dc9208a3303feef5b3839f4323d9beb36df0a9dd
 docker-init:
  Version:          0.18.0
  GitCommit:        fec3683
```

* `docker-compose version`:
```
docker-compose version 1.27.4, build 40524192
docker-py version: 4.3.1
CPython version: 3.7.7
OpenSSL version: OpenSSL 1.1.0l  10 Sep 2019
```

## Test
* `docker swarm init`: Enable Swarm mode
* `docker stack deploy -c docker-compose.yml swarm_dnsrr_test`:
```
Creating network swarm_dnsrr_test_default
Creating service swarm_dnsrr_test_server
Creating service swarm_dnsrr_test_client
```
* `docker ps -a`:
```
CONTAINER ID        IMAGE                          COMMAND                  CREATED             STATUS                 PORTS                                                                     NAMES
37f30d9bb82e        alpine:latest                  "sh -c 'apk add curl…"   20 seconds ago      Up 19 seconds                                                                                    swarm_dnsrr_test_client.1.985c9h7im1znrs5ab38o55jgl
a4cb53c531b4        golang:1.15.6-buster           "bash -c 'go run mai…"   25 seconds ago      Up 23 seconds                                                                                    swarm_dnsrr_test_server.3.bfy5q0ytgz42ymm28a72hdk31
3f97ac050886        golang:1.15.6-buster           "bash -c 'go run mai…"   25 seconds ago      Up 23 seconds
                                                          swarm_dnsrr_test_server.2.yyactmrfgb76cd2mdvg0j4cf4
9f2e3554886f        golang:1.15.6-buster           "bash -c 'go run mai…"   25 seconds ago      Up 22 seconds                                                                                    swarm_dnsrr_test_server.4.ssxhn0q3aaj60m81l441dc4n4
406dbd45f556        golang:1.15.6-buster           "bash -c 'go run mai…"   25 seconds ago      Up 23 seconds                                                                                    swarm_dnsrr_test_server.1.ivmjxw6lctajwfx0y8z6c1zcy
349a5ebb2f24        golang:1.15.6-buster           "bash -c 'go run mai…"   25 seconds ago      Up 24 seconds                                                                                    swarm_dnsrr_test_server.5.jm4ek3dyg6tiwth30u7o9bgd
```
* `docker exec -it 37f30d9bb82e curl http://server:8090/interfaces`: You will notice that everytime you query this name, the responder container is different.
```
IP: 127.0.0.1
IP: 10.0.2.5
IP: 172.18.0.6
```
* `docker exec -it 37f30d9bb82e curl http://server:8090/interfaces`:
```
IP: 127.0.0.1
IP: 10.0.2.7
IP: 172.18.0.4
```

* `docker exec -it 37f30d9bb82e ping server`: However, it seems that the responder of `ping` is the swarm DNS load balancer.

## References
* <https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go>
