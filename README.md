# fs
file/folder sync over tcp(client/server)

# install

## from go source

client
```bash
go get github.com/Chyroc/fs/cmd/fs-cli
```

server
```bash
go get github.com/Chyroc/fs/cmd/fs-svr
```

## from release binary

```bash
curl -L https://github.com/Chyroc/fs/releases/download/v0.1.0/fs_0.1.0_Linux_x86_64.tar.gz > fs.tar.gz && tar zxvf fs.tar.gz
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