# docker build . -t wj2015/redis-test-gin:latest

FROM alpine:3.13
COPY gin-redis-test /usr/local/bin

CMD ["/usr/local/bin/gin-redis-test"]
EXPOSE 80