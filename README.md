## damdo/mirror

[![Build Status](https://travis-ci.org/damdo/mirror.svg?branch=master)](https://travis-ci.org/damdo/mirror)

:microscope: HTTP(S) request mirror for Debugging

<br>

### FEATURES:
- Mirrors the received requests and their responses -> in the response body and in its STDOUT
- Returns Client IP and Server Hostname
- Single, cross-platform and statically Linked executable (written in Go)
- Dockerized version available

<br>

### 1) USAGE: 
#### 1.1) TURN ON MIRROR
##### OPTION 1: Start it in a docker container
```
docker run -e "PORT=80" -e "BODY=1" -p 80:80 damdo/mirror
```

##### OPTION 2: Download and start the binary
```
go get github.com/damdo/mirror;
PORT=80 BODY=1 mirror;
```

<br>

#### 1.2) MAKE REQUESTS
#### Example: Make a request with curl to mirror
For example if you make a request form the same host
the ip will be 127.0.0.1
```
curl -X POST -d 'This is the example body' 127.0.0.1:80
```


#### Example Response:
Depending on the headers of the requests the output will differ.
So here the various Headers, Types and Protocols can be inspected.

For example:
```
====================

CLIENT: "192.168.56.101:58072" REQUEST
--------------------------
Time: 1515851216261964065
GET / HTTP/1.1
Host: localhost:80
Accept: */*
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/7.57.0


This is the example body


SERVER: MacBook-Pro-di-Damiano.local RESPONSE
--------------------------
Time: 1515851216262018704
HTTP/1.1 200 OK
Connection: close
Content-Type: text/plain; charset=utf-8

====================
```
(logs can be inspected with `docker logs -f containername`)

<br>

#### 1.3) PARAMETERS

The parameters can be specified as: `PARAMETER:value`

If the docker deployment is used they have to be specified as ENV variables 
like:
```sh
docker run -e "PARAMETER:value" -e "PARAMETER2:value2" ...
```
If the executable binary deployment is used they have to be specified preponed to the mirror command 
```sh
PARAMETER:value PARAMETER2:value2 mirror
```

```
AVAILABLE PARAMETERS:

PORT is an integer (default: PORT=80)
Define the value of the port where mirror will be listening
!warning!: If the PORT is set, and the docker deployment is used,
       	   remember to change the docker port mapping (-p parameter)

BODY is an constrained integer (0 or 1) (default: BODY=1)
Define if the body will be included or excluded form the mirror
```

<br>

### 2) OTHER:
PR are more than welcome
