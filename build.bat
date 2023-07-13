@echo off
go build -ldflags "-w -s" -o release/client.exe client.go
go build -ldflags "-w -s" -o release/server.exe server.go

@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build -ldflags "-w -s" -o release/client32.exe client.go
go build -ldflags "-w -s" -o release/server32.exe server.go

@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/client client.go
go build -ldflags "-w -s" -o release/server server.go
