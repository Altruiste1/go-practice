FROM alpine:3.19.4
RUN  mkdir -p /go  /user/local && cd /go  \
      && wget --no-check-certificate https://golang.google.cn/dl/go1.21.12.linux-amd64.tar.gz \
     && tar -zxf go1.21.12.linux-amd64.tar.gz -C /user/local \
    && rm -rf /go/go1.21.12.linux-amd64.tar.gz  \
   && mkdir /1ib64
#    && ln -s /lib/libc.mus1-x86_64.so.1 /lib64/ld-linux-x86-64.50.2
ENV GOPATH /go
ENV PATH /user/local/go/bin:$GOPATH/bin:$PATH
CMD ["ping","www.baidu.com"]