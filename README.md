We are using a proxy server for the grpc server from https://github.com/improbable-eng/grpc-web/tree/master/go/grpcwebproxy, to convert http 1.1 requests to http 2

## Services:
- grpc-server: Grpc server for backend, this is the main service for development
- grpc-web-proxy: proxy for the grpc-server (port: 9090)
- mongodb: database (port: 27017)
- mongo-express: web interface to manage mongodb (port: 8081)
# How to use for dev
1. Place 'mailgun-sending-key.key' file in secrets folder
2. Open The folder in vscode
3. Using remote container extension, select reopen in container
4. Develop or go to [localhost:8081](localhost:8081) to manage the database
# How to use for production
TODO

#TODO
SetProfilePic doesn't check if it is avif
removeAccount remove only the account itself
non optional values are assumed that are optional, it can crash