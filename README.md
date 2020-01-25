# Supervisord status page

Simple supervisord statuspage for healthcheck. Just call simple `/healtz` via HTTP and get 200 or 500 and process list.

## Usage

### Supervisord configuration

```
[supervisord]
nodaemon=true

; turn on XML-RPC interface
[rpcinterface:supervisor]
supervisor.rpcinterface_factory = supervisor.rpcinterface:make_main_rpcinterface

; listen on localhost 9001
[inet_http_server]         ; inet (TCP) server disabled by default
port=127.0.0.1:9001


[program:backend]
command=/bin/bash -lc "/usr/local/bin/svd-statuspage"
autorestart=true
```

### Config variables

* `RPC_URL` - supervisord RPC endpoint, default `http://127.0.0.1:9001/RPC2`
* `LISTEN_IP` - listen IP, default `0.0.0.0`
* `LISTEN_PORT` - listen port, default `8080`

### Build & run

```
make
make run
./svd-statuspage
LISTEN_PORT=9090 LISTEN_IP=127.0.0.1 ./svd-statuspage
```

### Get status

```bash
curl -v localhost:8080/healtz | jq .
```

*Output ERROR*

Return `500 Internal Server Error`

```json
{
  "status": "ERROR",
  "processes": [
    {
      "name": "backend",
      "status": "FATAL",
      "description": "Exited too quickly (process log may have details)",
      "pid": "0"
    }
  ]
}
```

*Output OK* 

Return `200 OK`

```json
{
  "status": "OK",
  "processes": [
    {
      "name": "backend",
      "status": "RUNNING",
      "description": "pid 78855, uptime 0:00:05",
      "pid": "78855"
    }
  ]
}
```
