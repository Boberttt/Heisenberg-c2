go build -ldflags -H=windowsgui h.go
h-packers\upx.exe -9 -o h-packed.exe h.exe
powershell -c rm h.exe
