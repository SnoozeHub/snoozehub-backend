FROM ubuntu
WORKDIR /tmp
RUN apt update
RUN apt install -y wget unzip
RUN wget https://github.com/improbable-eng/grpc-web/releases/download/v0.15.0/grpcwebproxy-v0.15.0-linux-x86_64.zip
RUN unzip grpcwebproxy-v0.15.0-linux-x86_64.zip
RUN mv dist/grpcwebproxy-v0.15.0-linux-x86_64 /bin/grpcwebproxy
ENTRYPOINT grpcwebproxy --backend_addr=grpc-server:9090 --run_tls_server=false