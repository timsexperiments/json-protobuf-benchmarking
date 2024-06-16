# JSON vs. Protobuf Serialization Benchmarking

This project benchmarks the performance of JSON and Protobuf serialization and deserialization. It measures the time taken to serialize and deserialize a predefined message structure repeatedly using both JSON and Protobuf, and reports the statistics.

## Motivation

Serialization is a crucial part of many applications, especially those involving data transfer between different systems or components. JSON and Protobuf are two popular serialization formats, each with its own advantages:

- **JSON:** Human-readable and widely used, especially in web APIs.
- **Protobuf:** Efficient and compact binary format, often used in performance-critical applications.

This project aims to provide a simplified solution to compare the performance of these two serialization formats in a straightforward manner.

## Setup

### Prerequisites

- Go 1.22 or later
- Protocol Buffers compiler (`protoc`)

### Steps

1. **Clone the repository:**

   ```sh
   git clone https://github.com/timsexperiments/json-protobuf-benchmarking.git
   cd json-protobuf-benchmarking
   ```

2. **Install dependencies:**

   ```sh
   go mod download
   ```

3. **Install the go protoc generator:**

   ```sh
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   ```

4. **Compile the proto:**

   ```sh
   protoc --proto_path=proto --go_out=internal/pb --go_opt=paths=source_relative proto/person.proto
   ```

## Running the Benchmark

### From Source

1.  **Clone the repository:**

    ```sh
    git clone https://github.com/timsexperiments/json-protobuf-benchmarking.git
    cd json-protobuf-benchmarking
    ```

2.  **Build the project:**

    ```sh
    go build -o bin/json-pb-benchmark cmd/main.go
    ```

3.  **Run the executable:**

    ```sh
    ./bin/json-pb-benchmark -p 4 4000000
    ```

    - `-p`: Number of parallel workers (default is the number of CPU cores).
    - `4000000`: The number of iterations for the benchmark (default is 1,000,000).

## Using Docker

1. **Build the Docker image:**

   ```sh
   docker build -t json-pb-benchmark .
   ```

2. **Run the benchmark:**

   ```sh
   docker run --rm json-pb-benchmark -p 4 1000000
   ```

## Output

The benchmark will output statistics for both JSON and Protobuf serialization and deserialization, including the total operations, total execution time, average serialization and deserialization times, and total bytes processed.

Example output:

```
============== START JSON INFO ================
Total Operations: 1000000
Total Execution Time: 2m34s
Serialization Time: ...
Average Serialization Time: ...
Total Serialization Bytes: ...
Deserialization Time: ...
Average Deserialization Time: ...
============== END JSON INFO ================


============== START PB INFO ================
Total Operations: 1000000
Total Execution Time: 1m45s
Serialization Time: ...
Average Serialization Time: ...
Total Serialization Bytes: ...
Deserialization Time: ...
Average Deserialization Time: ...
============== END PB INFO ================
Completed in 4m19s
```

## License

Feel free to adjust the README content as needed.
