FROM alpine:3.19.4
COPY install_etcd.sh .
RUN chmod +x install_etcd.sh && sh install_etcd.sh
CMD ["nohup","./tmp/etcd-download-test/etcd"]