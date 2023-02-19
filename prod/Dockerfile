FROM golang
WORKDIR /app
ADD go.mod .
ADD go.sum .
ADD *.go .
ADD mail mail
ADD grpc_gen grpc_gen
RUN go mod download
RUN go build -tags prod
ENTRYPOINT ./snoozehub-backend