FROM golang
WORKDIR /app
ADD go.mod .
ADD go.sum .
ADD *.go .
ADD grpc_gen grpc_gen
RUN go mod download
RUN go build
ENTRYPOINT ./snoozehub-backend
