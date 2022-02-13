# Collect Data Service (Client & Server)
## Dependencies
- go version go1.15.6 linux/amd64
- protobuf
- Ubuntu 20.04
## Install protobuf
```bash
# Install  
apt install -y protobuf-compiler
```

```bash
# Check   
protoc --version
```
## Build

To build the project, use:

```bash
# Use 
go get .
```

## Run Collect Data Service (Server Side)
```bash
# Install  
root@hp:~/admin/go/src/github.com/thanhlam/home-collect-data-svc/server# go run server.go
```
## Run Collect Data Service (Client Side)
```bash
# Install  
root@hp:~/admin/go/src/github.com/thanhlam/home-collect-data-svc/client# go run client.go
```
