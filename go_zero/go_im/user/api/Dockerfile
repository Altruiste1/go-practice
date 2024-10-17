FROM alpine:3.19.4
RUN  mkdir -p /go/etc /go/user/etc  /go/user/log && cd /go/user
COPY etc/user.yaml /go/user/etc/.
COPY api-user-linux /go/user/.
RUN chmod +x /go/user/api-user-linux
RUN chmod 777 /go/user/log
WORKDIR  /go/user
ENTRYPOINT  ["./api-user-linux"]

# docker run -p 8888:8888  -v /Users/huanglin/go/src/go-practice/go_zero/go_im/user/api/log:/go/user/log --name go_zero_v12 -d e690de0