package main

import (
	"fmt"
	"time"

	"github.com/timsexperiments/json-protobuf-benchmarking/internal/args"
	"github.com/timsexperiments/json-protobuf-benchmarking/internal/exercise"
	"github.com/timsexperiments/json-protobuf-benchmarking/internal/json"
	"github.com/timsexperiments/json-protobuf-benchmarking/internal/pb"
	"github.com/timsexperiments/json-protobuf-benchmarking/internal/serialize"
)

func main() {
	startTime := time.Now()
	err := args.Parse()
	if err != nil {
		panic(err)
	}
	limit := args.Limit

	jsonSerializer := serialize.CreateJsonSerializer[*json.Person](&json.Person{})
	jsonMessage := json.Person{Name: "Pixel 🐶", Age: 3, Email: "pixel@timsexperiments.foo", Hobbies: []string{"barking", "sleeping", "eating", "playing"}}
	jsonExerciser := exercise.CreateExerciser[*json.Person](jsonSerializer).WithLimit(limit)
	jsonStats, err := jsonExerciser.RunExercise(&jsonMessage)
	if err != nil {
		panic(err)
	}
	fmt.Println("============== START JSON INFO ================")
	fmt.Println(jsonStats)
	fmt.Println("============== END JSON INFO ================")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	pbSerializer := serialize.CreatePbSerializer[*pb.Person](&pb.Person{})
	pbMessage := pb.Person{Name: "Pixel 🐶", Age: 3, Email: "pixel@timsexperiments.foo", Hobbies: []string{"barking", "sleeping", "eating", "playing"}}
	pbExerciser := exercise.CreateExerciser[*pb.Person](pbSerializer).WithLimit(limit)
	pbStats, err := pbExerciser.RunExercise(&pbMessage)
	if err != nil {
		panic(err)
	}
	fmt.Println("============== START PB INFO ================")
	fmt.Println(pbStats)
	fmt.Println("============== END PB INFO ================")
	fmt.Printf("Completed in %v\n", time.Since(startTime))
}
