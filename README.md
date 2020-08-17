# Open Directory Downloader using GO

### Build Instruction for linux
``` bat
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -o download-linux go.go
```

### Build Instruction for windows

``` bat
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0
go build -o download-windows go.go
```