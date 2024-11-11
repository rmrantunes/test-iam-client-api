watch-dev: 
	$(HOME)/go/bin/air --build.cmd "go build -o bin/api cmd/main.go" --build.bin "./bin/api"