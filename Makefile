build: gen gobuild

gen:
	go generate ./...

gobuild:
	go build -o demo ./...