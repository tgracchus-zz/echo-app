FROM golang:1.10.1
WORKDIR /server
COPY echo-linux-amd64 /server
COPY config/ /server/config
ENTRYPOINT ["/server/echo-linux-amd64"]