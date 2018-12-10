# fs
file/folder sync over tcp(client/server)

# install

client
```bash
go get github.com/Chyroc/fs/cmd/fs-cli
```

server
```bash
go get github.com/Chyroc/fs/cmd/fs-svr
```

# usage

server
```bash
fs-svr -port 1234 -mode pull
```

client
```bash
fs-cli -host <host> -port <port> -dir ./dir-to-sync
```