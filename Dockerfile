FROM golang
WORKDIR /app
ADD go.mod .
ADD go.sum .
ADD *.go .
ADD dev_vs_prod dev_vs_prod
ADD assets assets
ADD grpc_gen grpc_gen
RUN go mod download
RUN go build -tags prod
ENTRYPOINT ./snoozehub-backend