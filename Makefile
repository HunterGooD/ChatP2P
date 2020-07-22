all: linux windows

linux:
	go build -o build/bin/main cmd/chatP2P/main.go

windows:
	go build -o build/bin/main.exe cmd/chatP2P/main.go