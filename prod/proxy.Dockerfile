FROM ubuntu
WORKDIR /tmp
RUN apt update
RUN apt install -y wget unzip
RUN wget https://github.com/improbable-eng/grpc-web/releases/download/v0.14.0/grpcwebproxy-v0.14.0-arm64.zip
RUN unzip grpcwebproxy-v0.14.0-arm64.zip
RUN mv dist/grpcwebproxy-v0.14.0-arm64 /bin/grpcwebproxy
ENTRYPOINT /bin/grpcwebproxy --backend_addr=grpc-server:9090 --server_tls_cert_file=/secrets/live/snoozehub.it/cert.pem --server_tls_key_file=/secrets/live/snoozehub.it/privkey.pem