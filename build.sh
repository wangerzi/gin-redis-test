# bash build.sh v1.0.0
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
docker build . -t wj2015/gin-redis-test:"$1"
