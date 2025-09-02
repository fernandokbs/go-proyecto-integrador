# How to build

```bash
sudo apt install -y libvips-dev
```

```bash
go build -o bin/api ./cmd/server/main.go
```

# How to run
```bash
./bin/api
```

# Miniio
```bash
docker run -p 9000:9000 -p 9001:9001 \
  quay.io/minio/minio server /data --console-address ":9001"
```

```bash
mc alias set local http://localhost:9000 $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD
mc mb local/go-image
mc anonymous set download local/go-image
```