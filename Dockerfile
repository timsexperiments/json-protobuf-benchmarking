FROM --platform=$BUILDPLATFORM golang:1.22 AS builder

COPY . .

RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

RUN protoc --proto_path=proto --go_out=internal/pb --go_opt=paths=source_relative proto/person.proto

RUN go mod download

ARG TARGETOS
ARG TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /json-pb-benchmark cmd/main.go

FROM scratch

COPY --from=builder /json-pb-benchmark /json-pb-benchmark

ENTRYPOINT ["/json-pb-benchmark"]
