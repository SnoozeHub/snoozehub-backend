We are using a proxy server for the grpc server from https://github.com/improbable-eng/grpc-web/tree/master/go/grpcwebproxy, to convert http 1.1 requests to http 2

## Services:
- grpc-server: Grpc server for backend, this is the main service for development
- grpc-web-proxy: proxy for the grpc-server (port: 9090)
- mongodb: database (port: 27017)
- mongo-express: web interface to manage mongodb (port: 8081)
# How to use for dev
1. Open The folder in vscode
2. Using remote container extension, select reopen in container
3. Install reccomended tools
4. 
    - Develop (to manage the database go to [localhost:8081](localhost:8081))
    - Is exposed the port 9090 from grpc-web-proxy service if you want to make manual testing from an external grpc web client
# How to use for production
Note: Currenctly there is a public key filter for demo purposes
1. You need locally only "prod folder" and open it
2. Sign in in the container registry
3. Place "mailgun-sending-key.key" file in "secrets" folder
4. Spin up "prod/docker-compose.yaml" with `docker compose up`

# Proposals
- Implement tests for:
    - Book
    - Review
    - GetMyReview
    - RemoveReview
    - RemoveMyBed
    - GetReview
- add correct regex of telegram username
- add indexes for database in setupdb function
- SetProfilePic doesn't check if it is avif
- non optional values in grpc could be nil
