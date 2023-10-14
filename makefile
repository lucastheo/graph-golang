test:
	go test ./...

watch-test:
	nodemon --exec "go test ./..." -e go -w ./