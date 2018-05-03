FROM golang:1.10.1
WORKDIR /server
COPY echo-linux-amd64 /server
ENTRYPOINT /server/echo-linux-amd64