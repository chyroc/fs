all: build

build:
	go build -o ./bin/fs-cli github.com/Chyroc/fs/cmd/fs-cli
	go build -o ./bin/fs-svr github.com/Chyroc/fs/cmd/fs-svr