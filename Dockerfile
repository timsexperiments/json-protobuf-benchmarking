FROM golang:1.22 as builder

COPY . .

# Generate compiled proto files
RUN apt-get update
RUN apt install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

RUN protoc --proto_path=proto --go_out=internal/pb --go_opt=paths=source_relative proto/person.proto

# Build the binary
RUN go mod download

RUN go build -o bin/json-pb-benchmark cmd/main.go

FROM scratch

COPY --from=builder /go/bin/json-pb-benchmark /json-pb-benchmark