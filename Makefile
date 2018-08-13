clean:
	-rm bin/*

dep:
	dep ensure

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/cctocb_handler cmd/cctocb_handler/main.go
	-strip bin/*
	-upx bin/*
