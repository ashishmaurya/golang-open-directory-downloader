# Open Directory Downloader using GO

Simple Go Lang Project to Download Multiple file 

How to Use?

``` bat
download.exe -f file.txt -c 4 -d "D:\Downloads"
```

### Build Instruction for linux
``` bat
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -o download go.go
```

### Build Instruction for windows

``` bat
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0
go build -o download.exe go.go
```