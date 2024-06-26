package exercise

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/timsexperiments/json-protobuf-benchmarking/internal/args"
	"github.com/timsexperiments/json-protobuf-benchmarking/internal/serialize"
)

type Exerciser[T any] struct {
	serializer serialize.Serializer[T]
	limit      int
}

func CreateExerciser[T any](serializer serialize.Serializer[T]) *Exerciser[T] {
	return &Exerciser[T]{serializer: serializer, limit: 100}
}

func (e *Exerciser[T]) WithLimit(limit int) *Exerciser[T] {
	e.limit = limit
	return e
}

func (e *Exerciser[T]) RunExercise(data T) (*ExerciseStats, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	workersAvailable := make(chan struct{}, args.PoolSize)
	errors := make(chan error, 1)

	for range args.PoolSize {
		workersAvailable <- struct{}{}
	}

	stats := createExerciseStats()
	for range e.limit {
		<-workersAvailable
		wg.Add(1)
		go func(data T) {
			defer wg.Done()
			defer func() {
				workersAvailable <- struct{}{}
			}()

			startDeserialization := time.Now()
			serializedData, err := e.serializer.Serialize(data)
			if err != nil {
				if len(errors) == 0 {
					errors <- fmt.Errorf("unable to serialize data: %w", err)
				}
				return
			}
			serializationDuration := time.Since(startDeserialization)

			startSerialization := time.Now()
			_, err = e.serializer.Deserialize(serializedData)
			if err != nil {
				if len(errors) == 0 {
					errors <- fmt.Errorf("unable to deserialize data: %w", err)
				}
				return
			}
			deserializationDuration := time.Since(startSerialization)

			mu.Lock()
			stats.addBytes(len(serializedData))
			stats.serializationTime = append(stats.serializationTime, serializationDuration)
			stats.deserializationTime = append(stats.deserializationTime, deserializationDuration)
			mu.Unlock()
		}(data)
	}

	wg.Wait()
	close(workersAvailable)
	close(errors)

	if len(errors) > 0 {
		return nil, <-errors
	}

	return stats, nil
}

type ExerciseStats struct {
	serializationTime   []time.Duration
	deserializationTime []time.Duration
	totalBytes          int
}

func createExerciseStats() *ExerciseStats {
	return &ExerciseStats{
		serializationTime:   make([]time.Duration, 0),
		deserializationTime: make([]time.Duration, 0),
		totalBytes:          0,
	}
}

func (stats *ExerciseStats) TotalTime() time.Duration {
	return stats.SerializationTime() + stats.DeserializationTime()
}

func (stats *ExerciseStats) AverageSerializationTime() time.Duration {
	return time.Duration(stats.SerializationTime().Nanoseconds() / int64(len(stats.serializationTime)))
}

func (stats *ExerciseStats) SerializationTime() time.Duration {
	total := time.Duration(0)
	for _, time := range stats.serializationTime {
		total += time
	}
	return total
}

func (stats *ExerciseStats) AverageDeserializationTime() time.Duration {
	return time.Duration(stats.DeserializationTime().Nanoseconds() / int64(len(stats.deserializationTime)))
}

func (stats *ExerciseStats) DeserializationTime() time.Duration {
	total := time.Duration(0)
	for _, time := range stats.deserializationTime {
		total += time
	}
	return total
}

func (stats *ExerciseStats) BytesPerSecond(time time.Duration) int {
	return stats.totalBytes
}

func (stats *ExerciseStats) addBytes(totalBytes int) {
	stats.totalBytes += totalBytes
}

func (stats *ExerciseStats) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Total Operations: ")
	buffer.WriteString(fmt.Sprintf("%d", len(stats.serializationTime)))
	buffer.WriteString("\n\nTotal Execution Time: ")
	buffer.WriteString(stats.TotalTime().String())
	buffer.WriteString("\n\nSerialization Info:")
	buffer.WriteString("\nSerialization Time: ")
	buffer.WriteString(stats.SerializationTime().String())
	buffer.WriteString("\nAverage Serialization Time: ")
	buffer.WriteString(stats.AverageSerializationTime().String())
	buffer.WriteString("\nTotal Serialization Bytes: ")
	buffer.WriteString(formatBytes(uint64(stats.totalBytes)))
	buffer.WriteString("\n\nDeserialization Info:")
	buffer.WriteString("\nDeserialization Time: ")
	buffer.WriteString(stats.DeserializationTime().String())
	buffer.WriteString("\nAverage Deserialization Time: ")
	buffer.WriteString(stats.AverageDeserializationTime().String())
	return buffer.String()
}

func formatBytes(bytes uint64) string {
	const (
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)

	switch {
	case bytes >= EB:
		return fmt.Sprintf("%.2f EB", float64(bytes)/EB)
	case bytes >= PB:
		return fmt.Sprintf("%.2f PB", float64(bytes)/PB)
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/TB)
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}
