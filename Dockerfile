FROM alpine:latest
# replace China mirror
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

RUN apk update && \
    apk add \
    ca-certificates && \
    rm -rf /var/cache/apk/*

COPY drone-goreportcard /bin

ENTRYPOINT ["/bin/drone-goreportcard"]
